package deck

import (
	"reflect"
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine-API/pkg/tilesets"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/stack"
)

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
