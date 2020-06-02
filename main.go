package main

import (
	"fmt"
	rabbitmq "impl/rabbit-mq"
	"time"
)

func main() {
	client := rabbitmq.NewClient()

	err := client.Connect()
	if err != nil {
		fmt.Println("connect err")
		fmt.Println(err)
	}
	err = client.Publish()
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Hour)
}
