package deck

import (
	"reflect"
	"slices"
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/stack"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
)

func TestDeepClone(t *testing.T) {
	tileSet := tilesets.StandardTileSet()
	deckStack := stack.New(tileSet.Tiles)

	original := Deck{Stack: &deckStack, StartingTile: tileSet.StartingTile}
	expected := original.GetRemainingTileCount()
	clone := original.DeepClone()

	actual := clone.GetRemainingTileCount()
	if actual != expected {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}

	if _, err := clone.Next(); err != nil {
		t.Fatal(err.Error())
	}

	actual = clone.GetRemainingTileCount()
	if actual != expected-1 {
		t.Fatalf("expected %#v, got %#v instead", expected-1, actual)
	}
	actual = original.GetRemainingTileCount()
	if actual != expected {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}

	if !slices.Equal(original.StartingTile.Features, clone.StartingTile.Features) {
		t.Fatalf("expected %#v, got %#v instead", original.StartingTile, clone.StartingTile)
	}
}

func TestTileSet(t *testing.T) {
	tileSet := tilesets.StandardTileSet()
	deckStack := stack.New(tileSet.Tiles)
	deck := Deck{Stack: &deckStack, StartingTile: tileSet.StartingTile}

	expected := tileSet
	actual := deck.TileSet()

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}
}
