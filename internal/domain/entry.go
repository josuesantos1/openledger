package domain


import (
	"time"
)

type Entry struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	ConvertedAt time.Time
	ConversationRate int
	AccountID string
	Amount    string
	Status    string
	Side      string
}
