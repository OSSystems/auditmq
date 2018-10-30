package config

import (
	"os"

	"github.com/OSSystems/pkg/log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

type DataOptions struct {
	Offset float64 `mapstructure:"offset"`
}

type ServiceData struct {
	Type     string                 `mapstructure:"type"`
	Owner    string                 `mapstructure:"owner"`
	Replicas map[string]DataOptions `mapstructure:"replicas"`
}

type DataFields map[string]ServiceData

type Config struct {
	DSN           string     `mapstructure:"dsn"`
	Data          DataFields `mapstructure:"data"`
	Exchange      string     `mapstructure:"exchange"`
	RoutingKey    string
	ConsumerQueue string `mapstructure:"ConsumerQueue"`
	ConsumerName  string

	amqp *amqp.Connection
}

var config *Config

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func GetConfig() *Config {
	if config == nil {
		panic("Config not initialized")
	}

	return config
}

func LoadConfig() *Config {
	if _, ok := os.LookupEnv("DEBUG"); ok {
		log.SetLevel(logrus.DebugLevel)
	}

	viper.SetConfigType("yml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	panicIfErr(err)

	config = &Config{
		Exchange:      "auditmq_exchange",
		RoutingKey:    "auditmq",
		ConsumerQueue: "audit_in",
		ConsumerName:  "auditmq_worker",
	}

	err = viper.Unmarshal(config)
	panicIfErr(err)

	return config
}

func (c *Config) GetAMQP() *amqp.Connection {
	if c.amqp != nil {
		return c.amqp
	}

	conn, err := amqp.Dial(c.DSN)
	panicIfErr(err)

	ch, err := conn.Channel()
	panicIfErr(err)

	panicIfErr(ch.ExchangeDeclare(c.Exchange, "direct", true, false, false, false, nil))
	_, err = ch.QueueDeclare(c.ConsumerQueue, true, false, false, false, nil)
	panicIfErr(err)
	panicIfErr(ch.QueueBind(c.ConsumerQueue, c.RoutingKey, c.Exchange, false, nil))

	panicIfErr(ch.Close())

	c.amqp = conn
	return conn
}
