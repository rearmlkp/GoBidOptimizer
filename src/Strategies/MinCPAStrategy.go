package Strategies

import (
	"GoBidOptimizer/src/IO"
	"math"
	"GoBidOptimizer/src/Configurations"
	"fmt"
	"GoBidOptimizer/src/Inventory"
	"GoBidOptimizer/src/Managers"
	"strconv"
)

type MinCPAStrategy struct {
	Alpha float32
}

func (i MinCPAStrategy) PredictPrice(ortb IO.OpenRTBRequest) *map[string]IO.PredictionOutput {
	result := make(map[string]IO.PredictionOutput)
	//conf := Configurations.GetConfig()

	currentInventory := Inventory.GetInventory(ortb)
	fmt.Println(currentInventory.Id)
	// TODO: Call Analytic in another thread [ 90 -> 97 MinCPAStrategy ]

	// TODO: Call promise CTR [ 99 -> 109 MinCPAStrategy ]

	conversionProbability := getConversionProbability(ortb)
	if conversionProbability < 0 {
		return bidHistoricalCPM(ortb)
	}

	bidRequestTime := ortb.TimestampInSecond
	currentIndex := int32(math.Ceil(float64(bidRequestTime-currentInventory.StartTime) / float64(currentInventory.TimeInterval)))
	if currentInventory.LastIntervalIndex < 0 {
		currentInventory.LastIntervalIndex = currentIndex
		Inventory.UpdateMap(currentInventory.Id, currentInventory)
	} else {

	}

	return &result
}

func sigmoid(x float32) float32 {
	if x > 0.01 {
		return float32(2 / (1 + math.Exp(float64(-1*100*(x-0.01)))))
	} else {
		return float32(2 * (1/(1+math.Exp(float64(-1*500*x))) - 0.5))
	}
}

func getConversionProbability(ortb IO.OpenRTBRequest) float32 {
	exceptionList := [3]string{"adx", "smaato", "taboola"}
	exchange := ortb.Ext.Exchange
	if exchange == "" {
		exchange = ""
	}
	knorexId := ""
	containFlag := false
	for _, item := range exceptionList {
		if knorexId == item {
			containFlag = true
			break
		}

	}
	if containFlag && ortb.Ext.UserIds.Xchg != "" {
		knorexId = ortb.Ext.UserIds.Xchg
	} else {
		knorexId = ortb.Ext.UserIds.Prov
	}

	redisKey := Inventory.GenerateConversionProbKey(exchange, knorexId)
	data, err := Managers.GetRedisClient().HGetAll(redisKey).Result()
	if data == nil || err != nil || len(data) == 0 {
		if err != nil {
			fmt.Println(err)
		}
		return -1
	}

	advertiserId := ortb.Ext.Advertiser
	if advertiserId == "" {
		advertiserId = "na"
	}
	if val, ok := data[advertiserId]; ok {
		if i, err := strconv.ParseFloat(val, 32); err == nil {
			return float32(i)
		} else {
			fmt.Println(err)
		}
	} else {
		for _, v := range data {
			if i, err := strconv.ParseFloat(v, 32); err == nil {
				return Configurations.GetConfig().ConversionDiscountFactor * float32(i)
			} else {
				fmt.Println(err)
			}
		}
	}

	return -1
}

func bidHistoricalCPM(ortb IO.OpenRTBRequest) *map[string]IO.PredictionOutput {
	result := make(map[string]IO.PredictionOutput)
	return &result
}
