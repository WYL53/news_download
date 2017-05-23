package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"news_download/model"
)

var (
	dbfd     *sql.DB
	newsChan chan []*model.News = make(chan []*model.News, 256)
	started  bool
	stopChan chan int = make(chan int, 1)
)

func InitSqlite(dbName string) {
	var err error
	dbfd, err = sql.Open("sqlite3", dbName)
	checkErr(err)
}

func StartStorageLoop() {
	if started {
		return
	}
	started = true
	for started {
		select {
		case news := <-newsChan:
			addNews(news)
		case <-stopChan:
			started = false
		}
	}
	close(newsChan)
	close(stopChan)
	dbfd.Close()
}

func AddNews(news []*model.News) {
	if started {
		newsChan <- news
	}
}

func StopStorageLoop() {
	if started {
		stopChan <- 1
	}
}

//func AddNews()

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
