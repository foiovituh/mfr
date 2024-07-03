package util

import (
	"testing"

	"github.com/foiovituh/mfr/internal/static"
	"github.com/foiovituh/mfr/internal/util"
)

func TestLastCharacter(t *testing.T) {
	result := util.LastCharacter("xyz")
	expected := "z"

	if result != expected {
		t.Errorf(static.LastCharacterTest, result, expected)
	}
}
