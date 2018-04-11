package model

import "encoding/json"
import "github.com/google/uuid"

type Event struct {
	EventId     uuid.UUID       `json:"id"`
	Actor       Actor           `json:"actor"`
	Type        string          `json:"type"`
	ClientId    uuid.UUID       `json:"client_id"`
	Application string          `json:"application"`
	CreatedAt   string          `json:"created_at"`
	Data        json.RawMessage `json:"data"`
}

type Actor struct {
	ActorType string          `json:"type"`
	Data      json.RawMessage `json:"data"`
}
