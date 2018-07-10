package IO

import "github.com/json-iterator/go"

type Ext struct {
	CampaignID                 string             `json:"campaign_id"`
	AdgroupID                  string             `json:"adgroup_id"`
	Exchange                   string             `json:"exchange"`
	Advertiser                 string             `json:"advertiser"`
	Partner                    string             `json:"partner"`
	Strategy                   string             `json:"strategy"`
	CampaignSessionStartTime   int64              `json:"campaign_session_start_timestamp"`
	CampaignSessionEndTime     int64              `json:"campaign_session_end_timestamp"`
	CampaignStartDate          int32              `json:"campaign_start_date"`
	CampaignEndDate            int32              `json:"campaign_end_date"`
	CampaignSessionTotalBudget float32            `json:"campaign_session_budget"`
	ReferenceCPM               float32
	TargetCPC                  float32
	TargetCPA                  float32
	MultiReferenceCPM          map[string]float32 `json:"multi_reference_cpm_rate"`
	RemainSessionBudget        float32
	WinRatioLower              float32
	WinRatioUpper              float32
	BidPriceDeltaLower         float32
	BidPriceDeltaUpper         float32
	RtbkitExt                  RtbkitExt          `json:"rtbkit"`
	UserIds                    UserIds            `json:"userIds"`
}

type RtbkitExt struct {
	AugmentationList AugmentationList `json:"augmentationList"`
}

type AugmentationList struct {
	FrequencyCapRedis []FrequencyCapRedis `json:"frequency-cap-redis"`
}

type FrequencyCapRedis struct {
	Account      []string     `json:"account"`
	Augmentation Augmentation `json:"augmentation"`
}

type Augmentation struct {
	Data    map[string]string
	DataRaw jsoniter.RawMessage `json:"data,omitempty"`
	Tags    []string            `json:"tags"`
}

type UserIds struct {
	Prov string `json:"prov"`
	Xchg string `json:"xchg"`
}
