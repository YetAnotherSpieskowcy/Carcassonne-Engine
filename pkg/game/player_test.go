package game

import (
	"errors"
	"reflect"
	"testing"
)

type testBoard struct {
	tileCount func() int
	placeTile func(tile PlacedTile) (ScoreReport, error)
}

func (board *testBoard) TileCount() int {
	if board.tileCount == nil {
		return 0
	}
	return board.tileCount()
}

func (board *testBoard) GetLegalMovesFor(tile Tile) []LegalMove {
	return []LegalMove{}
}

func (board *testBoard) HasValidPlacement(tile Tile) bool {
	return true
}

func (board *testBoard) CanBePlaced(tile PlacedTile) bool {
	return true
}

func (board *testBoard) PlaceTile(tile PlacedTile) (ScoreReport, error) {
	if board.placeTile == nil {
		return ScoreReport{}, nil
	}
	return board.placeTile(tile)
}

func getTestTile() Tile {
	return SingleCityEdgeNoRoads().Rotate(2)
}

func getTestPlacedTile() PlacedTile {
	return PlacedTile{
		LegalMove: LegalMove{Tile: getTestTile(), pos: Position{0, 1}},
		meeple: Meeple{side: Bottom},
	}
}

func getTestPlacedTileWithMeeple(meeple Meeple) PlacedTile {
	return PlacedTile{
		LegalMove: LegalMove{Tile: getTestTile(), pos: Position{0, 1}},
		meeple: meeple,
	}
}

func getTestScoreReport() ScoreReport {
	return ScoreReport{
		ReceivedPoints: map[int]uint32{0: 5},
		ReturnedMeeples: map[int]uint8{},
	}
}

func TestPlayerPlaceTileErrorsWhenPlayerHasNoMeeples(t *testing.T) {
	player := NewPlayer(0)
	player.meepleCount = 0

	board := NewBoard(5)
	tile := getTestPlacedTile()
	_, err := player.PlaceTile(board, tile)
	if !errors.Is(err, NoMeepleAvailable) {
		t.Fatalf("expected NoMeepleAvailable error type, got %#v instead", err)
	}
}

func TestPlayerPlaceTileCallsBoardPlaceTile(t *testing.T) {
	player := NewPlayer(0)

	expectedScoreReport := getTestScoreReport()
	callCount := 0
	board := &testBoard{placeTile: func(tile PlacedTile) (ScoreReport, error) {
		callCount++
		return expectedScoreReport, nil
	}}

	tile := getTestPlacedTile()

	actualScoreReport, err := player.PlaceTile(board, tile)
	if err != nil {
		t.Fatal(err.Error())
	}

	if !reflect.DeepEqual(actualScoreReport, expectedScoreReport) {
		t.Fatalf("expected %#v, got %#v instead", expectedScoreReport, actualScoreReport)
	}

	if callCount != 1 {
		t.Fatal("expected board.PlaceTile() to be called once")
	}
}

func TestPlayerPlaceTileLowersMeepleCountWhenMeeplePlaced(t *testing.T) {
	player := NewPlayer(0)
	player.meepleCount = 2
	expectedMeepleCount := uint8(1)

	board := &testBoard{}
	tile := getTestPlacedTile()

	_, err := player.PlaceTile(board, tile)
	if err != nil {
		t.Fatal(err.Error())
	}

	actualMeepleCount := player.MeepleCount()
	if actualMeepleCount != expectedMeepleCount {
		t.Fatalf("expected %#v, got %#v instead", expectedMeepleCount, actualMeepleCount)
	}
}

func TestPlayerPlaceTileKeepsMeepleCountWhenMeeplePlaced(t *testing.T) {
	player := NewPlayer(0)
	player.meepleCount = 2
	expectedMeepleCount := uint8(2)

	board := &testBoard{}
	tile := getTestPlacedTileWithMeeple(Meeple{side: None})

	_, err := player.PlaceTile(board, tile)
	if err != nil {
		t.Fatal(err.Error())
	}

	actualMeepleCount := player.MeepleCount()
	if actualMeepleCount != expectedMeepleCount {
		t.Fatalf("expected %#v, got %#v instead", expectedMeepleCount, actualMeepleCount)
	}
}
