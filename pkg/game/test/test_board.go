package test

import (
	. "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
)

type TestBoard struct {
	TileCountFunc func() int
	PlaceTileFunc func(tile PlacedTile) (ScoreReport, error)
}

func (board *TestBoard) TileCount() int {
	if board.TileCountFunc == nil {
		return 0
	}
	return board.TileCountFunc()
}

func (board *TestBoard) Tiles() []PlacedTile {
	return []PlacedTile{}
}

func (board *TestBoard) GetTileAt(_ Position) (PlacedTile, bool) {
	return PlacedTile{}, true
}

func (board *TestBoard) GetLegalMovesFor(_ Tile) []LegalMove {
	return []LegalMove{}
}

func (board *TestBoard) HasValidPlacement(_ Tile) bool {
	return true
}

func (board *TestBoard) CanBePlaced(_ PlacedTile) bool {
	return true
}

func (board *TestBoard) PlaceTile(tile PlacedTile) (ScoreReport, error) {
	if board.PlaceTileFunc == nil {
		return GetTestScoreReport(), nil
	}
	return board.PlaceTileFunc(tile)
}
