package upc

import "testing"

func TestValid(t *testing.T) {
	got := Valid("632737715836")
	want := true

	if got != want {
		t.Errorf("failed")
	}
}
