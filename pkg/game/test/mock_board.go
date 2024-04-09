package test

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
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

func (board *BoardMock) GetTileAt(_ elements.Position) (elements.PlacedTile, bool) {
	return elements.PlacedTile{}, true
}

func (board *BoardMock) GetLegalMovesFor(_ elements.Tile) []elements.LegalMove {
	return []elements.LegalMove{}
}

func (board *BoardMock) HasValidPlacement(_ elements.Tile) bool {
	return true
}

func (board *BoardMock) CanBePlaced(_ elements.PlacedTile) bool {
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
