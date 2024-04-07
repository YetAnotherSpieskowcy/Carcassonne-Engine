package buildings

import (
	"testing"
)

func TestBuildingToString(t *testing.T) {

	if NONE_BULDING.String() != "NONE_BUILDING" {
		t.Fatalf("got %#v should be %#v", NONE_BULDING.String(), "NONE_BUILDING")
	}

	if MONASTERY.String() != "MONASTERY" {
		t.Fatalf("got %#v should be %#v", MONASTERY.String(), "MONASTERY")
	}

	if Bulding(100).String() != "ERROR" {
		t.Fatalf("got %#v should be %#v", Bulding(100).String(), "ERROR")
	}
}
