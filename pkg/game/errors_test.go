package game

import (
	"testing"
)

func TestInvalidMoveErrorReturnsMsgField(t *testing.T) {
	expected := InvalidPosition.msg
	actual := InvalidPosition.Error()
	if expected != actual {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}
}
