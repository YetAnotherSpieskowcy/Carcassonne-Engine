package connection

import (
	"testing"
)

func TestSideRotate(t *testing.T) { //nolint:gocyclo // simply testing all states

	if Top.Rotate(1) != Right {
		t.Fatalf("got %#v should be %#v after rotation", Top.Rotate(1), Right)
	}

	if Right.Rotate(1) != Bottom {
		t.Fatalf("got %#v should be %#v after rotation", Right.Rotate(1), Bottom)
	}

	if Bottom.Rotate(1) != Left {
		t.Fatalf("got %#v should be %#v after rotation", Bottom.Rotate(1), Left)
	}

	if Left.Rotate(1) != Top {
		t.Fatalf("got %#v should be %#v after rotation", Left.Rotate(1), Top)
	}
	// corners
	if TopLeftCorner.Rotate(1) != TopRightCorner {
		t.Fatalf("got %#v should be %#v after rotation", TopLeftCorner.Rotate(1), TopRightCorner)
	}

	if TopRightCorner.Rotate(1) != BottomRightCorner {
		t.Fatalf("got %#v should be %#v after rotation", TopRightCorner.Rotate(1), BottomRightCorner)
	}

	if BottomRightCorner.Rotate(1) != BottomLeftCorner {
		t.Fatalf("got %#v should be %#v after rotation", BottomRightCorner.Rotate(1), BottomLeftCorner)
	}

	if BottomLeftCorner.Rotate(1) != TopLeftCorner {
		t.Fatalf("got %#v should be %#v after rotation", BottomLeftCorner.Rotate(1), TopLeftCorner)
	}

	if Center.Rotate(1) != Center {
		t.Fatalf("got %#v should be %#v after rotation", Center.Rotate(1), Center)
	}

	if TopLeftEdge.Rotate(1) != RightTopEdge {
		t.Fatalf("got %#v should be %#v after rotation", TopLeftEdge.Rotate(1), RightTopEdge)
	}

	if TopRightEdge.Rotate(1) != RightBottomEdge {
		t.Fatalf("got %#v should be %#v after rotation", TopRightEdge.Rotate(1), RightBottomEdge)
	}

	if RightTopEdge.Rotate(1) != BottomRightEdge {
		t.Fatalf("got %#v should be %#v after rotation", RightTopEdge.Rotate(1), BottomRightEdge)
	}

	if RightBottomEdge.Rotate(1) != BottomLeftEdge {
		t.Fatalf("got %#v should be %#v after rotation", RightBottomEdge.Rotate(1), BottomLeftEdge)
	}

	if BottomRightEdge.Rotate(1) != LeftBottomEdge {
		t.Fatalf("got %#v should be %#v after rotation", BottomRightEdge.Rotate(1), LeftBottomEdge)
	}

	if BottomLeftEdge.Rotate(1) != LeftTopEdge {
		t.Fatalf("got %#v should be %#v after rotation", BottomLeftEdge.Rotate(1), LeftTopEdge)
	}

	if LeftBottomEdge.Rotate(1) != TopLeftEdge {
		t.Fatalf("got %#v should be %#v after rotation", LeftBottomEdge.Rotate(1), TopLeftEdge)
	}

	if LeftTopEdge.Rotate(1) != TopRightEdge {
		t.Fatalf("got %#v should be %#v after rotation", LeftTopEdge.Rotate(1), TopRightEdge)
	}

	if None.Rotate(1) != None {
		t.Fatalf("got %#v should be %#v after rotation", None.Rotate(1), None)
	}

	if Side(80).Rotate(1) != None {
		t.Fatalf("got %#v should be %#v after rotation", Side(80).Rotate(1), None)
	}
}

func TestSideToString(t *testing.T) { //nolint:gocyclo // simply testing all states
	if Top.String() != "TOP" {
		t.Fatalf("got %#v should be %#v", Top.String(), "TOP")
	}

	if Right.String() != "RIGHT" {
		t.Fatalf("got %#v should be %#v", Right.String(), "RIGHT")
	}

	if Left.String() != "LEFT" {
		t.Fatalf("got %#v should be %#v", Left.String(), "LEFT")
	}

	if Bottom.String() != "BOTTOM" {
		t.Fatalf("got %#v should be %#v", Bottom.String(), "BOTTOM")
	}

	if TopLeftCorner.String() != "TOP_LEFT_CORNER" {
		t.Fatalf("got %#v should be %#v", TopLeftCorner.String(), "TOP_LEFT_CORNER")
	}

	if TopRightCorner.String() != "TOP_RIGHT_CORNER" {
		t.Fatalf("got %#v should be %#v", TopRightCorner.String(), "TOP_RIGHT_CORNER")
	}

	if BottomLeftCorner.String() != "BOTTOM_LEFT_CORNER" {
		t.Fatalf("got %#v should be %#v", BottomLeftCorner.String(), "BOTTOM_LEFT_CORNER")
	}

	if BottomRightCorner.String() != "BOTTOM_RIGHT_CORNER" {
		t.Fatalf("got %#v should be %#v", BottomRightCorner.String(), "BOTTOM_RIGHT_CORNER")
	}

	if Center.String() != "CENTER" {
		t.Fatalf("got %#v should be %#v", Center.String(), "CENTER")
	}

	if TopLeftEdge.String() != "TOP_LEFT_EDGE" {
		t.Fatalf("got %#v should be %#v", TopLeftEdge.String(), "TOP_LEFT_EDGE")
	}
	if TopRightEdge.String() != "TOP_RIGHT_EDGE" {
		t.Fatalf("got %#v should be %#v", TopRightEdge.String(), "TOP_RIGHT_EDGE")
	}
	if RightTopEdge.String() != "RIGHT_TOP_EDGE" {
		t.Fatalf("got %#v should be %#v", RightTopEdge.String(), "RIGHT_TOP_EDGE")
	}
	if RightBottomEdge.String() != "RIGHT_BOTTOM_EDGE" {
		t.Fatalf("got %#v should be %#v", RightBottomEdge.String(), "RIGHT_BOTTOM_EDGE")
	}
	if LeftTopEdge.String() != "LEFT_TOP_EDGE" {
		t.Fatalf("got %#v should be %#v", LeftTopEdge.String(), "LEFT_TOP_EDGE")
	}
	if LeftBottomEdge.String() != "LEFT_BOTTOM_EDGE" {
		t.Fatalf("got %#v should be %#v", LeftBottomEdge.String(), "LEFT_BOTTOM_EDGE")
	}
	if BottomLeftEdge.String() != "BOTTOM_LEFT_EDGE" {
		t.Fatalf("got %#v should be %#v", BottomLeftEdge.String(), "BOTTOM_LEFT_EDGE")
	}
	if BottomRightEdge.String() != "BOTTOM_RIGHT_EDGE" {
		t.Fatalf("got %#v should be %#v", BottomRightEdge.String(), "BOTTOM_RIGHT_EDGE")
	}

	if None.String() != "NONE" {
		t.Fatalf("got %#v should be %#v", None.String(), "NONE")
	}

	if Side(20).String() != "ERROR" {
		t.Fatalf("got %#v should be %#v", Side(20).String(), "ERROR")
	}

}
