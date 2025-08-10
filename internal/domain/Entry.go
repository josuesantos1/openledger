package domain

type Entry struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	AccountID string
	Amount    string
	Status    string
	Side      string
}
