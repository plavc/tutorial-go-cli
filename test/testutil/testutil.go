package testutil

import (
	"testing"
)

func ExpectInt(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected %d; but got %d", expected, actual)
	}
}

func ExpectString(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("Expected %s; but got %s", expected, actual)
	}
}
