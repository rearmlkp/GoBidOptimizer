package IO

import "github.com/json-iterator/go"

type Imp struct {
	Id          string              `json:"id"`
	BidFloor    float32             `json:"bidfloor"`
	BidFloorCur string              `json:"bidfloorcur"`
	Presentable Presentable
	Banner      Banner
	Video       Video
	BannerRaw   jsoniter.RawMessage `json:"banner,omitempty"`
	VideoRaw    jsoniter.RawMessage `json:"video,omitempty"`
	Ext         ImpExt              `json:"ext"`
	Position    int32               `json:"position"`
	Format      string
	Type        string
}

func (m Imp) getWidth() int32 {
	return m.Presentable.getW()
}

func (m Imp) getHeight() int32 {
	return m.Presentable.getH()
}

type ImpExt struct {
	Viewability   int32   `json:"viewability"`
	HistoricalCTR float32 `json:"click_through_rate"`
	BidFloor      float32 `json:"bid_floor"`
}
