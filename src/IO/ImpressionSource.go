package IO

type ImpressionSource interface {
	getId() string
	getPublisher() Publisher
	getHost() string
	getURL() string
}
