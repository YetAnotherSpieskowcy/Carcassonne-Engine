package game

import (
	"errors"
	"reflect"
	"testing"

	. "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/test"
)


func getTestScoreReport() ScoreReport {
	return ScoreReport{
		ReceivedPoints: map[int]uint32{0: 5},
		ReturnedMeeples: map[int]uint8{},
	}
}

func TestPlayerPlaceTileErrorsWhenPlayerHasNoMeeples(t *testing.T) {
	player := NewPlayer(0)
	player.SetMeepleCount(0)

	board := NewBoard(5)
	tile := test.GetTestPlacedTile()
	_, err := player.PlaceTile(board, tile)
	if !errors.Is(err, NoMeepleAvailable) {
		t.Fatalf("expected NoMeepleAvailable error type, got %#v instead", err)
	}
}

func TestPlayerPlaceTileCallsBoardPlaceTile(t *testing.T) {
	player := NewPlayer(0)

	expectedScoreReport := getTestScoreReport()
	callCount := 0
	board := &test.TestBoard{PlaceTileFunc: func(tile PlacedTile) (ScoreReport, error) {
		callCount++
		return expectedScoreReport, nil
	}}

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
	player := NewPlayer(0)
	player.SetMeepleCount(2)
	expectedMeepleCount := uint8(1)

	board := &test.TestBoard{}
	tile := test.GetTestPlacedTile()

	_, err := player.PlaceTile(board, tile)
	if err != nil {
		t.Fatal(err.Error())
	}

	actualMeepleCount := player.MeepleCount()
	if actualMeepleCount != expectedMeepleCount {
		t.Fatalf("expected %#v, got %#v instead", expectedMeepleCount, actualMeepleCount)
	}
}

func TestPlayerPlaceTileKeepsMeepleCountWhenNoMeeplePlaced(t *testing.T) {
	player := NewPlayer(0)
	player.SetMeepleCount(2)
	expectedMeepleCount := uint8(2)

	board := &test.TestBoard{}
	tile := test.GetTestPlacedTileWithMeeple(Meeple{Side: None})

	_, err := player.PlaceTile(board, tile)
	if err != nil {
		t.Fatal(err.Error())
	}

	actualMeepleCount := player.MeepleCount()
	if actualMeepleCount != expectedMeepleCount {
		t.Fatalf("expected %#v, got %#v instead", expectedMeepleCount, actualMeepleCount)
	}
}

func TestPlayerPlaceTileKeepsMeepleCountWhenErrorReturned(t *testing.T) {
	player := NewPlayer(0)
	player.SetMeepleCount(2)
	expectedMeepleCount := uint8(2)

	board := &test.TestBoard{PlaceTileFunc: func(tile PlacedTile) (ScoreReport, error) {
		return ScoreReport{}, InvalidPosition
	}}
	tile := test.GetTestPlacedTile()

	_, err := player.PlaceTile(board, tile)
	if err == nil {
		t.Fatal("expected error to occur")
	}

	actualMeepleCount := player.MeepleCount()
	if actualMeepleCount != expectedMeepleCount {
		t.Fatalf("expected %#v, got %#v instead", expectedMeepleCount, actualMeepleCount)
	}
}
