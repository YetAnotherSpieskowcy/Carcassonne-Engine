package elements

import (
	"fmt"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
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

func (pos Position) MarshalText() ([]byte, error) {
	return fmt.Appendf([]byte{}, "%v,%v", pos.x, pos.y), nil
}

func (pos *Position) UnmarshalText(text []byte) error {
	_, err := fmt.Sscanf(string(text), "%v,%v", &pos.x, &pos.y)
	return err
}

// https://wikicarpedia.com/car/Game_Figures
type MeepleType uint8

const (
	NormalMeeple MeepleType = iota

	MeepleTypeCount int = iota
)

type Meeple struct {
	Player Player
	Side   side.Side
	Type   MeepleType
}

type LegalMove struct {
	tiles.Tile
	Pos Position
}

type PlacedTile struct {
	LegalMove
	// PlacedTile always has a `Meeple`. Whether it is actually placed is determined by
	// `Meeple.side` which will be `None`, if it isn't.
	Meeple Meeple
}

func GetStandardStartingPlacedTile() PlacedTile {
	return PlacedTile{
		LegalMove: LegalMove{
			Tile: tilesets.GetStandardStartingTile(),
			Pos:  NewPosition(0, 0),
		},
	}
}
