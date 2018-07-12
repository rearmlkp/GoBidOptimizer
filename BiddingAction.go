package main

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"fmt"
	"github.com/json-iterator/go"
	"GoBidOptimizer/src/IO"
	"GoBidOptimizer/src/Strategies"
	"github.com/json-iterator/go/extra"
)

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
		}
	}
	if ORTB.SiteRaw != nil {
		if err := jsoniter.Unmarshal(ORTB.SiteRaw, &ORTB.Site); err != nil {
			fmt.Println(err)
		}
	}

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
	for k, v := range bidResult {
		result.BidResponse = append(result.BidResponse, IO.ImpressionPricepair{
			Impid:        k,
			PredictedCTR: v.PredictedCTR,
			PredictedCVR: v.PredictedCVR,
			Price:        v.PredictedPrice,
		})
	}

	return result
}

func main() {
	extra.RegisterFuzzyDecoders()
	router := fasthttprouter.New()
	router.POST("/status", Status)
	router.POST("/bidding.optimizer.v2/bidding", Bidding)
	fasthttp.ListenAndServe(":8080", router.Handler)
}
