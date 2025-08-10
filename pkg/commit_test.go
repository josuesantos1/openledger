package pkg

import (
	"testing"
	"fmt"
	"crypto/sha256"
	"reflect"
)

func TestNewCommit(t *testing.T) {
	id := "1"
	event := "created"
	payload := map[string]interface{}{ "foo": "bar" }
	commit := NewCommit(id, event, payload)
	if commit.ID != id {
		t.Errorf("expected ID %s, got %s", id, commit.ID)
	}
	if commit.Event != event {
		t.Errorf("expected Event %s, got %s", event, commit.Event)
	}
	if !reflect.DeepEqual(commit.Payload, payload) {
		t.Errorf("expected Payload %v, got %v", payload, commit.Payload)
	}
	if commit.CheckSum != "" {
		t.Errorf("expected CheckSum to be empty, got %s", commit.CheckSum)
	}
}

func TestCommit_Apply(t *testing.T) {
	id := "2"
	event := "updated"
	payload := "some data"
	commit := NewCommit(id, event, payload)
	applied := commit.Apply()
	expectedSum := fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%s%s%v", id, event, payload))))
	if applied.CheckSum != expectedSum {
		t.Errorf("expected CheckSum %s, got %s", expectedSum, applied.CheckSum)
	}
	if commit.CheckSum != expectedSum {
		t.Errorf("expected original Commit CheckSum to be %s, got %s", expectedSum, commit.CheckSum)
	}
	if applied.ID != id || applied.Event != event || applied.Payload != payload {
		t.Errorf("applied Commit fields do not match original")
	}
}
