package util

import (
	"testing"
)

func TestLastCharacter(t *testing.T) {
	result := LastCharacter("xyz")
	expected := "z"

	if result != expected {
		t.Errorf("LastCharacter(\"xyd\") -> result: %s; expected: %s",
			result, expected)
	}
}
