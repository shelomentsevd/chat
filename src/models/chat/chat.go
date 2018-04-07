package chat

import (
	"github.com/m4rw3r/uuid"
)

type Chat struct {
	ID   uuid.UUID `json:"id" sql:"unique_index"`
	Name string    `json:"name"`
}
