package main

import (
	"context"
	"fmt"
	rabbitmq "impl/rabbit-mq"
)

func main() {
	client := rabbitmq.NewClient()
	if err := client.Connect(); err != nil {
		fmt.Println(err)
	}
	client.Publish()
}


func deferPanic() {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("recover success")
				// recover 返回的 err 是 panic 产生的
				fmt.Println(err)
			}
		}()
		panic("panic self")
	}()

	fmt.Println("main goroutine")
}

func chanClose()  {
	c := make(chan int)
	close(c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func selectChan()  {
	ctx := context.Background()
	//go func() {
	//	for i:=0; i<=3; i++ {
	//		c <- i
	//		fmt.Println("write to chan", i)
	//	}
	//}()
	//close(c)
	select {
	// ctx.Done 在 context 没有结束前返回 nil
		// 值为 nil 的 chan 同样可读
	case <-ctx.Done():
		fmt.Println("read from c")
	}
}


