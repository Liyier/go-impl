package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

type SnsType int32

const (
	UnknownSnsType   SnsType = 0
	WeChatSnsType    SnsType = 1
	WeChatMPSnsType  SnsType = 2
	WeChatCPSnsType  SnsType = 3
	WeChatMIPSnsType SnsType = 4
	DouYinSnsType    SnsType = 5
	MobileSnsType    SnsType = 6
)

//const SQL = "u.id,u.nickname,u.headimgurl,u.sex,u.language,u.country,u.city,u.province,u.mobile,wuu.unionid,wuu.wx_app_id,wuu.openid,wuu.subscribe,wuu.subscribe_tag,wuu.subscribe_scene,wuu.subscribe_time,wuu.qr_scene,wuu.qr_scene_str"


// 用户同步状态
type FosUser struct {
	Id             int64   `xorm:"pk autoincr bigint"`
	Unionid        string  `xorm:"unionid"`
	Appid          int64   `xorm:"wx_app_id"`
	Openid         string  `xorm:"openid"`
	Nickname       string  `xorm:"nickname"`
	HeadImgUrl     string  `xorm:"headimgurl"`
	Sex            int32   `xorm:"sex"`
	Language       string  `xorm:"language"`
	Country        string  `xorm:"country"`
	City           string  `xorm:"city"`
	Province       string  `xorm:"province"`
	Mobile         string  `xorm:"mobile"`
	Subscribe      int32   `xorm:"subscribe"`
	TagIdList      []int32 `xorm:"subscribe_tag"`
	SubscribeScene string  `xorm:"subscribe_scene"`
	SubscribeTime  int64   `xorm:"subscribe_time"`
	QrScene        int32   `xorm:"qr_scene"`
	QrSceneStr     string  `xorm:"qr_scene_str"`
}

type cookieSession struct {
	md5Salt string
	domain  string
	lives   int
	secure  bool
}


func main() {
	cs := cookieSession{}
	a := `{
  "md5Salt": "fWG^o07n#39nNV9K",
  "domain": "",
  "lives": "604800",
  "secure": "true"
}`
	var cfg map[string]string
	json.Unmarshal([]byte(a), &cfg)
	cs.md5Salt = cfg["md5Salt"]
	cs.domain = cfg["domain"]
	if v, ok := cfg["lives"]; ok {
		i, _ := strconv.Atoi(v)
		cs.lives = i * 10
	}
	if v, ok := cfg["secure"]; ok {
		b, _ := strconv.ParseBool(v)
		cs.secure = b
	}
	fmt.Println(cs)
}
	//json.Unmarshal([]byte(a), &cs)
