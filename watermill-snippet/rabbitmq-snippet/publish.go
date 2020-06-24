package main

import (
	"bytes"
	"fmt"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
)

const amqpURI = "amqp://guest:guest@localhost:5672/"

func main() {
	config := amqp.NewNonDurableQueueConfig(amqpURI)
	config.Exchange.GenerateName = func(topic string) string {
		return "micro-events"
	}
	publisher, err := amqp.NewPublisher(config, watermill.NewStdLogger(true, false))
	if err != nil {
		fmt.Println(err)
		return
	}

	var payload bytes.Buffer
	payload.WriteString(fmt.Sprintf("helloworld"))
	msg := message.NewMessage(watermill.NewUUID(), payload.Bytes())
	err = publisher.Publish("topic", msg)
	if err != nil {
		fmt.Println(err)
	}

}

