package pkg

import (
	"github.com/cockroachdb/pebble"
)

type Storage interface {
	Save(commit Commit) error
	Load(id string) (Commit, error)
	Start() error
	Close() error
	DB() *pebble.DB
}

type Storage struct {
	db *pebble.DB
}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) Start() error {
	db, err := pebble.Open("openledger", &pebble.Options{})
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (s *Storage) DB() *pebble.DB {
	return s.db
}

func (s *Storage) Save(commit Commit) error {
	return s.db.Set([]byte(commit.ID), commit.Data, pebble.Sync)
}

func (s *Storage) Load(id string) (Commit, error) {
	var commit Commit
	err := s.db.Get([]byte(id), &commit.Data)
	if err != nil {
		return Commit{}, fmt.Errorf("commit not found")
	}
	return commit, nil
}

func (s *Storage) Close() error {
	return s.db.Close()
}
