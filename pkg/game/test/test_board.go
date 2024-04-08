package test

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
)

type TestBoard struct {
	TileCountFunc func() int
	PlaceTileFunc func(tile elements.PlacedTile) (elements.ScoreReport, error)
}

func (board *TestBoard) TileCount() int {
	if board.TileCountFunc == nil {
		return 0
	}
	return board.TileCountFunc()
}

func (board *TestBoard) Tiles() []elements.PlacedTile {
	return []elements.PlacedTile{}
}

func (board *TestBoard) GetTileAt(_ elements.Position) (elements.PlacedTile, bool) {
	return elements.PlacedTile{}, true
}

func (board *TestBoard) GetLegalMovesFor(_ elements.Tile) []elements.LegalMove {
	return []elements.LegalMove{}
}

func (board *TestBoard) HasValidPlacement(_ elements.Tile) bool {
	return true
}

func (board *TestBoard) CanBePlaced(_ elements.PlacedTile) bool {
	return true
}

func (board *TestBoard) PlaceTile(
	tile elements.PlacedTile,
) (elements.ScoreReport, error) {
	if board.PlaceTileFunc == nil {
		return GetTestScoreReport(), nil
	}
	return board.PlaceTileFunc(tile)
}
