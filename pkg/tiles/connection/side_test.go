package connection

import (
	"testing"
)

func TestSideRotate(t *testing.T) {

	if TOP.Rotate(1) != RIGHT {
		t.Fatalf("got %#v should be %#v after rotation", TOP.Rotate(1), RIGHT)
	}

	if RIGHT.Rotate(1) != BOTTOM {
		t.Fatalf("got %#v should be %#v after rotation", RIGHT.Rotate(1), BOTTOM)
	}

	if BOTTOM.Rotate(1) != LEFT {
		t.Fatalf("got %#v should be %#v after rotation", BOTTOM.Rotate(1), LEFT)
	}

	if LEFT.Rotate(1) != TOP {
		t.Fatalf("got %#v should be %#v after rotation", LEFT.Rotate(1), TOP)
	}
	// corners
	if TOP_LEFT_CORNER.Rotate(1) != TOP_RIGHT_CORNER {
		t.Fatalf("got %#v should be %#v after rotation", TOP_LEFT_CORNER.Rotate(1), TOP_RIGHT_CORNER)
	}

	if TOP_RIGHT_CORNER.Rotate(1) != BOTTOM_RIGHT_CORNER {
		t.Fatalf("got %#v should be %#v after rotation", TOP_RIGHT_CORNER.Rotate(1), BOTTOM_RIGHT_CORNER)
	}

	if BOTTOM_RIGHT_CORNER.Rotate(1) != BOTTOM_LEFT_CORNER {
		t.Fatalf("got %#v should be %#v after rotation", BOTTOM_RIGHT_CORNER.Rotate(1), BOTTOM_LEFT_CORNER)
	}

	if BOTTOM_LEFT_CORNER.Rotate(1) != TOP_LEFT_CORNER {
		t.Fatalf("got %#v should be %#v after rotation", BOTTOM_LEFT_CORNER.Rotate(1), TOP_LEFT_CORNER)
	}

	if CENTER.Rotate(1) != CENTER {
		t.Fatalf("got %#v should be %#v after rotation", CENTER.Rotate(1), CENTER)
	}

	if TOP_LEFT_EDGE.Rotate(1) != RIGHT_TOP_EDGE {
		t.Fatalf("got %#v should be %#v after rotation", TOP_LEFT_EDGE.Rotate(1), RIGHT_TOP_EDGE)
	}

	if TOP_RIGHT_EDGE.Rotate(1) != RIGHT_BOTTOM_EDGE {
		t.Fatalf("got %#v should be %#v after rotation", TOP_RIGHT_EDGE.Rotate(1), RIGHT_BOTTOM_EDGE)
	}

	if RIGHT_TOP_EDGE.Rotate(1) != BOTTOM_RIGHT_EDGE {
		t.Fatalf("got %#v should be %#v after rotation", RIGHT_TOP_EDGE.Rotate(1), BOTTOM_RIGHT_EDGE)
	}

	if RIGHT_BOTTOM_EDGE.Rotate(1) != BOTTOM_LEFT_EDGE {
		t.Fatalf("got %#v should be %#v after rotation", RIGHT_BOTTOM_EDGE.Rotate(1), BOTTOM_LEFT_EDGE)
	}

	if BOTTOM_RIGHT_EDGE.Rotate(1) != LEFT_BOTTOM_EDGE {
		t.Fatalf("got %#v should be %#v after rotation", BOTTOM_RIGHT_EDGE.Rotate(1), LEFT_BOTTOM_EDGE)
	}

	if BOTTOM_LEFT_EDGE.Rotate(1) != LEFT_TOP_EDGE {
		t.Fatalf("got %#v should be %#v after rotation", BOTTOM_LEFT_EDGE.Rotate(1), LEFT_TOP_EDGE)
	}

	if LEFT_BOTTOM_EDGE.Rotate(1) != TOP_LEFT_EDGE {
		t.Fatalf("got %#v should be %#v after rotation", LEFT_BOTTOM_EDGE.Rotate(1), TOP_LEFT_EDGE)
	}

	if LEFT_TOP_EDGE.Rotate(1) != TOP_RIGHT_EDGE {
		t.Fatalf("got %#v should be %#v after rotation", LEFT_TOP_EDGE.Rotate(1), TOP_RIGHT_EDGE)
	}

	if NONE.Rotate(1) != NONE {
		t.Fatalf("got %#v should be %#v after rotation", NONE.Rotate(1), NONE)
	}

	if Side(80).Rotate(1) != NONE {
		t.Fatalf("got %#v should be %#v after rotation", Side(80).Rotate(1), NONE)
	}
}

