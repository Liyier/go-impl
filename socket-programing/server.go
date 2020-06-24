package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("err: %s", err)
	}
	defer listener.Close()

	c, err := listener.Accept()
	if err != nil {
		fmt.Println(err)
	}
	for {
		n, err := bufio.NewReader(c).ReadString('\n')
		 //c.Read(buffer)
		 fmt.Println("收到消息")
		fmt.Println(n)
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(time.Second)
	}


}
