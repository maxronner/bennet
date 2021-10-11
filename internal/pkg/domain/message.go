package domain

import "time"

type MessageID string

type Message struct {
	ID       MessageID `json:"id,omitempty"`
	Sent     time.Time `json:"sent,omitempty"`
	Body     string    `json:"body,omitempty"`
	PersonID `json:"person_id,omitempty"`
}
