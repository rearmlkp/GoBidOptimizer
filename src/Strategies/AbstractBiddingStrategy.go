package Strategies

import "GoBidOptimiser/src/IO"

type AbstractBiddingStrategy interface {
	PredictPrice(request IO.OpenRTBRequest) map[string]IO.PredictionOutput
	UpdateHandler()
}
