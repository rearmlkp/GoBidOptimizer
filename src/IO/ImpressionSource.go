package IO

type ImpressionSource interface {
	getId()
	getPublisher()
	getHost()
	getURL()
}
