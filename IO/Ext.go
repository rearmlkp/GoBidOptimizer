package IO

type Ext struct {
	CampaignID                 string  `json:"campaign_id"`
	AdgroupID                  string  `json:"adgroup_id"`
	Exchange                   string  `json:"exchange"`
	Advertiser                 string  `json:"advertiser"`
	Partner                    string  `json:"partner"`
	Strategy                   string  `json:"strategy"`
	CampaignSessionStartTime   int64   `json:"campaign_session_start_timestamp"`
	CampaignSessionEndTime     int64   `json:"campaign_session_end_timestamp"`
	CampaignStartDate          int32   `json:"campaign_start_date"`
	CampaignEndDate            int32   `json:"campaign_end_date"`
	CampaignSessionTotalBudget float32 `json:"campaign_session_budget"`
	ReferenceCPM               float32
	TargetCPC                  float32
	TargetCPA                  float32
	MultiReferenceCPM          map[string]float32
	RemainSessionBudget        float32
	WinRatioLower              float32
	WinRatioUpper              float32
	BidPriceDeltaLower         float32
	BidPriceDeltaUpper         float32
}

type RtbkitExt struct {
}

type AugmentationList struct {
}

type FrequencyCapRedis struct {
}

type Augmentation struct {
	Data map[string]string `json:"data"`


}
