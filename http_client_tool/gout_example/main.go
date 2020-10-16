package main

import "github.com/guonaihong/gout"

func main() {
	// get 可以完成携带 body
	// 可以单个请求开启 debug， 实际需求不大
	gout.GET("http://localhost:8080/set").Debug(true).SetBody("hello, body")
}
