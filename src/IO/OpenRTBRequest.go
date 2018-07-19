package IO

import "github.com/json-iterator/go"

type OpenRTBRequest struct {
	Id                string              `json:"id"`
	Timestamp         int64               `json:"timestamp"`
	TimestampInSecond int32
	Imp               []Imp               `json:"imp"`
	Ext               Ext                 `json:"ext"`
	Device            Device              `json:"device"`
	ImpressionSource  ImpressionSource
	Site              Site
	App               App
	SiteRaw           jsoniter.RawMessage `json:"site,omitempty"`
	AppRaw            jsoniter.RawMessage `json:"app,omitempty"`
	Segments          Segments            `json:"segments"`
}

func (o OpenRTBRequest) SourceString() string {
	return o.Ext.Exchange + ":" + o.ImpressionSource.getPublisher().Id + ":" + o.ImpressionSource.getId()
}
