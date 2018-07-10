package IO

type OpenRTBRequest struct {
	Id string `json:"id"`
	Timestamp int64 `json:"timestamp"`
	TimestampInSecond int32
	Imp []Imp `json:"imp"`

}