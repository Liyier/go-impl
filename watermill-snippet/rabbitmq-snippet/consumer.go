package main

import (
	"context"
	"fmt"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/pkg/amqp"
	"log"
)

const uri = "amqp://guest:guest@localhost:5672/"

func main() {
	config := amqp.NewDurableQueueConfig(uri)
	initSub(&config)
	sub, err := amqp.NewSubscriber(config, watermill.NewStdLogger(true, false))
	if err != nil {
		log.Panicf("there is err: %v", err)
	}
	fmt.Println(sub.IsConnected())
	msgChan, err := sub.Subscribe(context.Background(), "watermill-topic")
	if err != nil {
		log.Panicf("there is err: %v", err)
	}
	msg := <- msgChan
	msg.Nack()
	fmt.Println(msg.Payload)
}

func initSub(config *amqp.Config)  {
	// durable queue
	config.Queue.GenerateName = func(topic string) string {
		return "silo"
	}
}
