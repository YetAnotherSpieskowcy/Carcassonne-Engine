package game

import (
	"errors"
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/stack"
)

func TestFullGame(t *testing.T) {
	tiles := []elements.Tile{{ID: 1}, {ID: 2}}
	deck := stack.NewOrdered(tiles)
	game, err := NewGameWithDeck(&deck, nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	// correct move with tile 0
	tile, err := game.GetCurrentTile()
	if err != nil {
		t.Fatal(err.Error())
	}
	err = game.PlayTurn(
		elements.PlacedTile{LegalMove: elements.LegalMove{Tile: tile}},
	)
	if err != nil {
		t.Fatal(err.Error())
	}

	// incorrect move - try placing tile 0 when 1 should be placed
	tile = tiles[0]
	err = game.PlayTurn(
		elements.PlacedTile{LegalMove: elements.LegalMove{Tile: tile}},
	)
	if err == nil {
		t.Fatal("expected error to occur")
	}
	if !errors.Is(err, WrongTile) {
		t.Fatal(err.Error())
	}

	// correct move with tile 1
	tile, err = game.GetCurrentTile()
	if err != nil {
		t.Fatal(err.Error())
	}
	err = game.PlayTurn(
		elements.PlacedTile{LegalMove: elements.LegalMove{Tile: tile}},
	)
	if err != nil {
		t.Fatal(err.Error())
	}

	// check if out of bounds state is detected
	err = game.PlayTurn(
		elements.PlacedTile{LegalMove: elements.LegalMove{Tile: tiles[1]}},
	)
	if err == nil {
		t.Fatal("expected error to occur")
	}
	if !errors.Is(err, stack.ErrStackOutOfBounds) {
		t.Fatal(err.Error())
	}

	actualScores, err := game.Finalize()
	if err != nil {
		t.Fatal(err.Error())
	}

	expectedScores := []uint32{0, 0}
	for playerID, actual := range actualScores {
		expected := expectedScores[playerID]
		if actual != expected {
			t.Fatalf("expected %v, got %v for player %v instead", expected, actual, playerID)
		}
	}
}

func TestGameFinalizeErrorsBeforeGameIsFinished(t *testing.T) {
	game, err := NewGame(nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	// try finalizing before the game is finished
	_, err = game.Finalize()
	if err == nil {
		t.Fatal("expected error to occur")
	}
	if !errors.Is(err, ErrGameIsNotFinished) {
		t.Fatal(err.Error())
	}
}
