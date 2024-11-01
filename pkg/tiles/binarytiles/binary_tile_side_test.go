package binarytiles

import (
	"testing"
)

func TestCornersToSides(t *testing.T) {
	sides := []BinaryTileSide{
		SideTopLeftCorner,
		SideBottomRightCorner | SideBottomLeftCorner,
		SideBottomRightCorner,
		SideBottomRightCorner | SideTopRightCorner | SideTopLeftCorner,
		SideTopRightCorner | SideBottomRightCorner | SideBottomLeftCorner | SideTopLeftCorner,
	}
	expected := []BinaryTileSide{
		SideTop | SideLeft,
		SideBottom | SideRight | SideLeft,
		SideBottom | SideRight,
		SideBottom | SideRight | SideTop | SideLeft,
		SideBottom | SideRight | SideTop | SideLeft,
	}

	for i := range sides {
		actual := sides[i].CornersToSides()
		if actual != expected[i] {
			t.Fatalf("expected: %016b\ngot: %016b", expected[i], actual)
		}
	}
}

func TestSidesToCorners(t *testing.T) {
	sides := []BinaryTileSide{
		SideTop,
		SideBottom | SideRight,
		SideBottom,
		SideRight | SideTop | SideLeft,
		SideBottom | SideRight | SideTop | SideLeft,
	}
	expected := []BinaryTileSide{
		SideTopLeftCorner | SideTopRightCorner,
		SideBottomRightCorner | SideBottomLeftCorner | SideTopRightCorner,
		SideBottomRightCorner | SideBottomLeftCorner,
		SideTopRightCorner | SideBottomRightCorner | SideBottomLeftCorner | SideTopLeftCorner,
		SideTopRightCorner | SideBottomRightCorner | SideBottomLeftCorner | SideTopLeftCorner,
	}

	for i := range sides {
		actual := sides[i].SidesToCorners()
		if actual != expected[i] {
			t.Fatalf("expected: %016b\ngot: %016b", expected[i], actual)
		}
	}
}

func TestCornerFromSide(t *testing.T) {
	corners := []BinaryTileSide{
		SideTopLeftCorner,
		SideTopLeftCorner,
		SideTopLeftCorner | SideTopRightCorner,
		SideTopLeftCorner | SideTopRightCorner,
		SideTopLeftCorner | SideTopRightCorner,
		SideTopLeftCorner,
		SideTopLeftCorner | SideTopRightCorner,
	}
	directions := []BinaryTileSide{
		SideTop,
		SideLeft,
		SideTop,
		SideRight,
		SideLeft,
		SideRight,
		SideBottom,
	}

	expected := []BinaryTileSide{
		SideBottomLeftCorner,
		SideTopRightCorner,
		SideBottomRightCorner | SideBottomLeftCorner,
		SideTopLeftCorner,
		SideTopRightCorner,
		SideNone,
		SideNone,
	}

	for i := range corners {
		actual := CornerFromSide(corners[i], directions[i])
		if actual != expected[i] {
			t.Fatalf("%v expected: %016b\ngot: %016b", i, expected[i], actual)
		}
	}
}
