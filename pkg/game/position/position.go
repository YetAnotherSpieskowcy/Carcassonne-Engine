package position

import (
	"fmt"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
)

type Position struct {
	// int8 would be fine for base game (72 tiles) but let's be a bit more generous
	x int16
	y int16
}

func New(x int16, y int16) Position {
	return Position{x, y}
}

func (pos Position) X() int16 {
	return pos.x
}

func (pos Position) Y() int16 {
	return pos.y
}

func (pos Position) Add(other Position) Position {
	return New(pos.x+other.x, pos.y+other.y)
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
	for _, otherSide := range []side.Side{side.Top, side.Right, side.Left, side.Bottom} {
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
