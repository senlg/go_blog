package tests

import "testing"

func TestAdd(t *testing.T) {
	var num = 1 + 1
	if num != 2 {
		t.Errorf("num: %v", num)
	}
}
