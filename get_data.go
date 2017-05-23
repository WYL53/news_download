package main

import (
	"fmt"
	"math"
	"time"
	//	"bytes"
	"encoding/json"
	//	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	//	"runtime"
	"strings"
	//	"sync"
	//	"net"
)

func start() {
	//	wg := &sync.WaitGroup{}
	//	wg.Add(5)
	go func() {
		//		defer wg.Done()
		//		defer wg.Done()
		//		defer wg.Done()
		YYDataFunc(API_PIC_JOKE_URL, parsePicJoke)
		YYDataFunc(API_MIYU_URL, parseMiYu)
		YYDataFunc(API_TXT_JOKE_URL, parseTxtJoke)
	}()

	go func(url string) {
		//		defer wg.Done()
		getMingRenMingYan(url)
	}(API_MINGRENMINGYAN_URL)

	go func(url string) {
		//		defer wg.Done()
		getXieHouYu(url, parseXiehouyu)
	}(API_XIEHOUYU_URL)

	//	runtime.Gosched()
	//	wg.Wait()
	//	Log.Println("start eixt")
}

//易源 数据
func getYYData(url string, parseHandle func(map[string]interface{})) {
	defer func() {
		if x := recover(); x != nil {
			Log.Println(x)
		}
	}()
	payload := strings.NewReader(fmt.Sprintf("showapi_sign=%s&showapi_appid=%s",
		Config.String("yy_showapi_sign"), Config.String("yy_showapi_appid")))
	resp := postRequest(url, payload)
	showapiResCode, ok := resp["showapi_res_code"].(float64)
	if ok && isEqualFloat(showapiResCode, 0) {
		parseHandle(resp)
	}
}

func getXieHouYu(url string, parseHandle func(map[string]interface{})) {
	defer func() {
		if x := recover(); x != nil {
			Log.Println(x)
		}
	}()
	payload := strings.NewReader(fmt.Sprintf("key=%s", Config.String("avatar_key")))
	resp := postRequest(url, payload)
	showapiResCode, ok := resp["error_code"].(float64)
	if ok && isEqualFloat(showapiResCode, 0) {
		parseHandle(resp)
	} else {
		reason, ok := resp["reason"].(string)
		if ok {
			Log.Println("get xie hou yu fail:" + reason)
		}
	}
}

func getMingRenMingYan(url string) {
	defer func() {
		if x := recover(); x != nil {
			Log.Println(x)
		}
	}()
	resp := getRequest(url)
	parseMingRenMingYan(resp)
}

func postRequest(url string, reqBody io.Reader) map[string]interface{} {
	checkErr := func(err error) {
		if err != nil {
			panic(err)
		}
	}
	client := http.Client{
		Timeout: time.Duration(time.Second * time.Duration(10)),
	}

	request, err := http.NewRequest(http.MethodPost, url, reqBody)
	checkErr(err)
	request.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, err := client.Do(request)
	checkErr(err)
	defer res.Body.Close()
	response, err := ioutil.ReadAll(res.Body)
	checkErr(err)
	respMap := make(map[string]interface{})
	err = json.Unmarshal(response, &respMap)
	checkErr(err)
	return respMap
}

func isEqualFloat(a, b float64) bool {
	return math.Abs(a-b) < 0.5
}

func getRequest(url string) map[string]interface{} {
	checkErr := func(err error) {
		if err != nil {
			panic(err)
		}
	}
	client := http.Client{
		Timeout: time.Duration(time.Second * time.Duration(10)),
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	checkErr(err)
	resp, err := client.Do(req)
	checkErr(err)
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	checkErr(err)
	respMap := make(map[string]interface{})
	err = json.Unmarshal(content, &respMap)
	checkErr(err)
	return respMap
}
