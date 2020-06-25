package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)
func main() {
	// 非连接池 方式
	client, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()

	// set
	reply, err := client.Do("set", "key", "value你好")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("reply: %v\n", reply)
	fmt.Println(reply.(string))

	// get
	reply, err = client.Do("get", "key")
	if err != nil {
		fmt.Println(err)
	}

	// 从 redis 返回的数据应该是 []uint8, 即 []byte
	//fmt.Printf("reply: %v\n", reply)
	//fmt.Println(reply.(string))
	fmt.Println(redis.String(reply, err))
}
