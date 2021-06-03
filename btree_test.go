package gobtree

import (
	"testing"
)

func TestSomefunc(t *testing.T) {
	if got := somefunc(); 2 != got {
		t.Errorf("want 2 but got %v\n", got)
	}
}
