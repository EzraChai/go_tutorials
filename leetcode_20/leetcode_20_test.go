package leetcode20

import "testing"

func TestReverse(t *testing.T) {
	got := isValid("{}()[][}")
	if got != false {
		t.Errorf("Got = %v; Want false", got)
	}
}
