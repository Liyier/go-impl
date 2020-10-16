package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	net.ListenUDP()
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("err: %s", err)
	}
	defer listener.Close()

	for {
		// 该方法阻塞
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("客户端地址：%s\n", conn.RemoteAddr())  // 客户端端口随机
		//c.LocalAddr() // 本地地址
		//c.RemoteAddr() // 远程地址， 可用于做白名单
		go process(conn)
	}

}

func process(conn net.Conn)  {
	defer func() {
		fmt.Println("关闭连接")
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		//  buf[:n] 表示实际读到的内容
		fmt.Printf("读取信息: %s", string(buf[:n]))
		fmt.Printf("读取信息: %s", string(buf))
	}
}
