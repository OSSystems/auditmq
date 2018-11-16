package main

import (
	"os"
	"os/signal"

	"github.com/NeowayLabs/wabbit"
	"github.com/OSSystems/pkg/log"
	"github.com/OSSystems/auditmq/comparator"
	"github.com/OSSystems/auditmq/config"
	"github.com/OSSystems/auditmq/storage"
)

func main() {
	cfg := config.LoadConfig()
	storage.InitializeStorage(cfg.Data)

	conn := cfg.GetAMQP()

	go func() {
		panic(<-conn.NotifyClose(make(chan wabbit.Error)))
	}()

	consumer, err := NewConsumer()
	if err != nil {
		panic(err)
	}

	log.Info("Listering for new messages, please use CTRL+C to stop")
	consumer.Start()

	log.Info("Initializing comparator")
	comp, err := comparator.New()
	if err != nil {
		panic(err)
	}
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
