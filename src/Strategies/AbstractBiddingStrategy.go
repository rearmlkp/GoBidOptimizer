package Strategies

import "GoBidOptimizer/src/IO"

type AbstractBiddingStrategy interface {
	PredictPrice(request IO.OpenRTBRequest) *map[string]IO.PredictionOutput
	UpdateHandler()
}
