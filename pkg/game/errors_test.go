package game

import (
	"testing"
)

func TestInvalidMoveErrorReturnsMsgField(t *testing.T) {
	expected := ErrInvalidPosition.msg
	actual := ErrInvalidPosition.Error()
	if expected != actual {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}
}
