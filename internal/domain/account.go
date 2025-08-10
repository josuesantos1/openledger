package domain

import "time"

type Account struct {
	ID        string
	ExternalID string
	CreatedAt time.Time
	UpdatedAt time.Time
	OwnerID   string
	Type      string
	Currency  string
	Country   string
}
