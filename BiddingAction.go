package main

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"fmt"
)

func Status(ctx *fasthttp.RequestCtx)  {
	fmt.Fprint(ctx, "Status ok")
}

func main() {
	router := fasthttprouter.New()
	router.POST("/status", Status)

}
