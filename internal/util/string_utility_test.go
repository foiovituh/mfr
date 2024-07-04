package util

import (
	"testing"

	"github.com/foiovituh/mfr/internal/static"
)

func TestLastCharacter(t *testing.T) {
	result := LastCharacter("xyz")
	expected := "z"

	if result != expected {
		t.Errorf(static.LastCharacterTest, result, expected)
	}
}
