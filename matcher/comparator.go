package matcher

import (
	"time"

	"github.com/OSSystems/pkg/log"
	"github.com/rodrigoapereira/auditmq/storage"
)

type Comparator struct {
	ticker *time.Ticker
	stop   chan bool
	done   chan error
}

func NewComparator() *Comparator {
	return &Comparator{
		stop: make(chan bool),
		done: make(chan error),
	}
}

func (c *Comparator) compare() {
	for service, details := range m.Services {
		for field, option := range details.MatchedData {
			serviceData := storage.GetStorageData(service)
			data := serviceData.GetBuffer(field)

			ownerData := storage.GetStorageData(option.OwnedBy)
			dataByOwner := ownerData.GetBuffer(field)

			log.Debugf("%s - expected: %v ; actual: %v", service, dataByOwner.ToSlice(), data.ToSlice())
			result, err := option.Compare(dataByOwner.ToSlice(), data.ToSlice())
			if err != nil {
				log.Error(err)
				continue
			}

			if result {
				log.Infof("%s service stable - %s\n", service, result)
			} else {
				log.Errorf("%s service unstable - %s is not equal to the owner \n", service, field)
			}

		}
	}
}

func (c *Comparator) run() {
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

func (c *Comparator) Start() {
	c.ticker = time.NewTicker(5 * time.Second)
	go c.run()
}

func (c *Comparator) Shutdown() {
	c.stop <- true
	c.ticker.Stop()
	<-c.done
}
