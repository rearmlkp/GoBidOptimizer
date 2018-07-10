package IO

type Device struct {
	Ua       string `json:"ua"`
	Ip       string `json:"ip"`
	Language string `json:"language"`
	Country  string
	City     string
	Region   string
	Device   string `json:"device"`
	Browser  string `json:"browser"`
	Os       string `json:"os"`
	Geo      Geo    `json:"geo"`
}

type Geo struct {
	Country string `json:"country"`
	City    string `json:"city"`
	Region  string `json:"region"`
}
