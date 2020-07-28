package main

import (
	"fmt"
	"sync"
)

func main() {
	//concurrencyWriteMap()
	concurrencyReadMap()
}

// 测试 map 的 并发写入
// 协程数一多就会引发并发写入的错误
// 使用通用锁即可解决问题，但是效率有所下降
func concurrencyWriteMap()  {
	var (
		m = make(map[string]string)
		wg sync.WaitGroup
		mx sync.Mutex
	)
	for i:=0; i<1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i)
			mx.Lock()
			defer mx.Unlock()
			m[key] = "value"
		}(i)
	}
	wg.Wait()
	fmt.Println(len(m))
}


// 仅仅并发读不会出现问题
func concurrencyReadMap() {
	var (
		m = map[string]string{"key1":"value", "key2":"value"}
		wg sync.WaitGroup
	)
	for i:=0; i<10000000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = m["key1"]
			_ = m["key2"]
		}()
	}
	wg.Wait()
	fmt.Println(len(m))
}
