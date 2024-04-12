package player_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/test"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/player"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
)

func getTestScoreReport() elements.ScoreReport {
	return elements.ScoreReport{
		ReceivedPoints:  map[int]uint32{0: 5},
		ReturnedMeeples: map[int][]uint8{},
	}
}

func TestPlayerPlaceTileErrorsWhenPlayerHasNoMeeples(t *testing.T) {
	player := player.New(0)
	player.SetMeepleCount(elements.NormalMeeple, 0)

	board := game.NewBoard(tilesets.GetStandardTiles())
	tile := test.GetTestPlacedTile()
	_, err := player.PlaceTile(board, tile)
	if !errors.Is(err, elements.ErrNoMeepleAvailable) {
		t.Fatalf("expected NoMeepleAvailable error type, got %#v instead", err)
	}
}

func TestPlayerPlaceTileCallsBoardPlaceTile(t *testing.T) {
	player := player.New(0)

	expectedScoreReport := getTestScoreReport()
	callCount := 0
	board := &test.BoardMock{
		PlaceTileFunc: func(_ elements.PlacedTile) (elements.ScoreReport, error) {
			callCount++
			return expectedScoreReport, nil
		},
	}

	tile := test.GetTestPlacedTile()

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
	player := player.New(0)
	player.SetMeepleCount(elements.NormalMeeple, 2)
	expectedMeepleCount := uint8(1)

	board := &test.BoardMock{}
	tile := test.GetTestPlacedTile()

	_, err := player.PlaceTile(board, tile)
	if err != nil {
		t.Fatal(err.Error())
	}

	actualMeepleCount := player.MeepleCount(elements.NormalMeeple)
	if actualMeepleCount != expectedMeepleCount {
		t.Fatalf("expected %#v, got %#v instead", expectedMeepleCount, actualMeepleCount)
	}
}

func TestPlayerPlaceTileKeepsMeepleCountWhenNoMeeplePlaced(t *testing.T) {
	player := player.New(0)
	player.SetMeepleCount(elements.NormalMeeple, 2)
	expectedMeepleCount := uint8(2)

	board := &test.BoardMock{}
	tile := test.GetTestPlacedTileWithMeeple(elements.Meeple{Side: side.None})

	_, err := player.PlaceTile(board, tile)
	if err != nil {
		t.Fatal(err.Error())
	}

	actualMeepleCount := player.MeepleCount(elements.NormalMeeple)
	if actualMeepleCount != expectedMeepleCount {
		t.Fatalf("expected %#v, got %#v instead", expectedMeepleCount, actualMeepleCount)
	}
}

func TestPlayerPlaceTileKeepsMeepleCountWhenErrorReturned(t *testing.T) {
	player := player.New(0)
	player.SetMeepleCount(elements.NormalMeeple, 2)
	expectedMeepleCount := uint8(2)

	board := &test.BoardMock{
		PlaceTileFunc: func(_ elements.PlacedTile) (elements.ScoreReport, error) {
			return elements.ScoreReport{}, elements.ErrInvalidPosition
		},
	}
	tile := test.GetTestPlacedTile()

	_, err := player.PlaceTile(board, tile)
	if err == nil {
		t.Fatal("expected error to occur")
	}

	actualMeepleCount := player.MeepleCount(elements.NormalMeeple)
	if actualMeepleCount != expectedMeepleCount {
		t.Fatalf("expected %#v, got %#v instead", expectedMeepleCount, actualMeepleCount)
	}
}

func TestPlayerScoreUpdatesAfterSet(t *testing.T) {
	player := player.New(0)
	actualScore := player.Score()
	if actualScore != 0 {
		t.Fatalf("expected %#v, got %#v instead", 0, actualScore)
	}

	player.SetScore(2)

	expectedScore := uint32(2)
	actualScore = player.Score()
	if actualScore != expectedScore {
		t.Fatalf("expected %#v, got %#v instead", expectedScore, actualScore)
	}
}

func TestPlayerNewPlayerSetsId(t *testing.T) {
	expectedID := uint8(6)
	player := player.New(expectedID)
	actualID := player.ID()
	if actualID != expectedID {
		t.Fatalf("expected %#v, got %#v instead", expectedID, actualID)
	}
}
