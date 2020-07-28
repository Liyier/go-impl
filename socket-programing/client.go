package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()

	for {
		// 从标准输入读取数据
		reader := bufio.NewReader(os.Stdin)
		line,err := reader.ReadString('\n')
		if line == "q\n" {
			// 退出
			return
		}
		if err != nil {
			fmt.Println("从标准输入读取字符串错误")
		}
		n, err := conn.Write([]byte(line))

		fmt.Println("发送消息", n)
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(time.Second)
	}
}
