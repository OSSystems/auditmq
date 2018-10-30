package main

import (
	"encoding/json"
	"time"

	"github.com/OSSystems/pkg/log"
	"github.com/rodrigoapereira/auditmq/matcher"
	"github.com/rodrigoapereira/auditmq/pkg"
	"github.com/rodrigoapereira/auditmq/storage"
	"github.com/streadway/amqp"
)

type Handler struct{}

func (h *Handler) fillAnotherServices(actualService string, field string) {
	for srv, serviceOpt := range matcher.Get().Services {
		if srv == actualService {
			continue
		}

		if !serviceOpt.FieldExist(field) {
			log.Debugf("Field %s not configured for service %s", field, srv)
			continue
		}

		serviceStorage := storage.GetStorageData(srv)
		buffer := serviceStorage.GetBuffer(field)
		buffer.CopyLastValue()
	}
}

func (h *Handler) Handle(message amqp.Delivery) {
	servicePayload := &pkg.Payload{}
	err := json.Unmarshal(message.Body, servicePayload)
	if err != nil {
		log.Error("JSON format error")
		message.Nack(false, true)
		time.Sleep(5 * time.Second)
		return
	}

	if !matcher.Get().ServiceExist(servicePayload.Service) {
		log.Error("Service does not exists")
		message.Nack(false, true)
		time.Sleep(5 * time.Second)
		return
	}

	serviceStorage := storage.GetStorageData(servicePayload.Service)
	for field, data := range servicePayload.Data {
		if !matcher.Get().FieldExist(field) {
			log.Debugf("Field %s not present in configuration", field)
			continue
		}

		buffer := serviceStorage.GetBuffer(field)
		buffer.Push(data)

		h.fillAnotherServices(servicePayload.Service, field)
	}

	message.Ack(false)
}

func NewHandler() *Handler {
	return &Handler{}
}
