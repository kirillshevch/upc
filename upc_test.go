package upc

import "testing"

func TestValid(t *testing.T) {
	got := Valid("632737715836")
	want := true

	if got != want {
		t.Errorf("failed")
	}
}

func TestInvalidLength(t *testing.T) {
	want := false

	var invalid_values [2]string

	invalid_values[0] = "63"
	invalid_values[1] = "6327377158366"

	for _, value := range invalid_values {
		got := Valid(value)

		if got != want {
			t.Errorf("Value valid, but should not: %q", value)
		}
	}
}
