package main

import (
	"fmt"
	"time"

	"github.com/OSSystems/pkg/log"
	"github.com/rodrigoapereira/auditmq/config"
	"github.com/streadway/amqp"
)

type Consumer struct {
	cfg  *config.Config
	ch   *amqp.Channel
	done chan error
}

func (c *Consumer) handle(deliveries <-chan amqp.Delivery) {
	for x := range deliveries {
		fmt.Println("RECEIVED ", string(x.Body))
		time.Sleep(10 * time.Second)
		fmt.Println(x.Ack(false))
	}

	log.Debug("Worker stopped")
	c.done <- nil
}

func (c *Consumer) Start() error {
	c.done = make(chan error)
	deliveryChan, err := c.ch.Consume(
		c.cfg.ConsumerQueue,
		c.cfg.ConsumerName,
		false,
		false,
		false,
		false,
		nil,
	)
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

	return &Consumer{
		ch:  ch,
		cfg: cfg,
	}, nil
}
