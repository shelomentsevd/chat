package chat

import "github.com/m4rw3r/uuid"

type Chat struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