func TestSideToString(t *testing.T) {
	if TOP.String() != "TOP" {
		t.Fatalf("got %#v should be %#v", TOP.String(), "TOP")
	}

	if RIGHT.String() != "RIGHT" {
		t.Fatalf("got %#v should be %#v", RIGHT.String(), "RIGHT")
	}

	if LEFT.String() != "LEFT" {
		t.Fatalf("got %#v should be %#v", LEFT.String(), "LEFT")
	}

	if BOTTOM.String() != "BOTTOM" {
		t.Fatalf("got %#v should be %#v", BOTTOM.String(), "BOTTOM")
	}

	if TOP_LEFT_CORNER.String() != "TOP_LEFT_CORNER" {
		t.Fatalf("got %#v should be %#v", TOP_LEFT_CORNER.String(), "TOP_LEFT_CORNER")
	}

	if TOP_RIGHT_CORNER.String() != "TOP_RIGHT_CORNER" {
		t.Fatalf("got %#v should be %#v", TOP_RIGHT_CORNER.String(), "TOP_RIGHT_CORNER")
	}

	if BOTTOM_LEFT_CORNER.String() != "BOTTOM_LEFT_CORNER" {
		t.Fatalf("got %#v should be %#v", BOTTOM_LEFT_CORNER.String(), "BOTTOM_LEFT_CORNER")
	}

	if BOTTOM_RIGHT_CORNER.String() != "BOTTOM_RIGHT_CORNER" {
		t.Fatalf("got %#v should be %#v", BOTTOM_RIGHT_CORNER.String(), "BOTTOM_RIGHT_CORNER")
	}

	if CENTER.String() != "CENTER" {
		t.Fatalf("got %#v should be %#v", CENTER.String(), "CENTER")
	}

	if TOP_LEFT_EDGE.String() != "TOP_LEFT_EDGE" {
		t.Fatalf("got %#v should be %#v", TOP_LEFT_EDGE.String(), "TOP_LEFT_EDGE")
	}
	if TOP_RIGHT_EDGE.String() != "TOP_RIGHT_EDGE" {
		t.Fatalf("got %#v should be %#v", TOP_RIGHT_EDGE.String(), "TOP_RIGHT_EDGE")
	}
	if RIGHT_TOP_EDGE.String() != "RIGHT_TOP_EDGE" {
		t.Fatalf("got %#v should be %#v", RIGHT_TOP_EDGE.String(), "RIGHT_TOP_EDGE")
	}
	if RIGHT_BOTTOM_EDGE.String() != "RIGHT_BOTTOM_EDGE" {
		t.Fatalf("got %#v should be %#v", RIGHT_BOTTOM_EDGE.String(), "RIGHT_BOTTOM_EDGE")
	}
	if LEFT_TOP_EDGE.String() != "LEFT_TOP_EDGE" {
		t.Fatalf("got %#v should be %#v", LEFT_TOP_EDGE.String(), "LEFT_TOP_EDGE")
	}
	if LEFT_BOTTOM_EDGE.String() != "LEFT_BOTTOM_EDGE" {
		t.Fatalf("got %#v should be %#v", LEFT_BOTTOM_EDGE.String(), "LEFT_BOTTOM_EDGE")
	}
	if BOTTOM_LEFT_EDGE.String() != "BOTTOM_LEFT_EDGE" {
		t.Fatalf("got %#v should be %#v", BOTTOM_LEFT_EDGE.String(), "BOTTOM_LEFT_EDGE")
	}
	if BOTTOM_RIGHT_EDGE.String() != "BOTTOM_RIGHT_EDGE" {
		t.Fatalf("got %#v should be %#v", BOTTOM_RIGHT_EDGE.String(), "BOTTOM_RIGHT_EDGE")
	}

	if NONE.String() != "NONE" {
		t.Fatalf("got %#v should be %#v", NONE.String(), "NONE")
	}

	if Side(20).String() != "ERROR" {
		t.Fatalf("got %#v should be %#v", Side(20).String(), "ERROR")
	}

}
