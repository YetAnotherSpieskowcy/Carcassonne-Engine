package game

import (
	"testing"

	. "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/stack"
)

func TestFullGame(t *testing.T) {
	tiles := []Tile{Tile{}, Tile{}}
    deck := stack.NewOrdered(tiles)
    game, err := NewGameWithDeck(&deck, nil)
    if err != nil {
    	t.Fatal(err.Error())
    }

    for range 2 {
		tile, err := game.GetCurrentTile()
		if err != nil {
			t.Fatal(err.Error())
		}
		err = game.PlayTurn(PlacedTile{LegalMove: LegalMove{Tile: tile}})
		if err != nil {
			t.Fatal(err.Error())
		}
	}

    actualScores, err := game.Finalize()
    if err != nil {
    	t.Fatal(err)
    }

    expectedScores := []uint32{0, 0}
    for playerId, actual := range actualScores {
    	expected := expectedScores[playerId]
    	if actual != expected {
    		t.Fatalf("expected %v, got %v for player %v instead", expected, actual, playerId)
    	}
    }
}
