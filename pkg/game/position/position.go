package position

import (
	"fmt"

	sideMod "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
)

type Position struct {
	// int8 would be fine for base game (72 tiles) but let's be a bit more generous
	x int16
	y int16
}

func NewPosition(x int16, y int16) Position {
	return Position{x, y}
}

func (pos Position) X() int16 {
	return pos.x
}

func (pos Position) Y() int16 {
	return pos.y
}

func (pos Position) Add(other Position) Position {
	return NewPosition(pos.x+other.x, pos.y+other.y)
}

/*
Returns relative position directed by the side.
Caution! It is supposed to be used with side directing only one cardinal direction (or two edges connected by corner)!
Otherwise it will return undesired value!
*/
func PositionFromSide(side sideMod.Side) Position {
	position := NewPosition(0, 0)

	if side&sideMod.Top != 0 {
		position = position.Add(NewPosition(0, 1))
	}

	if side&sideMod.Right != 0 {
		position = position.Add(NewPosition(1, 0))
	}

	if side&sideMod.Bottom != 0 {
		position = position.Add(NewPosition(0, -1))
	}

	if side&sideMod.Left != 0 {
		position = position.Add(NewPosition(-1, 0))
	}

	return position
}

func (pos Position) MarshalText() ([]byte, error) {
	return fmt.Appendf([]byte{}, "%v,%v", pos.x, pos.y), nil
}

func (pos *Position) UnmarshalText(text []byte) error {
	_, err := fmt.Sscanf(string(text), "%v,%v", &pos.x, &pos.y)
	return err
}
