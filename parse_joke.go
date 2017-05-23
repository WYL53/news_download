package main

import (
	"strings"

	"news_download/model"
)

var lastTxtJokeid string
var lastPicJokeid string

func parseTxtJoke(m map[string]interface{}) {
	showapiResBody, ok := m["showapi_res_body"].(map[string]interface{})
	if !ok {
		return
	}
	contentList, ok := showapiResBody["contentlist"].([]interface{})
	if !ok || len(contentList) == 0 {
		return
	}
	newId := ""
	for _, elem := range contentList {
		joke, ok := elem.(map[string]interface{})
		if !ok {
			continue
		}
		id, ok := joke["id"].(string)
		if ok && strings.Compare(lastTxtJokeid, id) == 0 {
			break
		}
		if len(newId) == 0 {
			newId = id
		}
		text := joke["text"].(string)
		storageTxtJoke(text)
	}
	if len(newId) > 0 && lastTxtJokeid != newId {
		lastTxtJokeid = newId
	}
}

func storageTxtJoke(txt string) {
	storageNews(model.NewNews("", txt, "", model.TYPE_JOKE))
}

func parsePicJoke(m map[string]interface{}) {
	showapiResBody, ok := m["showapi_res_body"].(map[string]interface{})
	if !ok {
		return
	}
	contentList, ok := showapiResBody["contentlist"].([]interface{})
	if !ok || len(contentList) == 0 {
		return
	}
	newId := ""
	for _, elem := range contentList {
		joke, ok := elem.(map[string]interface{})
		if !ok {
			continue
		}
		id, ok := joke["id"].(string)
		if ok && strings.Compare(id, lastPicJokeid) == 0 {
			break
		}
		if len(newId) == 0 {
			newId = id
		}
		title := joke["title"].(string)
		img, ok := joke["img"].(string)
		if !ok {
			img = ""
		}
		storagePicJoke(title, img)
	}
	if len(newId) > 0 && newId != lastPicJokeid {
		lastPicJokeid = newId
	}
}

func storagePicJoke(txt, img string) {
	storageNews(model.NewNews("", txt, img, model.TYPE_JOKE))
}

//过滤字符
func filterQuot(s string) string {
	s = strings.Replace(s, "&amp", "", -1)
	s = strings.Replace(s, "</br>", "", -1)
	return s
}
