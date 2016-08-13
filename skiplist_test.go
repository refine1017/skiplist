package skiplist

import (
	"testing"
)

func TestNewSkipList(t *testing.T) {
	skiplist := NewSkipList()

	if skiplist == nil {
		t.Errorf("skiplist is nil")
	}
}

func TestNewSkipListNode(t *testing.T) {
	node := NewSkipListNode(1, "key", "data")

	if node == nil {
		t.Errorf("node is nil")
	}
}
