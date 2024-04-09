package game

import (
	"errors"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
)

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
	tiles []elements.PlacedTile
	// tilesMap is used by the engine for faster lookups
	// but contains the same information as the `tiles` slice.
	tilesMap map[elements.Position]elements.PlacedTile
}

func NewBoard(maxTileCount int32) elements.Board {
	tiles := make([]elements.PlacedTile, 0, maxTileCount)
	tiles = append(tiles, elements.StartingTile)
	return &board{
		tiles: tiles,
		tilesMap: map[elements.Position]elements.PlacedTile{
			elements.NewPosition(0, 0): elements.StartingTile,
		},
	}
}

func (board *board) TileCount() int {
	return len(board.tilesMap)
}

func (board *board) Tiles() []elements.PlacedTile {
	return board.tiles
}

func (board *board) GetTileAt(pos elements.Position) (elements.PlacedTile, bool) {
	elem, ok := board.tilesMap[pos]
	return elem, ok
}

//revive:disable-next-line:unused-parameter Until the TODO is finished.
func (board *board) GetLegalMovesFor(tile elements.Tile) []elements.LegalMove {
	// TODO for future tasks:
	// - implement generation of legal moves
	return []elements.LegalMove{}
}

// early return variant of above
//
//revive:disable-next-line:unused-parameter Until the TODO is finished.
func (board *board) HasValidPlacement(tile elements.Tile) bool {
	// TODO for future tasks:
	// - implement generation of legal moves
	return true
}

//revive:disable-next-line:unused-parameter Until the TODO is finished.
func (board *board) CanBePlaced(tile elements.PlacedTile) bool {
	// TODO for future tasks:
	// - implement a way to check if a specified move is valid
	return true
}

func (board *board) PlaceTile(tile elements.PlacedTile) (elements.ScoreReport, error) {
	if len(board.tiles) == cap(board.tiles) {
		return elements.ScoreReport{}, errors.New(
			"Board's tiles capacity exceeded, logic error?",
		)
	}
	// TODO for future tasks:
	// - determine if the tile can placed at a given position,
	//   or return InvalidMove otherwise
	board.tiles = append(board.tiles, tile)
	board.tilesMap[tile.Pos] = tile
	scoreReport, err := board.checkCompleted(tile)
	return scoreReport, err
}

func (board *board) checkCompleted(
	//revive:disable-next-line:unused-parameter Until the TODO is finished.
	tile elements.PlacedTile,
) (elements.ScoreReport, error) {
	// TODO for future tasks:
	// - identify all completed features
	// - resolve control of the completed features
	// - award points
	scoreReport := elements.ScoreReport{
		ReceivedPoints:  map[int]uint32{},
		ReturnedMeeples: map[int][]uint8{},
	}
	return scoreReport, nil
}
