package main

import (
	"news_download/model"
)

func parseMingRenMingYan(m map[string]interface{}) {
	content := m["cn"].(string)
	contentEn := m["en"].(string)
	pic, ok := m["picSquare"].(string)
	if !ok {
		pic, ok = m["picSmall"].(string)
		if !ok {
			pic, ok = m["pic"].(string)
			if !ok {
				pic = ""
			}
		}
	}
	storageMingRenMingYan(content, contentEn, pic)

}

func storageMingRenMingYan(content, contentEn, pic string) {
	storageNews(model.NewNews("", content+"/n"+contentEn, pic, model.TYPE_MINGRENMINGYAN))
}
