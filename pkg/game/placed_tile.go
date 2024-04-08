package game

import "fmt"


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

func (pos Position) MarshalText() ([]byte, error) {
	return fmt.Appendf([]byte{}, "%v,%v", pos.x, pos.y), nil
}

type Meeple struct {
	player Player
	side   Side
}

type LegalMove struct {
	Tile
	pos  Position
}

type PlacedTile struct {
	LegalMove
	// PlacedTile always has a `Meeple`. Whether it is actually placed is determined by
	// `Meeple.side` which will be `None`, if it isn't.
	meeple    Meeple
}
