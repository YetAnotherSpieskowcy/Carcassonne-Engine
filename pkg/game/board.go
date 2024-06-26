package game

import (
	"errors"
	"slices"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
)

// mutable type
// Position coordinates example on the board:
// (-1, +1)  (+0, +1)  (+1, +1)
// (-1, +0)  (+0, +0)  (+1, +0)
// (-1, -1)  (+0, -1)  (+1, -1)
//
// Starting tile is placed at (+0, +0) position.
type board struct {
	tileSet tilesets.TileSet
	// The information about the tile and its placement is stored sparsely
	// in a slice of size equal to the number of tiles in the set.
	// `tiles[0]` is always the starting tile.
	// Indexes in `tiles[1:]` correspond to positions of equivalent tiles in `tileSet`.
	tiles []elements.PlacedTile
	// tilesMap is used by the engine for faster lookups
	// but contains the same information as the `tiles` slice.
	tilesMap map[elements.Position]elements.PlacedTile

	placeablePositions []elements.Position
}

func NewBoard(tileSet tilesets.TileSet) elements.Board {
	tiles := make([]elements.PlacedTile, len(tileSet.Tiles)+1)
	startingTile := elements.NewStartingTile(tileSet)
	tiles[0] = startingTile
	return &board{
		tileSet: tileSet,
		tiles:   tiles,
		tilesMap: map[elements.Position]elements.PlacedTile{
			elements.NewPosition(0, 0): startingTile,
		},
		placeablePositions: []elements.Position{
			elements.NewPosition(0, 1),
			elements.NewPosition(1, 0),
			elements.NewPosition(0, -1),
			elements.NewPosition(-1, 0),
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

func (board *board) GetTilePlacementsFor(tile tiles.Tile) []elements.PlacedTile {
	valid := []elements.PlacedTile{}
	rotations := tile.GetTileRotations()
	for _, currentTile := range rotations {
		for _, placeable := range board.placeablePositions {
			tilePlacement := elements.ToPlacedTile(currentTile)
			tilePlacement.Position = placeable
			if board.isPositionValid(tilePlacement) {
				valid = append(valid, tilePlacement)
			}
		}
	}
	return valid
}

// Verifies if a certain side of a adjacent tile on the board matches an expected feature type.
// The method takes a board, a position, the expected side, and the expected feature type.
// Returns a boolean indicating whether the tile has an expected feature on specified side.
func (board *board) testSide(position elements.Position, expectedSide side.Side, expectedFeatureType feature.Type) bool {
	var tile elements.PlacedTile
	var ok bool
	switch expectedSide {
	case side.Bottom:
		tile, ok = board.tilesMap[elements.NewPosition(position.X(), position.Y()+1)]
	case side.Top:
		tile, ok = board.tilesMap[elements.NewPosition(position.X(), position.Y()-1)]
	case side.Left:
		tile, ok = board.tilesMap[elements.NewPosition(position.X()+1, position.Y())]
	case side.Right:
		tile, ok = board.tilesMap[elements.NewPosition(position.X()-1, position.Y())]
	}
	if !ok {
		return true
	}
	for _, tileFeature := range tile.Features {
		if tileFeature.FeatureType == expectedFeatureType {
			if tileFeature.Sides&expectedSide == expectedSide {
				return true
			}
		}
	}
	return false
}

func (board *board) TileHasValidPlacement(tile tiles.Tile) bool {
	rotations := tile.GetTileRotations()
	for _, currentTile := range rotations {
		for _, placeable := range board.placeablePositions {
			tilePlacement := elements.ToPlacedTile(currentTile)
			tilePlacement.Position = placeable
			if board.isPositionValid(tilePlacement) {
				return true
			}
		}
	}
	return false
}

//revive:disable-next-line:unused-parameter Until the TODO is finished.
func (board *board) GetLegalMovesFor(tile elements.PlacedTile) []elements.PlacedTile {
	// TODO for future tasks:
	// - implement generation of legal moves
	// - to be implemented after #18, #19 and #20 to avoid code duplication
	return []elements.PlacedTile{}
}

func (board *board) isPositionValid(tile elements.PlacedTile) bool {
	for _, f := range tile.Features {
		s := side.Top
		for range 4 {
			if f.Sides&s == s &&
				!board.testSide(tile.Position, s.Rotate(2), f.FeatureType) {
				return false
			}
			s = s.Rotate(1)
		}
	}
	return true
}

//revive:disable-next-line:unused-parameter Until the TODO is finished.
func (board *board) CanBePlaced(tile elements.PlacedTile) bool {
	// TODO for future tasks:
	// - implement generation of legal moves
	// - to be implemented after #18, #19 and #20 to avoid code duplication
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
	//   or return ErrInvalidMove otherwise
	setTiles := board.tileSet.Tiles
	actualIndex := 1
	for {
		index := slices.IndexFunc(setTiles, func(candidate tiles.Tile) bool {
			return elements.ToTile(tile).Equals(candidate)
		})
		if index == -1 {
			return elements.ScoreReport{}, errors.New(
				"Placed tile not found in the tile set, logic error?",
			)
		}
		actualIndex += index
		if !elements.ToTile(board.tiles[actualIndex]).Equals(elements.ToTile(tile)) {
			break
		}
		// position already taken, gotta find another next matching tile
		actualIndex++
		setTiles = setTiles[index+1:]
	}

	board.updateValidPlacements(tile)
	board.tiles[actualIndex] = tile
	board.tilesMap[tile.Position] = tile
	scoreReport, err := board.checkCompleted(tile)
	return scoreReport, err
}

func (board *board) updateValidPlacements(tile elements.PlacedTile) {
	tileIndex := slices.Index(board.placeablePositions, tile.Position)
	if tileIndex == -1 {
		panic("Invalid move was played")
	}
	board.placeablePositions = slices.Delete(board.placeablePositions, tileIndex, tileIndex+1)
	validNewPositions := []elements.Position{
		elements.NewPosition(tile.Position.X()+1, tile.Position.Y()),
		elements.NewPosition(tile.Position.X()-1, tile.Position.Y()),
		elements.NewPosition(tile.Position.X(), tile.Position.Y()+1),
		elements.NewPosition(tile.Position.X(), tile.Position.Y()-1),
	}
	for _, position := range validNewPositions {
		_, ok := board.tilesMap[position]
		if !ok && !slices.Contains(board.placeablePositions, position) {
			board.placeablePositions = append(board.placeablePositions, position)
		}
	}
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
		ReceivedPoints:  map[uint8]uint32{},
		ReturnedMeeples: map[uint8][]uint8{},
	}
	return scoreReport, nil
}
