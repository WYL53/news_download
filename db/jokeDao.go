package db

import (
	"time"

	"news_download/model"
)

//插入数据
func AddJoke(txt, img string) int64 {
	createTime := time.Now().Unix()
	stmt, err := dbfd.Prepare("INSERT INTO joke(content, img, upload,create_time) VALUES(?,?,0,?)")
	checkErr(err)
	res, err := stmt.Exec(txt, img, createTime)
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	return id
}

//查询数据
func QueryUnUploadJokes() []*model.Joke {
	rows, err := dbfd.Query("SELECT content,img FROM joke WHERE upload=0 ORDER BY create_time DESC")
	checkErr(err)

	jokes := make([]*model.Joke, 0)
	for rows.Next() {
		joke := new(model.Joke)
		err = rows.Scan(&joke.Content, &joke.Img)
		checkErr(err)
		jokes = append(jokes, joke)
	}
	return jokes
}

func UpdateJokeUploadStatus(id int) int64 {
	//更新数据
	stmt, err := dbfd.Prepare("UPDATE joke SET isupload=1 where id=?")
	checkErr(err)

	res, err := stmt.Exec(id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)
	return affect
}
