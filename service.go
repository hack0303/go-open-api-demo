package main

import (
	"log"

	"github.com/valyala/fasthttp"
)

func requestBaidu() {
	url := "http://www.baidu.com"
	status, resp, err := fasthttp.Get(nil, url)
	if err != nil {
		log.Println("调用异常,%s", err.Error())
		return
	}
	if status != fasthttp.StatusOK {
		log.Println("错误状态：%s", status)
	}
	log.Println("%s", string(resp))

}
