package main

import (
	"fmt"
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
	fmt.Println(content, answer)
}
