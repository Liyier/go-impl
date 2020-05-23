package main

import (
	"fmt"
	"math"
	"sync"
)

var (
	dataChan   = make(chan int, 10000)
	// 使用通道存储结果而不是切片或者数组， 因为通道是并发安全的
	resultChan = make(chan int, 10000)
	wg         = sync.WaitGroup{}
)

// 通过协程的方式， 求一个 1-10000 中哪些数是素数
func main() {

	go putNum(10000)
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go consume()
	}
	wg.Wait()
	fmt.Printf("总计 %d\n", len(resultChan))
}

func isPrime(number int) bool {
	if number < 2 {
		return false
	}
	squareRoot := math.Sqrt(float64(number))
	for i := 2; i <= int(squareRoot); i++ {
		// 能够被整除
		if number%i == 0 {
			return false
		}
	}
	return true
}

// 将数据放入通道
func putNum(scope int) {
	for i := 0; i < scope; i++ {
		dataChan <- i
	}
	close(dataChan)
}

func consume() {
	defer wg.Done()
	for {
		number, ok := <-dataChan
		if !ok {
			return
		}
		if isPrime(number) {
			fmt.Println("素数", number)
			resultChan <- number
		}
	}

}
