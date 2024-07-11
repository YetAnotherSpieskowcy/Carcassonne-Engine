package position

import (
	"slices"
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
)

func TestAddPositions(t *testing.T) {
	pos1 := NewPosition(3, 5)
	pos2 := NewPosition(1, 2)
	added := pos1.Add(pos2)

	if !(added.X() == 4 && added.Y() == 7) {
		t.Fatalf("expected (%#v,%#v), got (%#v,%#v) instead", 4, 7, added.X(), added.Y())
	}

	pos1 = NewPosition(-3, 5)
	pos2 = NewPosition(1, -4)
	added = pos1.Add(pos2)

	if !(added.X() == -2 && added.Y() == 1) {
		t.Fatalf("expected (%#v,%#v), got (%#v,%#v) instead", -2, 1, added.X(), added.Y())
	}

	pos1 = NewPosition(0, 5)
	pos2 = NewPosition(-1, 0)
	added = pos1.Add(pos2)

	if !(added.X() == -1 && added.Y() == 5) {
		t.Fatalf("expected (%#v,%#v), got (%#v,%#v) instead", -1, 5, added.X(), added.Y())
	}
}

func TestPositionFromSide(t *testing.T) {
	sides := []side.Side{
		side.Top,
		side.TopLeftEdge,
		side.TopRightEdge,

		side.Right,
		side.RightTopEdge,
		side.RightBottomEdge,

		side.Bottom,
		side.BottomRightEdge,
		side.BottomLeftEdge,

		side.Left,
		side.LeftBottomEdge,
		side.LeftTopEdge,
	}

	expected := []Position{
		NewPosition(0, 1),
		NewPosition(0, 1),
		NewPosition(0, 1),

		NewPosition(1, 0),
		NewPosition(1, 0),
		NewPosition(1, 0),

		NewPosition(0, -1),
		NewPosition(0, -1),
		NewPosition(0, -1),

		NewPosition(-1, 0),
		NewPosition(-1, 0),
		NewPosition(-1, 0),
	}

	for i, side := range sides {
		actual := PositionFromSide(side)
		if actual != expected[i] {
			t.Fatalf("PositionFromSide(%#v): expected %#v, got %#v instead", side.String(), expected[i], actual)
		}
	}

}

func TestPositionMarshalTextWithPositiveCoords(t *testing.T) {
	pos := NewPosition(1, 3)
	expected := []byte("1,3")
	actual, err := pos.MarshalText()
	if err != nil {
		t.Fatal(err.Error())
	}
	if !slices.Equal(actual, expected) {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}
}

func TestPositionMarshalTextWithNegativeCoords(t *testing.T) {
	pos := NewPosition(-31, -5)
	expected := []byte("-31,-5")
	actual, err := pos.MarshalText()
	if err != nil {
		t.Fatal(err.Error())
	}
	if !slices.Equal(actual, expected) {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}
}

func TestPositionUnmarshalTextWithPositiveCoords(t *testing.T) {
	text := []byte("1,3")
	expectedX := int16(1)
	expectedY := int16(3)

	actual := Position{}
	err := actual.UnmarshalText(text)
	if err != nil {
		t.Fatal(err.Error())
	}
	if actual.X() != expectedX {
		t.Fatalf("expected %#v, got %#v instead", expectedX, actual)
	}
	if actual.Y() != expectedY {
		t.Fatalf("expected %#v, got %#v instead", expectedY, actual)
	}
}

func TestPositionUnmarshalTextWithNegativeCoords(t *testing.T) {
	text := []byte("-31,-5")
	expectedX := int16(-31)
	expectedY := int16(-5)

	actual := Position{}
	err := actual.UnmarshalText(text)
	if err != nil {
		t.Fatal(err.Error())
	}
	if actual.X() != expectedX {
		t.Fatalf("expected %#v, got %#v instead", expectedX, actual)
	}
	if actual.Y() != expectedY {
		t.Fatalf("expected %#v, got %#v instead", expectedY, actual)
	}
}
