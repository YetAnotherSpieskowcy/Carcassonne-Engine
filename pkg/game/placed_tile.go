package game

import "fmt"


type Position struct {
	// int8 would be fine for base game (72 tiles) but let's be a bit more generous
	x int16
	y int16
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
	side   any
}

type LegalMove struct {
	Tile
	pos  Position
}

type PlacedTile struct {
	LegalMove
	Meeple    Meeple
}
