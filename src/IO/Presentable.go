package IO

type Presentable interface {
	getW() int32
	getH() int32
}

type Banner struct {
	Id string `json:"id"`
	W  int32  `json:"w"`
	H  int32  `json:"h"`
}

type Native struct {
}

type Video struct {
	W int32 `json:"w"`
	H int32 `json:"h"`
}

func (b Banner) getW() int32 {
	return b.W
}

func (b Banner) getH() int32 {
	return b.H
}

func (n Native) getW() int32 {
	return -1
}

func (n Native) getH() int32 {
	return -1
}

func (v Video) getW() int32 {
	return v.W
}

func (v Video) getH() int32 {
	return v.H
}
