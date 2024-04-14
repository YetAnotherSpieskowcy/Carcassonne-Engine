package elements

import (
	"fmt"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
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

// represents a legal position (and rotation) of a tile on the board
type TilePlacement struct {
	tiles.Tile
	Pos Position
}

func (placement TilePlacement) Rotate(_ uint) TilePlacement {
	panic("Rotate() not supported on TilePlacement")
}

// represents a legal position of a meeple on the tile
type MeeplePlacement struct {
	Feature feature.Feature
	Type    MeepleType
}

// represents a legal move (tile placement and meeple placement) on the board
type LegalMove struct {
	TilePlacement
	// LegalMove always has a `Meeple`. Whether it is actually placed
	// is determined by `MeeplePlacement.Feature.FeatureType` which will be `None`,
	// if it isn't.
	Meeple MeeplePlacement
}

// represents a tile placed on the board, including the player who placed it
type PlacedTile struct {
	LegalMove
	// Although the player field is always set, it technically is only crucial to
	// the game state *if* a meeple was placed.
	// For starting tile, Player with ID 0 is used.
	Player Player
}

func NewStartingTile(tileSet tilesets.TileSet) PlacedTile {
	return PlacedTile{
		LegalMove: LegalMove{
			TilePlacement: TilePlacement{
				Tile: tileSet.StartingTile,
				Pos:  NewPosition(0, 0),
			},
		},
	}
}
