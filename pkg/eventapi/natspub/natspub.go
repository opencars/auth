package natspub

import (
	"encoding/json"

	"github.com/nats-io/nats.go"

	"github.com/opencars/auth/pkg/eventapi"
)

// Publisher is an implementation ofn pubsub.Publisher interface for NATS.
type Publisher struct {
	conn *nats.Conn
}

// New returns new connection to NATS.
func New(url string, enabled bool) (*Publisher, error) {
	if !enabled {
		return &Publisher{}, nil
	}

	conn, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}

	return &Publisher{
		conn: conn,
	}, nil
}

// Publish publishes event into the queue.
func (p *Publisher) Publish(event *eventapi.Event) error {
	if p.conn == nil {
		return nil
	}

	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	if err := p.conn.Publish("events.auth.new", data); err != nil {
		return err
	}

	return nil
}
