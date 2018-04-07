package message

import (
	"time"

	"github.com/m4rw3r/uuid"
)

type Message struct {
	ID        uuid.UUID `json:"id" sql:"unique_index"`
	ChatID    uuid.UUID `json:"chat_id"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
