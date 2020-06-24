package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
	}
	for {
		n, err := conn.Write([]byte("hello world\n"))

		fmt.Println("发送消息", n)
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(time.Second)
	}
}
