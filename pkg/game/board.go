package game

import (
	"errors"
	"fmt"
)


type ScoreReport struct {
	ReceivedPoints  map[int]uint32
	ReturnedMeeples map[int]uint8
}

// mutable type
type Board interface {
	TileCount() int
	Tiles() []PlacedTile
	GetTileAt(pos Position) (PlacedTile, bool)
	GetLegalMovesFor(tile Tile) []LegalMove
	HasValidPlacement(tile Tile) bool
	CanBePlaced(tile PlacedTile) bool
	PlaceTile(tile PlacedTile) (ScoreReport, error)
}

// mutable type
// Position coordinates example on the board:
// (-1, +1)  (+0, +1)  (+1, +1)
// (-1, +0)  (+0, +0)  (+1, +0)
// (-1, -1)  (+0, -1)  (+1, -1)
//
// Starting tile is placed at (+0, +0) position.
type board struct {
	// The information about the tile and its placement is stored sparsely
	// in a slice of size equal to the number of tiles in the set.
	tiles    []PlacedTile
	// tilesMap is used by the engine for faster lookups
	// but contains the same information as the `tiles` slice.
	tilesMap map[Position]PlacedTile
}

func NewBoard(maxTileCount int32) Board {
	tiles := make([]PlacedTile, maxTileCount)
	tiles = append(tiles, StartingTile)
	return &board{
		tiles: tiles,
		tilesMap: map[Position]PlacedTile{
			{0, 0}: StartingTile,
		},
	}
}

func (board *board) TileCount() int {
	return len(board.tiles)
}

func (board *board) Tiles() []PlacedTile {
	return board.tiles
}

func (board *board) GetTileAt(pos Position) (PlacedTile, bool) {
	elem, ok := board.tilesMap[pos]
	return elem, ok
}

func (board *board) GetLegalMovesFor(tile Tile) []LegalMove {
	// TODO for future tasks:
	// - implement generation of legal moves
	panic("not implemented")
}

// early return variant of above
func (board *board) HasValidPlacement(tile Tile) bool {
	panic("not implemented")
}

func (board *board) CanBePlaced(tile PlacedTile) bool {
	// TODO for future tasks:
	// - implement a way to check if a specified move is valid
	panic("not implemented")
}

func (board *board) PlaceTile(tile PlacedTile) (ScoreReport, error) {
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

func (board *board) checkCompleted(tile PlacedTile) (ScoreReport, error) {
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
