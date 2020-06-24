package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"log"
	"net"
)

func main() {
	pubSub := gochannel.NewGoChannel(gochannel.Config{}, watermill.NewStdLogger(true, true))

	// publish
	go func() {
		for i := 0; i < 10; i++ {
			var payload bytes.Buffer
			payload.WriteString(fmt.Sprintf("helloworld %d", i))
			msg := message.NewMessage(watermill.NewUUID(), payload.Bytes())
			_ = pubSub.Publish("example.topic", msg)
		}
	}()

	// subscribe
	messageChan, err := pubSub.Subscribe(context.Background(), "example.topic")
	if err != nil {
		fmt.Println(err)
	}
	// consume
	go func() {
		for {
			msg := <- messageChan
			fmt.Printf("receive message %s ", msg.Payload)
			msg.Ack()
		}
	}()
	// forever
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal("cannot start a server")
	}
	_, _ = listener.Accept()
}
