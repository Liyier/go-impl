package main

import (
	"encoding/json"
	"fmt"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
)

const amqpURI = "amqp://guest:guest@localhost:5672/"

func main() {
	config := amqp.NewNonDurableQueueConfig(amqpURI)
	config.Exchange.GenerateName = func(topic string) string {
		return "test-events"
	}
	config.Exchange.Type = "topic"
	config.Exchange.Durable = true
	//fmt.Println("exchange type", config.Exchange.Type)
	publisher, err := amqp.NewPublisher(config, watermill.NewStdLogger(true, false))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(publisher.IsConnected())
	type UserMergeEvent struct {
		MergedUserId int64 `json:"merged_user_id"` // 被合并的用户 id
		MergeToUserId int64 `json:"merge_to_user_id"` // 合并至用户 id
	}
	event := UserMergeEvent{MergedUserId: 2, MergeToUserId: 1}
	data, _ := json.Marshal(event)
	for i:=0;i<3;i++ {
		msg := message.NewMessage(watermill.NewUUID(), data)
		err = publisher.Publish("silo.daemon.backend.mass", msg)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(publisher.IsConnected())
	}
}

