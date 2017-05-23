package main

import (
	"news_download/model"
)

//import(
//	""
//)

func parseMiYu(m map[string]interface{}) {
	showapiResBody, ok := m["showapi_res_body"].(map[string]interface{})
	if !ok {
		return
	}
	pagebean, ok := showapiResBody["pagebean"].(map[string]interface{})
	if !ok {
		return
	}
	miyuList, ok := pagebean["contentlist"].([]interface{})
	if !ok || len(miyuList) == 0 {
		return
	}
	for _, miyu := range miyuList {
		my, ok := miyu.(map[string]interface{})
		if !ok {
			continue
		}
		typeName := my["typeName"].(string)
		answer := my["Answer"].(string)
		title := my["Title"].(string)
		storageMiyu(typeName, answer, title)
	}
}

func storageMiyu(typeName, answer, title string) {
	storageNews(model.NewNews(typeName, title+"/n"+answer, "", model.TYPE_MIYU))
}
