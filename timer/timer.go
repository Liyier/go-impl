package main

import (
	"fmt"
	"sync"
	"time"
)

func main()  {
	t := time.NewTimer(time.Second)

	select {
	case c := <-t.C:
		fmt.Println(c.Format("2006-01-02 15:04:05"))
	}

	var wg sync.WaitGroup

	wg.Add(1)
	time.AfterFunc(time.Second, func() {
		// 此函数异步执行，
		defer wg.Done()
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	})

	wg.Wait()
}
