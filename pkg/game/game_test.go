package game

import (
	"errors"
	"testing"

	. "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/stack"
)

func TestFullGame(t *testing.T) {
	tiles := []Tile{Tile{Id: 0}, Tile{Id: 1}}
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
	err = game.PlayTurn(PlacedTile{LegalMove: LegalMove{Tile: tile}})
	if err != nil {
		t.Fatal(err.Error())
	}

    // incorrect move - try placing tile 0 when 1 should be placed
	tile = tiles[0]
	err = game.PlayTurn(PlacedTile{LegalMove: LegalMove{Tile: tile}})
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
	err = game.PlayTurn(PlacedTile{LegalMove: LegalMove{Tile: tile}})
	if err != nil {
		t.Fatal(err.Error())
	}

	// check if out of bounds state is detected
	err = game.PlayTurn(PlacedTile{LegalMove: LegalMove{Tile: tiles[1]}})
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
    for playerId, actual := range actualScores {
    	expected := expectedScores[playerId]
    	if actual != expected {
    		t.Fatalf("expected %v, got %v for player %v instead", expected, actual, playerId)
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
    if !errors.Is(err, GameIsNotFinished) {
    	t.Fatal(err.Error())
    }
}
