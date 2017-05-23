package db

import (
	"testing"

	"news_download/model"
)

func TestAddNews(t *testing.T) {
	InitSqlite("../data.db")
	nl := make([]*model.News, 2)
	n := model.NewNews("testTitle", "testConten??t", "testImgg", 1)
	n2 := model.NewNews("test2Title", "tes2tConten??t", "testIm2gg", 2)
	nl[0] = n
	nl[1] = n2
	id := addNews(nl)
	if id <= 0 {
		t.Fatal("id <=0 ")
	}
	jokes := QueryUnUploadNews()
	if len(jokes) == 0 {
		t.Fatal("len(jokes)==0")
	}
	find := false
	for _, news := range jokes {
		if news.Content == n.Content && news.Img == n.Img && news.Type == n.Type && news.Title == n.Title {
			find = true
		}
	}
	if !find {
		t.Fatal("insert fail")
	}
}
