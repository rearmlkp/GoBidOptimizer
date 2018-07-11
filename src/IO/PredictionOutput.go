package IO

type PredictionOutput struct {
	PredictedCTR   float32 `json:"predicted_ctr"`
	PredictedCVR   float32 `json:"predicted_cvr"`
	PredictedPrice float32 `json:"predicted_price"`
}
