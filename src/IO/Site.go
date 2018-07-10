package IO

type Site struct {
	Page          string    `json:"page"`
	PrivacyPolicy int32     `json:"privacypolicy"`
	Id            string    `json:"id"`
	Publisher     Publisher `json:"publisher"`
}

func (s Site) getId() string {
	return s.Id
}

func (s Site) getPublisher() Publisher {
	return s.Publisher
}

func (s Site) getHost() string {
	host := ""
	if s.Page != "" {
		host = s.Page
	}
	return host
}

func (s Site) getURL() string {
	return s.Page
}
