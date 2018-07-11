package IO

type CampaignConfig struct {
	WinRatioUpper            float32
	WinRatioLower            float32
	PriceDeltaLower          float32
	PriceDeltaUpper          float32
	IsPriority               bool
	CampaignSessionTargetCPM float32
	MultiReferenceCPM        map[string]float32
}
