package comparator

import (
	"time"

	"github.com/OSSystems/pkg/log"
	"github.com/rodrigoapereira/auditmq/storage"
)

type comparator struct {
	storage storage.Storage
	ticker  *time.Ticker
	stop    chan bool
	done    chan error
}

func New() (*comparator, error) {
	storage, err := storage.GetStorage()
	if err != nil {
		return nil, err
	}

	return &comparator{
		storage: storage,
		stop:    make(chan bool),
		done:    make(chan error),
	}, nil
}

func (c *comparator) compare() {
	for _, service := range c.storage.Services() {
		for fieldName, field := range service.FieldsToMatch() {
			ownerService := field.OwnedBy
			ownerField, err := ownerService.GetField(fieldName)
			if err != nil {
				log.Error(err)
				continue
			}

			log.Debugf("%s - expected: %v ; actual: %v", service.Name(), ownerField.ToSlice(), field.ToSlice())
			result, err := ownerField.Equals(field)
			if err != nil {
				log.Error(err)
				continue
			}

			if result {
				log.Infof("%s service stable", service.Name())
				service.SetStatus(storage.Sync)
			} else {
				log.Errorf("%s service unstable - %s is not equal to the owner", service.Name(), field.Name)
				service.SetStatus(storage.NotSync)
			}
		}
	}
}

func (c *comparator) run() {
	var lastErr error

	for {
		select {
		case <-c.stop:
			c.done <- lastErr
			break
		case <-c.ticker.C:
			log.Debug("Running comparison")
			c.compare()
			log.Debug("Finished comparison")
		}
	}
}

func (c *comparator) Start() {
	c.ticker = time.NewTicker(5 * time.Second)
	go c.run()
}

func (c *comparator) Shutdown() {
	c.stop <- true
	c.ticker.Stop()
	<-c.done
}
