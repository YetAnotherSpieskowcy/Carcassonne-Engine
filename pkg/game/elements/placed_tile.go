package elements

import (
	"fmt"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	sideMod "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
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

func (pos Position) Add(other Position) Position {
	return NewPosition(pos.x+other.x, pos.y+other.y)
}

func PositionFromSide(side sideMod.Side) Position { //nolint:gocyclo // similar problem to sides
	switch side {
	case sideMod.Top:
		return NewPosition(0, 1)
	case sideMod.Right:
		return NewPosition(1, 0)
	case sideMod.Left:
		return NewPosition(-1, 0)
	case sideMod.Bottom:
		return NewPosition(0, -1)

	case sideMod.TopLeftEdge:
		return NewPosition(0, 1)
	case sideMod.TopRightEdge:
		return NewPosition(0, 1)
	case sideMod.RightTopEdge:
		return NewPosition(1, 0)
	case sideMod.RightBottomEdge:
		return NewPosition(1, 0)

	case sideMod.LeftTopEdge:
		return NewPosition(-1, 0)
	case sideMod.LeftBottomEdge:
		return NewPosition(-1, 0)
	case sideMod.BottomLeftEdge:
		return NewPosition(0, -1)
	case sideMod.BottomRightEdge:
		return NewPosition(0, -1)

	default:
		return NewPosition(0, 0)
	}
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

// note: this is implemented to avoid inheriting Tile's implementation
// but is just a simple passthrough and its output may not be a legal placement.
func (placement TilePlacement) Rotate(rotations uint) TilePlacement {
	panic("Rotate() not supported on TilePlacement")
}

// represents a legal position of a meeple on the tile
type MeeplePlacement struct {
	Side sideMod.Side
	Type MeepleType
}

// note: this is implemented to aid with LegalMove.Rotate() implementation
// but is just a simple passthrough and its output may not be a legal placement.
func (placement MeeplePlacement) Rotate(rotations uint) MeeplePlacement {
	return MeeplePlacement{
		Side: placement.Side.Rotate(rotations),
		Type: placement.Type,
	}
}

// represents a legal move (tile placement and meeple placement) on the board
type LegalMove struct {
	TilePlacement
	// LegalMove always has a `Meeple`. Whether it is actually placed
	// is determined by `MeeplePlacement.Side` which will be `None`, if it isn't.
	Meeple MeeplePlacement
}

// note: this is implemented to avoid inheriting the implementation
// but is just a simple passthrough and its output may not be a legal move.
func (move LegalMove) Rotate(rotations uint) LegalMove {
	return LegalMove{
		TilePlacement: move.TilePlacement.Rotate(rotations),
		Meeple:        move.Meeple.Rotate(rotations),
	}
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
