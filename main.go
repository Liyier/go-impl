package main

import (
	"fmt"
)

func main() {
	//client := rabbitmq.NewClient()
	//
	//err := client.Connect()
	//if err != nil {
	//	fmt.Println("connect err")
	//	fmt.Println(err)
	//}
	//err = client.Publish()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//time.Sleep(time.Hour)
	m()
}

func m()  {
	var (
		appOpenidList = make(map[string][]string)
		// 推送统计信息
		statistics = make(map[string]int)
	)
	users := []map[string]string{{
		"appid": "ddsd",
		"openid": "dsgre",
	},
	{
		"appid": "dsdsgf",
		"openid": "dsgdfdsgre",
	},
	}
	for _, user := range users {
		statistics[user["appid"]]++
		appOpenidList[user["appid"]] = append(appOpenidList[user["appid"]], user["openid"])
	}
	fmt.Println(statistics)
	fmt.Println(appOpenidList)
}
