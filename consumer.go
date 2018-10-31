package main

import (
	"github.com/NeowayLabs/wabbit"
	"github.com/OSSystems/pkg/log"
	"github.com/rodrigoapereira/auditmq/config"
)

type ConsumerHandler interface {
	Handle(wabbit.Delivery)
}

type Consumer struct {
	Handler ConsumerHandler
	cfg     *config.Config
	ch      wabbit.Channel
	done    chan error
}

func (c *Consumer) handle(deliveries <-chan wabbit.Delivery) {
	for message := range deliveries {
		c.Handler.Handle(message)
	}

	log.Debug("Worker stopped")
	c.done <- nil
}

func (c *Consumer) Start() error {
	c.done = make(chan error)
	deliveryChan, err := c.ch.Consume(c.cfg.ConsumerQueue, c.cfg.ConsumerName, nil)
	if err != nil {
		return err
	}

	go c.handle(deliveryChan)

	return nil
}

func (c *Consumer) Shutdown() error {
	log.Info("Waiting for handler stops")
	err := c.ch.Cancel(c.cfg.ConsumerName, false)
	if err != nil {
		return err
	}

	<-c.done

	log.Info("Closing channel")
	err = c.ch.Close()
	if err != nil {
		return err
	}

	return nil
}

func NewConsumer() (*Consumer, error) {
	cfg := config.GetConfig()
	amqp := cfg.GetAMQP()
	ch, err := amqp.Channel()
	if err != nil {
		return nil, err
	}

	handler, err := NewHandler()
	if err != nil {
		return nil, err
	}

	return &Consumer{
		Handler: handler,
		ch:      ch,
		cfg:     cfg,
	}, nil
}
