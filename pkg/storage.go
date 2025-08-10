package pkg

import (
	"fmt"
	"github.com/cockroachdb/pebble"
	"log"
)

type StorageSystem interface {
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

	s.db = db

	return nil
}

func (s *Storage) DB() *pebble.DB {
	return s.db
}

func (s *Storage) Save(key string, data []byte) error {
	return s.db.Set([]byte(key), data, pebble.Sync)
}

func (s *Storage) Load(id string) ([]byte, error) {
	value, closer, err := s.db.Get([]byte(id))
	if err != nil {
		return nil, fmt.Errorf("data not found")
	}
	defer closer.Close()
	return value, nil
}

func (s *Storage) Close() error {
	return s.db.Close()
}
