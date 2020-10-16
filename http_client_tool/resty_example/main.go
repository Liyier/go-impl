package main

import (
	"bytes"
	"fmt"
	resty "github.com/go-resty/resty/v2"
	"io/ioutil"
)

func main()  {
	client := resty.New()
	// 开启 debug
	client.Debug = true
	req := client.R()
	req.SetCookie()
	resp, err := req.Get("http://localhost:8080/test")
	if err != nil {
		fmt.Println(resp.Cookies())
	}
}

func SimpleGetRequest() {
	client := resty.New()
	// 开启 debug
	client.Debug = true
	req := client.R()
	// 添加 query 参数
	req.SetQueryParam("page", "10")
	// 添加 token
	req.SetAuthToken("token")
	// 添加 header 传递 json
	req.Header.Add("Content-Type", "Application/json")
	req.SetBody(map[string]string{"key": "value"})
	// body 为空， resty get 请求不能带 body
	fmt.Println(req.Body)
	resp, err := req.Get("http://localhost:8080/")
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println(resp.StatusCode())
	fmt.Println(resp.Body()) // []byte
	fmt.Println(resp.RawBody()) // io.ReadCloser
}

func UploadFile()  {
	client := resty.New()
	client.Debug = true
	req := client.R()
	file, err := ioutil.ReadFile("/Users/liyi/go/src/gee/gee.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("--------分割线线-------")
	fmt.Println(string(file))
	req.SetFileReader("file", "test.go", bytes.NewReader(file)).Post("http://localhost:8080/upload")
}
