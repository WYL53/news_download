package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	//	"runtime"
	"sync"
	"time"

	"github.com/astaxie/beego/config"

	"news_download/db"
	"news_download/model"
)

const (
	API_MIYU_URL           = "http://route.showapi.com/151-2"
	API_HISTORY_TODAY_URL  = "http://route.showapi.com/119-42"
	API_TXT_JOKE_URL       = "http://route.showapi.com/341-1"
	API_PIC_JOKE_URL       = "http://route.showapi.com/341-2"
	API_MINGRENMINGYAN_URL = "https://api.hzy.pw/saying/v1/ciba"
	API_XIEHOUYU_URL       = "http://api.avatardata.cn/XieHouYu/Random"
)

var (
	Config   config.Configer
	Log      *log.Logger
	newsChan = make(chan *model.News, 100)
)

func init() {
	logger := log.New(os.Stdout, "[news_download]", log.Lshortfile|log.LstdFlags)
	Log = logger

	conf, err := config.NewConfig("json", "config.json")
	if err != nil {
		log.Fatal(err)
	}
	Config = conf

	db.InitSqlite(conf.DefaultString("database_name", "data.db"))
}

var YYDataFunc = func(url string, parseHandle func(map[string]interface{})) {
	getYYData(url, parseHandle)
}

func main() {
	sigChan := make(chan os.Signal, 1)
	defer close(sigChan)
	signal.Notify(sigChan, syscall.SIGKILL, syscall.SIGINT, syscall.SIGTERM)
	go db.StartStorageLoop()
	defer db.StopStorageLoop()
	go storageLoop()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		loop := true
		minutTtick := time.NewTicker(time.Minute)
		defer minutTtick.Stop()
		dayTick := time.NewTicker(time.Hour * time.Duration(24))
		defer dayTick.Stop()

		go YYDataFunc(API_HISTORY_TODAY_URL, parseHistoryTody)
		for loop {
			select {
			case <-minutTtick.C:
				Log.Println("start get data")
				start()
			case <-dayTick.C:
				Log.Println("start get history day data")
				go YYDataFunc(API_HISTORY_TODAY_URL, parseHistoryTody)
			case <-sigChan:
				loop = false
			}
		}
	}()
	wg.Wait()
	fmt.Println("exit success.")
}

func storageLoop() {
	tick := time.NewTicker(time.Second * time.Duration(5))
	for {
		select {
		case <-tick.C:
			n, ok := <-newsChan
			if ok {
				nl := []*model.News{n}
				timer := time.NewTimer(time.Second)
				timeout := false
				for !timeout {
					select {
					case <-timer.C:
						timeout = true
					case news := <-newsChan:
						nl = append(nl, news)
					}
				}
				if len(nl) > 0 {
					db.AddNews(nl)
				}

			}
		}
	}
}

func storageNews(news *model.News) {
	newsChan <- news
}
