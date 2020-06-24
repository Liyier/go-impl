package main

import (
	"fmt"
	"github.com/ThreeDotsLabs/watermill-amqp/pkg/amqp"
)

type MqConfig struct {
	amqp.Config
	Username string `env:"RABBIT_USERNAME" envDefault:"guest"`
	Password string `env:"RABBIT_PASSWORD" envDefault:"guest"`
	Host string  `env:"RABBIT_HOST" envDefault:"localhost"`
}

func GetAmqpURI()  {
	fmt.Sprintf("%")
}
