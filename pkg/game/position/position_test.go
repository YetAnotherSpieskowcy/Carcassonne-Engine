package position

import (
	"slices"
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
)

func TestAddPositions(t *testing.T) {
	pos1 := New(3, 5)
	pos2 := New(1, 2)
	added := pos1.Add(pos2)

	if !(added.X() == 4 && added.Y() == 7) {
		t.Fatalf("expected (%#v,%#v), got (%#v,%#v) instead", 4, 7, added.X(), added.Y())
	}

	pos1 = New(-3, 5)
	pos2 = New(1, -4)
	added = pos1.Add(pos2)

	if !(added.X() == -2 && added.Y() == 1) {
		t.Fatalf("expected (%#v,%#v), got (%#v,%#v) instead", -2, 1, added.X(), added.Y())
	}

	pos1 = New(0, 5)
	pos2 = New(-1, 0)
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

		side.None,
	}

	expected := []Position{
		New(0, 1),
		New(0, 1),
		New(0, 1),

		New(1, 0),
		New(1, 0),
		New(1, 0),

		New(0, -1),
		New(0, -1),
		New(0, -1),

		New(-1, 0),
		New(-1, 0),
		New(-1, 0),

		New(0, 0),
	}

	for i, side := range sides {
		actual := FromSide(side)
		if actual != expected[i] {
			t.Fatalf("PositionFromSide(%#v): expected %#v, got %#v instead", side.String(), expected[i], actual)
		}
	}
}

func TestPositionFromSidePanic(t *testing.T) {
	sides := []side.Side{
		side.Top | side.Left,
		side.RightBottomEdge | side.BottomRightEdge,
		side.Top | side.Right | side.Bottom | side.Left,
		side.Top | side.BottomLeftEdge,
	}
	for _, side := range sides {
		// weird nested anonymous function and defer combo
		// I don't know how or why it works, but without nesting "defer func()" in "func()", only a single iteration is tested
		func() {
			defer func() {
				r := recover()
				if r == nil {
					t.Errorf("position.FromSide(%08b) did not panic", side)
				}
			}()
			FromSide(side)
		}()
	}
}

func TestPositionMarshalTextWithPositiveCoords(t *testing.T) {
	pos := New(1, 3)
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
	pos := New(-31, -5)
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
