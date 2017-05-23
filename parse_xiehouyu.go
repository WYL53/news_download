package main

import (
	"news_download/model"
)

func parseXiehouyu(m map[string]interface{}) {
	result, ok := m["result"].(map[string]interface{})
	if !ok {
		return
	}
	question := result["question"].(string)
	answer := result["answer"].(string)
	storageXieHouyu(question, answer)
}

func storageXieHouyu(content, answer string) {
	storageNews(model.NewNews("", content+"/n"+answer, "", model.TYPE_XIEHOUYU))
}
