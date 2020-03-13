package eventapi

import (
	"encoding/json"
)

type EventKind string

const (
	EventAuthorizationKind EventKind = "authorization"
)

type Event struct {
	Kind EventKind       `json:"kind"`
	Data json.RawMessage `json:"data"`
}

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
