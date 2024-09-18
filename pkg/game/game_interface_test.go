package game

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/deck"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/stack"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
)

func TestGetBoard(t *testing.T) {
	tileSet := tilesets.StandardTileSet()
	deckStack := stack.NewOrdered(tileSet.Tiles)
	deck := deck.Deck{Stack: &deckStack, StartingTile: tileSet.StartingTile}

	game, err := NewFromDeck(deck, nil, 2)
	if err != nil {
		t.Fatal(err.Error())
	}

	if game.GetBoard() == nil {
		t.Fatalf("Couldn't get board")
	}
}
