package game

import (
	"errors"
	"slices"

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
	tileSet []elements.Tile
	// The information about the tile and its placement is stored sparsely
	// in a slice of size equal to the number of tiles in the set.
	// `tiles[0]` is always the starting tile.
	// Indexes in `tiles[1:]` correspond to positions of equivalent tiles in `tileSet`.
	tiles []elements.PlacedTile
	// tilesMap is used by the engine for faster lookups
	// but contains the same information as the `tiles` slice.
	tilesMap map[elements.Position]elements.PlacedTile
}

func NewBoard(tileSet []elements.Tile) elements.Board {
	tiles := make([]elements.PlacedTile, len(tileSet)+1)
	tiles[0] = elements.StartingTile
	return &board{
		tileSet: tileSet,
		tiles:   tiles,
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
	if board.TileCount() == cap(board.tiles) {
		return elements.ScoreReport{}, errors.New(
			"Board's tiles capacity exceeded, logic error?",
		)
	}
	// TODO for future tasks:
	// - determine if the tile can placed at a given position,
	//   or return InvalidMove otherwise
	tileSet := board.tileSet
	actualIndex := 1
	for {
		index := slices.IndexFunc(tileSet, func(candidate elements.Tile) bool {
			return tile.Tile == candidate
		})
		if index == -1 {
			return elements.ScoreReport{}, errors.New(
				"Placed tile not found in the tile set, logic error?",
			)
		}
		actualIndex += index
		if board.tiles[actualIndex].Tile != tile.Tile {
			break
		}
		// position already taken, gotta find another next matching tile
		actualIndex++
		tileSet = tileSet[index+1:]
	}
	board.tiles[actualIndex] = tile
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
