package domain

import "time"

type Transaction struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	Currency  string
	Type      string
	Status    string
	Description string
	ExternalID string
}
