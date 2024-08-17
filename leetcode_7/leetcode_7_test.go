package leetcode7

import "testing"

func TestReverse(t *testing.T) {
	got := Reverse(-123)
	if got != -321 {
		t.Errorf("Got = %d; Want -321", got)
	}
}
