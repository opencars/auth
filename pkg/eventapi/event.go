package eventapi

import (
	"encoding/json"
)

// EventKind is an enumation for kinds of an event.
type EventKind string

const (
	// EventAuthorizationKind represetns kind of authorization event.
	EventAuthorizationKind EventKind = "authorization"
)

// Event represents a message for event API.
type Event struct {
	Kind EventKind       `json:"kind"`
	Data json.RawMessage `json:"data"`
}

// NewEvent returns newly allocated event.
func NewEvent(kind EventKind, v interface{}) (*Event, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	return &Event{
		Kind: kind,
		Data: data,
	}, nil
}
