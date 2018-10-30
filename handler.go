package main

import (
	"encoding/json"
	"time"

	"github.com/OSSystems/pkg/log"
	"github.com/rodrigoapereira/auditmq/pkg"
	"github.com/streadway/amqp"
)

type Handler struct{}

func (h *Handler) Handle(message amqp.Delivery) {
	servicePayload := &pkg.Payload{}
	err := json.Unmarshal(message.Body, servicePayload)
	if err != nil {
		log.Error("JSON format error")
		message.Nack(false, true)
		time.Sleep(5 * time.Second)
		return
	}

	message.Ack(false)
}

func NewHandler() *Handler {
	return &Handler{}
}
