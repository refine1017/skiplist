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
