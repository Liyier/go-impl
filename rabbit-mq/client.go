package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
)

// Client rabbitmq 的客户端封装
type Client struct {
	Conn   *amqp.Connection
	Config RabbitConfig
}

func NewClient() *Client {
	fmt.Println(config.getConnString())
	return &Client{Config: config}
}

func (c *Client) Connect() error {
	if c.Conn != nil {
		// 连接是否有效
		if c.Ping() {
			return nil
		}
		c.Close()
	}
	// 拨号
	fmt.Println(c.Config.getConnString())
	conn, err := amqp.Dial(c.Config.getConnString())
	if err != nil {
		return err
	}
	c.Conn = conn
	return nil
}

func (c *Client) Ping() bool {
	channel, err := c.Conn.Channel()
	if err != nil {
		return false
	}
	_ = channel.Close()
	return true
}

func (c *Client) Close() {
	if c.Conn != nil {
		c.Conn.Close()
	}
}

func (c *Client) Publish() error {
	channel, err := c.Conn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	err = channel.ExchangeDeclare(
		"exchangeName", // name of the exchange
		"topic",       // type
		true,           // durable
		false,          // delete when complete
		false,          // internal
		false,          // noWait
		nil,            // arguments
	)
	if err != nil {
		return err
	}

	_, _ = channel.QueueDeclare(
		"queueName",   // 队列实名, 队列名和
		true,                        // durable
		false,                       // delete when usused
		false,                       // exclusive
		false,                       // noWait
		nil, // arguments
	)
	_ = channel.QueueBind("queueName", "routingkey.#", "exchangeName", false, nil)

	err = channel.Publish(
		"exchangeName",
		"routingkey.com",
		false,
		false,
		amqp.Publishing{
			Headers:         make(amqp.Table),
			ContentType:     "application/json",
			ContentEncoding: "",
			Body:            []byte("mq msg"),
		})
	if err != nil {
		return err
	}
	return nil
}
