package main

import (
	"encoding/json"
	"time"

	"github.com/NeowayLabs/wabbit"
	"github.com/OSSystems/pkg/log"
	"github.com/OSSystems/auditmq/pkg"
	storagePkg "github.com/OSSystems/auditmq/storage"
)

var nackRetryTime = 5

type Handler struct {
	storage storagePkg.Storage
}

func (h *Handler) fillAnotherServices(actualService string, fieldName string) {
	for _, service := range h.storage.Services() {
		if actualService == service.Name() {
			continue
		}

		if !service.HasField(fieldName) {
			continue
		}

		service.DuplicateLast(fieldName)
	}
}

func (h *Handler) Handle(message wabbit.Delivery) {
	servicePayload := &pkg.Payload{}
	err := json.Unmarshal(message.Body(), servicePayload)
	if err != nil {
		log.Error("JSON format error")
		message.Nack(false, true)
		time.Sleep(time.Second * time.Duration(nackRetryTime))
		return
	}

	service, err := h.storage.GetService(servicePayload.Service)
	if err != nil {
		log.Errorf("Service %s does not exists", servicePayload.Service)
		message.Nack(false, true)
		time.Sleep(time.Second * time.Duration(nackRetryTime))
		return
	}

	for field, data := range servicePayload.Data {
		service.Push(field, data)

		again, err := service.GetField(field)
		if err == nil {
			log.Debugf("Service %s.%s - data %v", service.Name(), again.Name, again.ToSlice())
		}

		h.fillAnotherServices(servicePayload.Service, field)
	}

	message.Ack(false)
}

func NewHandler() (*Handler, error) {
	storage, err := storagePkg.GetStorage()
	if err != nil {
		return nil, err
	}

	return &Handler{
		storage: storage,
	}, nil
}
