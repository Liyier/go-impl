package main

import "github.com/go-resty/resty/v2"

func main() {
	c := resty.New()
	c.SetRetryCount(3)
	c.AddRetryCondition(func(response *resty.Response, err error) bool {
		if err != nil {
			return true
		}
		return false
	})
	c.Debug = true
	_, _ = c.R().Get("www.bai")
	_, _ = c.R().Get("www.bai.cc")
	_, _ = c.R().Get("http://www.bai.com")
}
