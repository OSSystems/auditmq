package main

import (
	"os"
	"os/signal"

	"github.com/OSSystems/pkg/log"
	"github.com/rodrigoapereira/auditmq/config"
	"github.com/rodrigoapereira/auditmq/matcher"
	"github.com/streadway/amqp"
)

func main() {
	cfg := config.LoadConfig()
	matcher.InitializeWithData(cfg.Data)

	conn := cfg.GetAMQP()

	go func() {
		panic(<-conn.NotifyClose(make(chan *amqp.Error)))
	}()

	consumer, err := NewConsumer()
	if err != nil {
		panic(err)
	}

	log.Info("Listering for new messages, please use CTRL+C to stop")
	consumer.Start()

	log.Info("Initializing comparator")
	comp := matcher.NewComparator()
	comp.Start()

	signalCh := make(chan os.Signal)
	signal.Notify(signalCh, os.Interrupt)

	<-signalCh
	log.Info("Comparator shutdown")
	comp.Shutdown()

	log.Info("Worker shutdown")
	err = consumer.Shutdown()
	if err != nil {
		panic(err)
	}

	os.Exit(0)
}
