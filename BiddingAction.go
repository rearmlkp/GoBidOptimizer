package main

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"fmt"
	"github.com/json-iterator/go"
	"./src/IO"
	"github.com/json-iterator/go/extra"
)

func Status(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Status ok")
}

func Bidding(ctx *fasthttp.RequestCtx) {
	extra.RegisterFuzzyDecoders()
	var ORTB IO.OpenRTBRequest
	err := jsoniter.Unmarshal(ctx.Request.Body(), &ORTB)
	if err != nil {
		fmt.Println(err)
	}

	// Simple processing for all raw
	if ORTB.AppRaw != nil {
		err = jsoniter.Unmarshal(ORTB.AppRaw, &ORTB.App)
		if err != nil {
			fmt.Println(err)
		}
	}
	if ORTB.SiteRaw != nil {
		err = jsoniter.Unmarshal(ORTB.SiteRaw, &ORTB.Site)
		if err != nil {
			fmt.Println(err)
		}
	}

	k, err := jsoniter.MarshalToString(&ORTB)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(k)
	fmt.Fprintln(ctx, "It work...?")
}

func main() {
	router := fasthttprouter.New()
	router.POST("/status", Status)
	router.POST("/bidding.optimizer.v2/bidding", Bidding)
	fasthttp.ListenAndServe(":8080", router.Handler)
}
