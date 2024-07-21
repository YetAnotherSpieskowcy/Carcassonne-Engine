package side

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

	if LeftTopEdge.Rotate(4) != LeftTopEdge {
		t.Fatalf("got %#v should be %#v after rotation", LeftTopEdge.Rotate(4), LeftTopEdge)
	}

	if None.Rotate(1) != None {
		t.Fatalf("got %#v should be %#v after rotation", None.Rotate(1), None)
	}
}

func TestSideRotateReturnsSideRotatedTwice(t *testing.T) {
	expected := Bottom
	actual := Top.Rotate(2)
	if expected != actual {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}
}

func TestSideMirror(t *testing.T) {
	if None.Mirror() != None {
		t.Fatalf("expected %#v, got %#v instead", None, None.Mirror())
	}

	if All.Mirror() != All {
		t.Fatalf("expected %#v, got %#v instead", All, All.Mirror())
	}

	if Right.Mirror() != Left {
		t.Fatalf("expected %#v, got %#v instead", Left, Right.Mirror())
	}

	if TopLeftEdge.Mirror() != BottomLeftEdge {
		t.Fatalf("expected %#v, got %#v instead", BottomLeftEdge, TopLeftEdge.Mirror())
	}

	if (TopLeftEdge | Right).Mirror() != (BottomLeftEdge | Left) {
		t.Fatalf("expected %#v, got %#v instead", (BottomLeftEdge | Left), (TopLeftEdge | Right).Mirror())
	}
}

func TestSideFlipSides(t *testing.T) {
	if None.FlipSides() != None {
		t.Fatalf("expected %#v, got %#v instead", None, None.FlipSides())
	}

	if All.FlipSides() != All {
		t.Fatalf("expected %#v, got %#v instead", All, All.FlipSides())
	}

	if Right.FlipSides() != Right {
		t.Fatalf("expected %#v, got %#v instead", Left, Right.FlipSides())
	}

	if TopLeftEdge.FlipSides() != TopRightEdge {
		t.Fatalf("expected %#v, got %#v instead", TopRightEdge, TopLeftEdge.FlipSides())
	}

	if (TopLeftEdge | Right).FlipSides() != (TopRightEdge | Right) {
		t.Fatalf("expected %#v, got %#v instead", (TopRightEdge | Right), (TopLeftEdge | Right).FlipSides())
	}
}

func TestSideFlipCorners(t *testing.T) {
	if None.FlipCorners() != None {
		t.Fatalf("expected %#v, got %#v instead", None, None.FlipCorners())
	}

	if All.FlipCorners() != All {
		t.Fatalf("expected %#v, got %#v instead", All, All.FlipCorners())
	}

	if Right.FlipCorners() != (TopRightEdge | BottomRightEdge) {
		t.Fatalf("expected %#v, got %#v instead", (TopRightEdge | BottomRightEdge), Right.FlipCorners())
	}

	if TopLeftEdge.FlipCorners() != LeftTopEdge {
		t.Fatalf("expected %#v, got %#v instead", BottomLeftEdge, TopLeftEdge.FlipCorners())
	}

	if (TopLeftEdge | Right).FlipCorners() != (LeftTopEdge | TopRightEdge | BottomRightEdge) {
		t.Fatalf("expected %#v, got %#v instead", (LeftTopEdge | TopRightEdge | BottomRightEdge), (TopLeftEdge | Right).FlipCorners())
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
}
