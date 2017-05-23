package db

import (
	"time"

	"news_download/model"
)

const TABLE_NAME = "news"

//插入数据
func addNews(newsList []*model.News) int {
	createTime := time.Now().Unix()
	stmt, err := dbfd.Prepare("INSERT INTO " + TABLE_NAME + "(title ,content, img,type,create_time) VALUES(?,?,?,?,?)")
	checkErr(err)
	success := 0
	for _, news := range newsList {
		_, err := stmt.Exec(news.Title, news.Content, news.Img, news.Type, createTime)
		if err != nil {
			continue
		}
		success++
	}

	return success
}

//查询数据
func QueryUnUploadNews() []*model.News {
	rows, err := dbfd.Query("SELECT title,content,img,type FROM " + TABLE_NAME + " WHERE upload=0 ORDER BY create_time DESC")
	checkErr(err)

	newsList := make([]*model.News, 0)
	for rows.Next() {
		news := new(model.News)
		err = rows.Scan(&news.Title, &news.Content, &news.Img, &news.Type)
		checkErr(err)
		newsList = append(newsList, news)
	}
	return newsList
}

//更新数据
func UpdateNewsUploadStatus(id int) int64 {
	stmt, err := dbfd.Prepare("UPDATE " + TABLE_NAME + " SET isupload=1 where id=?")
	checkErr(err)

	_, err = stmt.Exec(id)
	checkErr(err)

	return 1
}
