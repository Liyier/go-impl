package rabbitmq

import (
	"fmt"
	"github.com/caarlos0/env"
)

var config RabbitConfig

func init()  {
	if err := env.Parse(&config); err != nil {
		fmt.Println(err)
	}
}

// rabbitmq 连接配置
type RabbitConfig struct {
	Username    string `env:"RABBITMQ_USER" envDefault:"guest"`
	Password    string `env:"RABBITMQ_PASSWORD" envDefault:"guest"`
	VirtualHost string `env:"RABBITMQ_VIRTUAL_HOST" envDefault:""`
	Host        string `env:"RABBITMQ_HOST" envDefault:"localhost"`
	Port        int    `env:"RABBIT_PORT" envDefault:"5672"`
}

func (conf *RabbitConfig) getConnString() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d/%s",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.VirtualHost,
	)
}
