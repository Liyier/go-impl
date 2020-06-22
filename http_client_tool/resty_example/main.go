package main

import (
	"bytes"
	"fmt"
	resty "github.com/go-resty/resty/v2"
	"io/ioutil"
	"unicode/utf8"
)

func main()  {
	//SimpleGetRequest()
	//UploadFile()
	//req()
	s := "控制"
	fmt.Println(utf8.RuneCountInString(s))
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

func req()  {
	data := map[string]interface{}{
		"content": `🎉618年中大促，5重福利重磅来袭

10元定金抵扣650

🎁首发Python学习地图，夏日最爱循环扇，仅限1000份
	
点击底部菜单栏--购买课程	

立即参与活动
	
更多活动进度消息，回复“618”持续关注`,
		"user_sql": "select 'wxe4885224fe1e1c95' as app_id, 'o6EtvvyUnWIDw6ch9fouwBWDBeCk' as openid",
		"title": "给 wusidi的预览",
		"send_type": "mass",
	}
	client := resty.New()
	client.Debug = true
	r := client.R()
	r.SetBody(data).Post("https://dev.pandateacher.com/fast-silo/silo-go-feature-custom/push")
}
