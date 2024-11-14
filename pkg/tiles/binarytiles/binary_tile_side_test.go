package binarytiles

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
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

func TestPositionFromSide(t *testing.T) {
	sides := []BinaryTileSide{
		SideTop,
		SideRight,
		SideBottom,
		SideLeft,
	}
	expectedPositions := []position.Position{
		position.New(0, 1),
		position.New(1, 0),
		position.New(0, -1),
		position.New(-1, 0),
	}

	for i := range sides {
		actualPosition := sides[i].PositionFromSide()
		if actualPosition != expectedPositions[i] {
			t.Fatalf("%v expected: %016b\ngot: %016b", i, expectedPositions[i], actualPosition)
		}
	}
}

func TestSideMirror(t *testing.T) {
	if SideNone.Mirror() != SideNone {
		t.Fatalf("expected %#v, got %#v instead", SideNone, SideNone.Mirror())
	}

	if SideAllOrthogonal.Mirror() != SideAllOrthogonal {
		t.Fatalf("expected %#v, got %#v instead", SideAllOrthogonal, SideAllOrthogonal.Mirror())
	}

	if SideAllDiagonal.Mirror() != SideAllDiagonal {
		t.Fatalf("expected %#v, got %#v instead", SideAllDiagonal, SideAllDiagonal.Mirror())
	}

	if SideRight.Mirror() != SideLeft {
		t.Fatalf("expected %#v, got %#v instead", SideLeft, SideRight.Mirror())
	}

	if SideTopLeftCorner.Mirror() != SideBottomRightCorner {
		t.Fatalf("expected %#v, got %#v instead", SideBottomRightCorner, SideTopLeftCorner.Mirror())
	}

	if (SideTopRightCorner | SideTop).Mirror() != (SideBottomLeftCorner | SideBottom) {
		t.Fatalf("expected %#v, got %#v instead", (SideBottomLeftCorner | SideBottom), (SideTopRightCorner | SideTop).Mirror())
	}
}

func TestHasSide(t *testing.T) {
	type testDataEntry struct {
		side1    BinaryTileSide
		side2    BinaryTileSide
		expected bool
	}
	testData := []testDataEntry{
		{
			side1:    SideRight,
			side2:    SideRight,
			expected: true,
		},
		{
			side1:    SideRight,
			side2:    SideLeft,
			expected: false,
		},
		{
			side1:    SideTopRightCorner | SideTopLeftCorner,
			side2:    SideTopRightCorner,
			expected: true,
		},
		{
			side1:    SideBottomLeftCorner | SideBottomRightCorner,
			side2:    SideTopLeftCorner,
			expected: false,
		},
		{
			side1:    SideAllDiagonal,
			side2:    SideRight,
			expected: false,
		},
		{
			side1:    SideRight | SideTop,
			side2:    SideRight | SideBottom,
			expected: false,
		},
	}

	for i, entry := range testData {
		actual := entry.side1.HasSide(entry.side2)
		if actual != entry.expected {
			t.Fatalf("%v expected: %v\ngot: %v", i, entry.expected, actual)
		}
	}
}

func TestOverlapsSide(t *testing.T) {
	type testDataEntry struct {
		side1    BinaryTileSide
		side2    BinaryTileSide
		expected bool
	}
	testData := []testDataEntry{
		{
			side1:    SideRight,
			side2:    SideRight,
			expected: true,
		},
		{
			side1:    SideRight,
			side2:    SideLeft,
			expected: false,
		},
		{
			side1:    SideTopRightCorner | SideTopLeftCorner,
			side2:    SideTopRightCorner,
			expected: true,
		},
		{
			side1:    SideBottomLeftCorner | SideBottomRightCorner,
			side2:    SideBottomLeftCorner | SideTopLeftCorner,
			expected: true,
		},
		{
			side1:    SideAllDiagonal,
			side2:    SideRight,
			expected: false,
		},
		{
			side1:    SideRight | SideTop,
			side2:    SideRight | SideBottom,
			expected: true,
		},
	}

	for i, entry := range testData {
		actual := entry.side1.OverlapsSide(entry.side2)
		if actual != entry.expected {
			t.Fatalf("%v expected: %v\ngot: %v", i, entry.expected, actual)
		}
	}
}
