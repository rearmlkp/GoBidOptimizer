package main

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"fmt"
	"github.com/json-iterator/go"
	"GoBidOptimizer/src/IO"
	"GoBidOptimizer/src/Strategies"
	"github.com/json-iterator/go/extra"
	"flag"
	"os"
	"log"
	"runtime/pprof"
	"os/signal"
	"syscall"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func Status(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Status ok")
}

func Bidding(ctx *fasthttp.RequestCtx) {

	var (
		strContentType     = []byte("Content-Type")
		strApplicationJSON = []byte("application/json")
	)

	var ORTB IO.OpenRTBRequest
	if err := jsoniter.Unmarshal(ctx.Request.Body(), &ORTB); err != nil {
		fmt.Println(err)
	}

	// Simple processing for all raw
	if ORTB.AppRaw != nil {
		if err := jsoniter.Unmarshal(ORTB.AppRaw, &ORTB.App); err != nil {
			fmt.Println(err)
		} else {
			ORTB.ImpressionSource = ORTB.App
		}
	}
	if ORTB.SiteRaw != nil {
		if err := jsoniter.Unmarshal(ORTB.SiteRaw, &ORTB.Site); err != nil {
			fmt.Println(err)
		} else {
			ORTB.ImpressionSource = ORTB.Site
		}
	}

	ORTB.TimestampInSecond = int32(ORTB.Timestamp / 1000)

	//k, err := jsoniter.MarshalToString(&ORTB)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(k)

	// Optimizing bid
	prediction := BiddingOptimizer(ORTB)
	outputResult, _ := jsoniter.Marshal(&prediction)
	ctx.Response.SetBody(outputResult)
	ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)
	//fmt.Fprintln(ctx, outputResult)
}

func BiddingOptimizer(ORTB IO.OpenRTBRequest) IO.BiddingOutput {
	isPriority := ORTB.Ext.GetCampaignConfig().IsPriority
	//biddingStrategy := ORTB.Ext.Strategy
	var r Strategies.AbstractBiddingStrategy
	if isPriority {
		r = Strategies.IgnoreBiddingStrategy{}
	} else {
		r = Strategies.IgnoreBiddingStrategy{}
	}
	bidResult := r.PredictPrice(ORTB)
	result := IO.BiddingOutput{
		Id: ORTB.Id,
	}
	for k, v := range *bidResult {
		result.BidResponse = append(result.BidResponse, IO.ImpressionPricePair{
			Impid:        k,
			PredictedCTR: v.PredictedCTR,
			PredictedCVR: v.PredictedCVR,
			Price:        v.PredictedPrice,
		})
	}

	return result
}

func cleanup() {
	fmt.Print("Stop profiler")
	pprof.StopCPUProfile()
}

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			fmt.Println(err)
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			fmt.Println(err)
			log.Fatal("could not start CPU profile: ", err)
		}
		fmt.Println("Using profiler")
		defer func() {
			fmt.Print("Stop profiler")
			pprof.StopCPUProfile()
		}()
	}
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		os.Exit(1)
	}()
	extra.RegisterFuzzyDecoders()
	router := fasthttprouter.New()
	router.POST("/status", Status)
	router.POST("/bidding.optimizer.v2/bidding", Bidding)
	fasthttp.ListenAndServe(":8080", router.Handler)
}
