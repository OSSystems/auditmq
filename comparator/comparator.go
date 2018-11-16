package comparator

import (
	"time"

	"github.com/OSSystems/auditmq/publisher"
	"github.com/OSSystems/auditmq/storage"
	"github.com/OSSystems/pkg/log"
)

type comparator struct {
	pub     publisher.Publisher
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

	pub, err := publisher.New()
	if err != nil {
		return nil, err
	}

	return &comparator{
		pub:     pub,
		storage: storage,
		stop:    make(chan bool),
		done:    make(chan error),
	}, nil
}

func (c *comparator) logExpectatives(service storage.Service) {
	serviceName := service.Name()
	for fieldName, field := range service.FieldsToMatch() {
		ownerService := field.OwnedBy
		ownerField, err := ownerService.GetField(fieldName)
		if err != nil {
			log.Errorf("Error when logging expectatives for %s", serviceName)
			continue
		}

		ownerFieldSlice := ownerField.ToSlice()
		log.Infof("%s.%s - expected: %v ; actual: %v", serviceName, fieldName, ownerFieldSlice, field.ToSlice())
	}
}

func (c *comparator) buildFieldStatusList(service storage.Service) []storage.Status {
	statusList := []storage.Status{}

	for fieldName, field := range service.FieldsToMatch() {
		ownerService := field.OwnedBy
		ownerField, err := ownerService.GetField(fieldName)
		if err != nil {
			log.Error(err)
			continue
		}

		ownerFieldSlice := ownerField.ToSlice()
		serviceFieldSlice := field.ToSlice()

		log.Debugf("%s.%s - expected: %v ; actual: %v", service.Name(), fieldName, ownerFieldSlice, serviceFieldSlice)
		result, err := ownerField.Equals(field)
		if err != nil {
			log.Debugf("Comparator raises: %v", err)
			continue
		}

		var newStatus storage.Status
		if result {
			newStatus = storage.Sync
		} else {
			newStatus = storage.NotSync
		}

		statusList = append(statusList, newStatus)
	}

	return statusList
}

func (c *comparator) compare() error {
	for _, service := range c.storage.Services() {
		serviceName := service.Name()
		oldStatus := service.Status()

		statusList := c.buildFieldStatusList(service)
		newStatus := storage.DetermineMinStatus(statusList)
		if oldStatus == newStatus {
			log.Debugf("%s service keep the same status (%s)", serviceName, newStatus.String())
			continue
		}

		log.Infof("%s service status: (%s)", serviceName, newStatus.String())
		service.SetStatus(newStatus)
		c.logExpectatives(service)

		err := c.pub.PublishServiceStatus(service)
		if err != nil {
			return err
		}

		err = c.pub.PublishGlobalStat()
		if err != nil {
			return err
		}
	}

	return nil
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

			lastErr = c.compare()
			if lastErr != nil {
				log.Errorf("Comparator raises an error %v", lastErr)
			}

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
