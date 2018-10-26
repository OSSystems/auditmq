package main

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
	"github.com/rodrigoapereira/auditmq/config"
	"github.com/rodrigoapereira/auditmq/pkg"
	"github.com/streadway/amqp"
)

type Handler struct {
	db *gorm.DB
}

func (h *Handler) Handle(message amqp.Delivery) {
	servicePayload := &pkg.Payload{}
	err := json.Unmarshal(message.Body, servicePayload)
	if err != nil {
		message.Nack(false, true)
		return
	}

	_, err = json.Marshal(servicePayload.Data)
	if err != nil {
		message.Nack(false, true)
		return
	}

	message.Ack(false)
}

func NewHandler() *Handler {
	cfg := config.GetConfig()
	db := cfg.GetDB()

	return &Handler{
		db: db,
	}
}
