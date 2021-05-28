package net

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ifls/gocore/utils/log"
	"strings"
)

//封装gin api
//应该随时可替换为其他web框架
var ginHandlers map[string]gin.HandlerFunc

func init() {
	ginHandlers = map[string]gin.HandlerFunc{}
}

func Serve(addr string) error {
	r := gin.Default()
	for typeUrl, h := range ginHandlers {
		spli := strings.Split(typeUrl, "_")
		typ := spli[0]
		url := spli[1]

		if typ == "post" {
			r.POST(url, h)
		} else if typ == "get" {
			r.GET(url, h)
		} else {
			log.LogErr(fmt.Errorf("find unmatched http method %s", typeUrl))
			continue
		}
	}
	return r.Run(addr)
}

func AddHandler(typeUrl string, h gin.HandlerFunc) {
	ginHandlers[typeUrl] = h
}
