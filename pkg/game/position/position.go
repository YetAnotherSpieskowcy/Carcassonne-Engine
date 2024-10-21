package position

import (
	"fmt"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
)

type Position struct {
	x int8
	y int8
}

func New(x int8, y int8) Position {
	if x == 127 || x == -128 || y == 127 || y == -128 {
		// technically not necessary, but many functions depending on finding neighbouring positions
		//   can potentially return invalid results or enter infinite loops due to (under/over)flows
		panic(fmt.Sprintf("Position (%#v, %#v) is outside of allowed range: ([-127, 126], [-127, 126])", x, y))
	}
	return Position{x, y}
}

func (pos Position) X() int8 {
	return pos.x
}

func (pos Position) Y() int8 {
	return pos.y
}

func (pos Position) Add(other Position) Position {
	return New(pos.x+other.x, pos.y+other.y)
}

func (pos Position) Rotate(rotations uint) Position {
	rotations %= 4

	for range rotations {
		pos = New(pos.y, -pos.x)
	}
	return pos
}

/*
Returns position neighbouring (0,0) from the given side.
Valid sides are either a single half-edge side (e.g. side.TopLeftEdge) or a single primary side (e.g. side.Right)
Examples:
position.FromSide(side.Right) ---> (1,0)
position.FromSide(side.BottomLeftEdge) == position.FromSide(side.BottomRightEdge) ---> (0, -1)
*/
func FromSide(checkedSide side.Side) Position {
	primarySides := 0
	for _, otherSide := range side.PrimarySides {
		if checkedSide.OverlapsSide(otherSide) {
			primarySides++
		}
	}

	if primarySides == 0 {
		return New(0, 0)
	} else if primarySides == 1 {
		switch {
		case checkedSide.OverlapsSide(side.Top):
			return New(0, 1)
		case checkedSide.OverlapsSide(side.Right):
			return New(1, 0)
		case checkedSide.OverlapsSide(side.Left):
			return New(-1, 0)
		case checkedSide.OverlapsSide(side.Bottom):
			return New(0, -1)
		}
	}
	panic(fmt.Sprintf("position.FromSide called with more than one primary side. 'side' = %08b", checkedSide))
}

func (pos Position) MarshalText() ([]byte, error) {
	return fmt.Appendf([]byte{}, "%v,%v", pos.x, pos.y), nil
}

func (pos *Position) UnmarshalText(text []byte) error {
	_, err := fmt.Sscanf(string(text), "%v,%v", &pos.x, &pos.y)
	return err
}
