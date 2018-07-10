package IO

type App struct {
	Bundle    string    `json:"bundle"`
	Id        string    `json:"id"`
	Publisher Publisher `json:"publisher"`
}

func (a App) getId() string {
	return a.Id
}

func (a App) getPublisher() Publisher {
	return a.Publisher
}

func (a App) getHost() string {
	return a.Bundle
}

func (a App) getURL() string {
	return ""
}
