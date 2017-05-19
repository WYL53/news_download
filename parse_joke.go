package main

import (
	"fmt"
	"strings"
)

func parseTxtJoke(m map[string]interface{}) {
	showapiResBody, ok := m["showapi_res_body"].(map[string]interface{})
	if !ok {
		return
	}
	contentList, ok := showapiResBody["contentlist"].([]interface{})
	if !ok || len(contentList) == 0 {
		return
	}
	for i, elem := range contentList {
		joke, ok := elem.(map[string]interface{})
		if !ok {
			continue
		}
		text := joke["text"].(string)
		storageTxtJoke(i, text)
	}
}

func storageTxtJoke(i int, txt string) {
	fmt.Println(i, txt)
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
	for _, elem := range contentList {
		joke, ok := elem.(map[string]interface{})
		if !ok {
			continue
		}
		title := joke["title"].(string)
		img, ok := joke["img"].(string)
		if !ok {
			img = ""
		}
		storagePicJoke(title, img)
	}
}

func storagePicJoke(txt, img string) {
	fmt.Println(filterQuot(txt), img)
}

//过滤字符
func filterQuot(s string) string {
	s = strings.Replace(s, "&amp", "", -1)
	s = strings.Replace(s, "</br>", "", -1)
	return s
}
