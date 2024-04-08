package game

import (
	"errors"
	"fmt"
)


type ScoreReport struct {
	ReceivedPoints  map[int]uint32
	ReturnedMeeples map[int]uint8
}

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

// mutable type
type Board struct {
	// The information about the tile and its placement is stored sparsely
	// in a slice of size equal to the number of tiles in the set.
	tiles    []PlacedTile
	// tilesMap is used by the engine for faster lookups
	// but contains the same information as the `tiles` slice.
	tilesMap map[Position]PlacedTile
}

func NewBoard(maxTileCount int32) *Board {
	tiles := make([]PlacedTile, maxTileCount)
	tiles = append(tiles, StartingTile)
	return &Board{
		tiles: tiles,
		tilesMap: map[Position]PlacedTile{
			{0, 0}: StartingTile,
		},
	}
}

func (board *Board) TileCount() int {
	return len(board.tiles)
}

func (board *Board) GetLegalMovesFor(tile Tile) []LegalMove {
	// TODO for future tasks:
	// - implement generation of legal moves
	panic("not implemented")
}

// early return variant of above
func (board *Board) HasValidPlacement(tile Tile) bool {
	panic("not implemented")
}

func (board *Board) CanBePlaced(tile PlacedTile) bool {
	// TODO for future tasks:
	// - implement a way to check if a specified move is valid
	panic("not implemented")
}

// XXX: `PlacedTile` may just become `Tile` if the meeple field does not get split out:
// see https://github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pull/9#discussion_r1554723567
func (board *Board) PlaceTile(tile PlacedTile) (ScoreReport, error) {
	if len(board.tiles) == cap(board.tiles) {
		return ScoreReport{}, errors.New(
			"Board's tiles capacity exceeded, logic error?",
		)
	}
	// TODO for future tasks:
	// - determine if the tile can placed at a given position,
	//   or return InvalidMove otherwise
	board.tiles = append(board.tiles, tile)
	board.tilesMap[tile.pos] = tile
	scoreReport, err := board.checkCompleted(tile)
	if err != nil {
		return scoreReport, nil
	}
	panic("not implemented")
}

func (board *Board) checkCompleted(tile PlacedTile) (ScoreReport, error) {
	// TODO for future tasks:
	// - identify all completed features
	// - resolve control of the completed features
	// - award points
	scoreReport := ScoreReport{
		ReceivedPoints:  map[int]uint32{},
		ReturnedMeeples: map[int]uint8{},
	}
	panic("not implemented")
	return scoreReport, nil
}
