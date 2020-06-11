package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	CompareTypeOfWithValueOf()
}

func CompareTypeOfWithValueOf() {
	var s io.Writer = os.Stdout
	// Type 是一个接口 rtype 是具体实现
	t := reflect.TypeOf(s)
	fmt.Printf("kind: %s, name: %s\n", t.Kind(), t.Name())

	// Value 是一个结构体
	v := reflect.ValueOf(s)
	fmt.Printf("kind: %s, name: %s\n", v.Kind(), t.Name())
	// 将 value 转换为 interface{}
	iv := v.Interface()
	// 类型断言
	w, ok := iv.(io.Writer)
	if !ok {
		panic("not writer")
	}
	w.Write([]byte("hello world\n"))
}
