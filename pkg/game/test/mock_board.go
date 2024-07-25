package test

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
)

type BoardMock struct {
	TileCountFunc func() int
	PlaceTileFunc func(tile elements.PlacedTile) (elements.ScoreReport, error)
}

func (board *BoardMock) TileCount() int {
	if board.TileCountFunc == nil {
		return 0
	}
	return board.TileCountFunc()
}

func (board *BoardMock) Tiles() []elements.PlacedTile {
	return []elements.PlacedTile{}
}

func (board *BoardMock) GetTileAt(pos position.Position) (elements.PlacedTile, bool) {
	_ = pos
	return elements.PlacedTile{}, true
}

func (board *BoardMock) GetTilePlacementsFor(tile tiles.Tile) []elements.PlacedTile {
	_ = tile
	return []elements.PlacedTile{}
}

func (board *BoardMock) TileHasValidPlacement(tile tiles.Tile) bool {
	_ = tile
	return true
}

func (board *BoardMock) GetLegalMovesFor(tile elements.PlacedTile) []elements.PlacedTile {
	_ = tile
	return []elements.PlacedTile{}
}

func (board *BoardMock) CanBePlaced(tile elements.PlacedTile) bool {
	_ = tile
	return true
}

func (board *BoardMock) PlaceTile(
	tile elements.PlacedTile,
) (elements.ScoreReport, error) {
	if board.PlaceTileFunc == nil {
		return GetTestScoreReport(), nil
	}
	return board.PlaceTileFunc(tile)
}

func (board *BoardMock) RemoveMeeple(meeple elements.Meeple, position position.Position, side side.Side, featureType feature.Type) {
}
