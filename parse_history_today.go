package main

import (
	"news_download/model"
)

func parseHistoryTody(m map[string]interface{}) {
	showapiResBody, ok := m["showapi_res_body"].(map[string]interface{})
	if !ok {
		return
	}
	list, ok := showapiResBody["list"].([]interface{})
	if !ok || len(list) == 0 {
		return
	}
	//	fmt.Println(list)
	for _, one := range list {
		history, ok := one.(map[string]interface{})
		if !ok {
			continue
		}
		title := history["title"].(string)
		year := history["year"].(string)
		img, ok := history["img"].(string)
		if !ok {
			img = ""
		}
		storageHistoryToday(title, year, img)
	}
}

func storageHistoryToday(title, year, img string) {
	storageNews(model.NewNews(year, title, img, model.TYPE_HISTORY_TODAY))
	//	nl := make([]s)
}
