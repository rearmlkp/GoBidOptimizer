package IO

type BiddingOutput struct {
	Id          string                `json:"id"`
	BidResponse []ImpressionPricepair `json:"bid_response"`
	Success     bool                  `json:"success"`
}

type ImpressionPricepair struct {
	Impid        string  `json:"impid"`
	Price        float32 `json:"price"`
	PredictedCTR float32 `json:"pred_ctr"`
	PredictedCVR float32 `json:"pred_cvr"`
}
