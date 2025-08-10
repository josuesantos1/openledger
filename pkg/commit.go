package pkg

import (
	"fmt"
	"crypto/sha256"
)

type CommitSystem interface {
	Apply() error
}

type Commit struct {
	ID      string
	Event   string
	Payload any
	CheckSum string
}

func NewCommit(id string, event string, payload any) *Commit {
	return &Commit{
		ID:      id,
		Event:   event,
		Payload: payload,
	}
}

func (c *Commit) Apply() Commit {
	c.CheckSum = fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%s%s%v", c.ID, c.Event, c.Payload))))
	return *c
}
