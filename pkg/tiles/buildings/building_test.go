package buildings

import (
	"testing"
)

func TestBuildingToString(t *testing.T) {

	if None.String() != "NONE_BUILDING" {
		t.Fatalf("got %#v should be %#v", None.String(), "NONE_BUILDING")
	}

	if Monastery.String() != "MONASTERY" {
		t.Fatalf("got %#v should be %#v", Monastery.String(), "MONASTERY")
	}

	if Bulding(100).String() != "ERROR" {
		t.Fatalf("got %#v should be %#v", Bulding(100).String(), "ERROR")
	}
}