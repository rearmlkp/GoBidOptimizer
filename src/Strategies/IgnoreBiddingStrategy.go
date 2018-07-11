package Strategies

import "GoBidOptimiser/src/IO"

type IgnoreBiddingStrategy struct {
}

func (i IgnoreBiddingStrategy) PredictPrice(ortb IO.OpenRTBRequest) map[string]IO.PredictionOutput {
	result := map[string]IO.PredictionOutput{}
	for _, item := range ortb.Imp {
		result[item.Id] = IO.PredictionOutput{PredictedPrice: 0}
	}
	return result
}

func (i IgnoreBiddingStrategy) UpdateHandler() {

}
