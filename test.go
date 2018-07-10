package main

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"fmt"
	"github.com/json-iterator/go"
)

func Index(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Welcome")
}

type ColorGroup struct {
	ID     int       `json:"id"`
	Name   string    `json:"name"`
	Colors []string  `json:"colors"`
	Test   TestGroup `json:"test"`
}

type TestGroup struct {
	Inside int `json:"inside"`
}

func main() {
	router := fasthttprouter.New()
	router.GET("/", Index)
	val := []byte(`{"id":1,"name":"Reds","colors":["Crimson","Red","Ruby","Maroon"], "test": {"inside": 1}}`)
	//fmt.Println(jsoniter.Get(val, "test", "inside").ToString())
	var k ColorGroup
	err := jsoniter.Unmarshal(val, &k)
	if err == nil {
		fmt.Println(k.Test.Inside)
	} else {
		fmt.Println(err)
	}
	fasthttp.ListenAndServe(":8080", router.Handler)
}
