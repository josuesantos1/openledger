package domain

import (
	"time"
)

type Client struct {
	ID        string
	ExternalID string
	CreatedAt time.Time
	UpdatedAt time.Time
}

