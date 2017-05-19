package main

import (
	"log"
	"os"

	"github.com/astaxie/beego/config"
)

const (
	API_MIYU_URL           = "http://route.showapi.com/151-2"
	API_HISTORY_TODAY_URL  = "http://route.showapi.com/119-42"
	API_TXT_JOKE_URL       = "http://route.showapi.com/341-1"
	API_PIC_JOKE_URL       = "http://route.showapi.com/341-2"
	API_MINGRENMINGYAN_URL = "https://api.hzy.pw/saying/v1/ciba"
	API_XIEHOUYU_URL       = "http://api.avatardata.cn/XieHouYu/Random"
)

var Config config.Configer
var Log *log.Logger

func init() {
	logger := log.New(os.Stdout, "[news_download]", log.Lshortfile|log.LstdFlags)
	Log = logger

	conf, err := config.NewConfig("json", "config.json")
	if err != nil {
		log.Fatal(err)
	}
	Config = conf
}

func main() {
	//	log.Println(Config.DefaultString("avatar_key", "a2e0fc146f174c67a5d81f5b531f0c0d"))
	start()
}
