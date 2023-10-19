package unit_test

import "testing"

func TestExample(t *testing.T) {
	hi := 1 + 1
	if hi != 2 {
		t.Error("hi isn't 2")
	}
}
