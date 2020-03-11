package eventapi

type Publisher interface {
	Publish(event *Event) error
}
