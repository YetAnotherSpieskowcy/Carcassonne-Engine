package game

import (
	"errors"
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/deck"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/stack"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/tiletemplates"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
)

func TestFullGame(t *testing.T) {
	tileSet := tilesets.StandardTileSet()
	tileSet.Tiles = []tiles.Tile{
		tiletemplates.SingleCityEdgeNoRoads(),
		tiletemplates.StraightRoads(),
	}
	deckStack := stack.NewOrdered(tileSet.Tiles)
	deck := deck.Deck{Stack: &deckStack, StartingTile: tileSet.StartingTile}
	game, err := NewFromDeck(deck, nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	// correct move with tile 0
	tile, err := game.GetCurrentTile()
	if err != nil {
		t.Fatal(err.Error())
	}
	ptile := elements.ToPlacedTile(tile)
	ptile.Position = elements.NewPosition(0, -1)
	err = game.PlayTurn(ptile)
	if err != nil {
		t.Fatal(err.Error())
	}

	// incorrect move - try placing tile 0 when 1 should be placed
	tile = tileSet.Tiles[0]
	ptile = elements.ToPlacedTile(tile)
	ptile.Position = elements.NewPosition(0, 1)
	err = game.PlayTurn(ptile)
	if err == nil {
		t.Fatal("expected error to occur")
	}
	if !errors.Is(err, elements.ErrWrongTile) {
		t.Fatal(err.Error())
	}

	// correct move with tile 1
	tile, err = game.GetCurrentTile()
	if err != nil {
		t.Fatal(err.Error())
	}
	ptile = elements.ToPlacedTile(tile)
	ptile.Position = elements.NewPosition(0, 1)
	err = game.PlayTurn(ptile)
	if err != nil {
		t.Fatal(err.Error())
	}

	ptile = elements.ToPlacedTile(tileSet.Tiles[1])
	ptile.Position = elements.NewPosition(0, 0)
	// check if out of bounds state is detected
	err = game.PlayTurn(ptile)
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
	game, err := NewFromTileSet(tilesets.StandardTileSet(), nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	// try finalizing before the game is finished
	_, err = game.Finalize()
	if err == nil {
		t.Fatal("expected error to occur")
	}
	if !errors.Is(err, elements.ErrGameIsNotFinished) {
		t.Fatal(err.Error())
	}
}
