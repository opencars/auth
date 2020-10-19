package eventapi

// Publisher pushlished event into a queue.
type Publisher interface {
	Publish(event *Event) error
}
