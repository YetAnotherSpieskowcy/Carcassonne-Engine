package connection

import (
	"slices"
	"testing"
)

func TestConnectionRotate(t *testing.T) {

	var connect Connection
	connect.Sides = []Side{RIGHT, BOTTOM}

	var expected Connection
	expected.Sides = []Side{BOTTOM, LEFT}

	var rotated = connect.Rotate(1)

	if !slices.Equal(rotated.Sides, expected.Sides) {
		t.Fatalf("%#v should be %#v after rotation", connect.String(), expected.String())
	}
}
func TestConnectionToString(t *testing.T) {
	var connec Connection
	connec.Sides = []Side{TOP, RIGHT}

	if connec.String() != TOP.String()+" "+RIGHT.String()+" " {
		t.Fatalf("%#v should be %#v ", connec.String(), TOP.String()+" "+RIGHT.String()+" ")
	}
}
