package connection

import (
	"slices"
	"testing"
)

func TestConnectionRotate(t *testing.T) {

	var connect Connection
	connect.Sides = []Side{Right, Bottom}

	var expected Connection
	expected.Sides = []Side{Bottom, Left}

	var rotated = connect.Rotate(1)

	if !slices.Equal(rotated.Sides, expected.Sides) {
		t.Fatalf("%#v should be %#v after rotation", connect.String(), expected.String())
	}
}
func TestConnectionToString(t *testing.T) {
	var connec Connection
	connec.Sides = []Side{Top, Right}

	if connec.String() != Top.String()+" "+Right.String()+" " {
		t.Fatalf("%#v should be %#v ", connec.String(), Top.String()+" "+Right.String()+" ")
	}
}
