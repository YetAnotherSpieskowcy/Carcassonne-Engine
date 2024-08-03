package player_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/test"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/player"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
)

func TestPlayerDeepClone(t *testing.T) {
	meepleType := elements.NormalMeeple

	original := player.New(1)
	expected := original.MeepleCount(meepleType)
	clone := original.DeepClone()

	clone.SetMeepleCount(meepleType, expected+1)
	actual := original.MeepleCount(meepleType)

	if actual != expected {
		t.Fatalf("expected %v, got %v instead", expected, actual)
	}
}

func TestPlayerGetEligibleMovesFromReturnsAllMovesWhenPlayerHasMeeples(t *testing.T) {
	player := player.New(1)
	input := []elements.PlacedTile{test.GetTestPlacedTile()}
	expected := input[0]

	actual := player.GetEligibleMovesFrom(input)

	if !reflect.DeepEqual(expected, actual[0]) {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}
}

func TestPlayerGetEligibleMovesFromReturnsMovesWithoutMeepleWhenPlayerHasNoMeeples(t *testing.T) {
	player := player.New(1)
	player.SetMeepleCount(elements.NormalMeeple, 0)
	input := []elements.PlacedTile{test.GetTestPlacedTile()}
	expected := []elements.PlacedTile{input[0]}

	actual := player.GetEligibleMovesFrom(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}
}

func TestPlayerPlaceTileErrorsWhenPlayerHasNoMeeples(t *testing.T) {
	board := game.NewBoard(tilesets.StandardTileSet())
	tile := test.GetTestPlacedTile()
	player := player.New(1)
	player.SetMeepleCount(elements.NormalMeeple, 0)
	tile.Features[0].Meeple.Type = elements.NormalMeeple
	tile.Features[0].PlayerID = player.ID()

	_, err := player.PlaceTile(board, tile)
	if !errors.Is(err, elements.ErrNoMeepleAvailable) {
		t.Fatalf("expected NoMeepleAvailable error type, got %#v instead", err)
	}
}

func TestPlayerPlaceTileCallsBoardPlaceTile(t *testing.T) {
	expectedScoreReport := test.GetTestScoreReport()
	callCount := 0
	board := &test.BoardMock{
		PlaceTileFunc: func(_ elements.PlacedTile) (elements.ScoreReport, error) {
			callCount++
			return expectedScoreReport, nil
		},
	}

	tile := test.GetTestPlacedTile()
	player := player.New(1)

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
	board := &test.BoardMock{}
	tile := test.GetTestPlacedTile()
	player := player.New(1)
	player.SetMeepleCount(elements.NormalMeeple, 2)
	expectedMeepleCount := uint8(1)
	tile.Features[0].Meeple.Type = elements.NormalMeeple
	tile.Features[0].PlayerID = player.ID()

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
	board := &test.BoardMock{}
	tile := test.GetTestPlacedTile()
	player := player.New(1)
	player.SetMeepleCount(elements.NormalMeeple, 2)

	expectedMeepleCount := uint8(2)

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
	board := &test.BoardMock{
		PlaceTileFunc: func(_ elements.PlacedTile) (elements.ScoreReport, error) {
			return elements.ScoreReport{}, elements.ErrInvalidPosition
		},
	}
	tile := test.GetTestPlacedTile()
	player := player.New(1)
	player.SetMeepleCount(elements.NormalMeeple, 2)
	expectedMeepleCount := uint8(2)

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
	player := player.New(1)
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
	expectedID := elements.ID(6)
	player := player.New(expectedID)
	actualID := player.ID()
	if actualID != expectedID {
		t.Fatalf("expected %#v, got %#v instead", expectedID, actualID)
	}
}
