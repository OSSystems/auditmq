package publisher

import (
	"encoding/json"

	"github.com/NeowayLabs/wabbit"
	"github.com/OSSystems/auditmq/config"
	"github.com/OSSystems/auditmq/pkg"
	"github.com/OSSystems/auditmq/storage"
	"github.com/OSSystems/pkg/log"
)

type Publisher interface {
	PublishServiceStatus(service storage.Service) error
	PublishGlobalStat() error
}

type publisher struct {
	storage    storage.Storage
	conn       wabbit.Conn
	ch         wabbit.Channel
	exchange   string
	routingKey string
}

func (p *publisher) getChannel() (wabbit.Channel, error) {
	if p.ch == nil {
		ch, err := p.conn.Channel()
		if err != nil {
			return nil, err
		}

		p.ch = ch
	}

	return p.ch, nil
}

func (p *publisher) publish(report pkg.ReportPayload) error {
	ch, err := p.getChannel()
	if err != nil {
		return err
	}

	message, err := json.Marshal(report)
	if err != nil {
		return err
	}

	return ch.Publish(p.exchange, p.routingKey, message, wabbit.Option{
		"contentType": "application/json",
	})
}

func (p *publisher) PublishServiceStatus(service storage.Service) error {
	serviceName := service.Name()
	status := service.Status()

	log.Debugf("Reporting service %s status (%s)", serviceName, status.String())

	return p.publish(pkg.ReportPayload{
		Service: serviceName,
		Status:  int(status),
	})
}

func (p *publisher) PublishGlobalStat() error {
	serviceStatusList := []storage.Status{}

	for _, service := range p.storage.Services() {
		status := service.Status()
		if status == storage.Unknown {
			continue
		}

		serviceStatusList = append(serviceStatusList, status)
	}

	newStatus := storage.DetermineMinStatus(serviceStatusList)
	if newStatus == storage.Unknown {
		log.Debugf("Ignoring report for all service status")
		return nil
	}

	log.Debugf("Reporting all status with (%s)", newStatus.String())
	return p.publish(pkg.ReportPayload{
		Service: "all",
		Status:  int(newStatus),
	})
}

func New() (Publisher, error) {
	cfg := config.GetConfig()

	store, err := storage.GetStorage()
	if err != nil {
		return nil, err
	}

	return &publisher{
		conn:       cfg.GetAMQP(),
		exchange:   cfg.Exchange,
		routingKey: cfg.ReportRoutingKey,
		storage:    store,
	}, nil
}
