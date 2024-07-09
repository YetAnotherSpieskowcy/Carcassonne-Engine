package game

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/tiletemplates"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
)

func TestStartingTilePlacement(t *testing.T) {
	board := NewBoard(tilesets.StandardTileSet())
	actual := len(board.GetTilePlacementsFor(tilesets.StandardTileSet().StartingTile))
	expected := 6
	if actual != expected {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}
}

func TestStraightRoadsPlacement(t *testing.T) {
	board := NewBoard(tilesets.StandardTileSet())
	actual := len(board.GetTilePlacementsFor(tiletemplates.StraightRoads()))
	expected := 3
	if actual != expected {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}
}
func TestMultipleStraightRoadsPlacement(t *testing.T) {
	board := NewBoard(tilesets.StandardTileSet())
	actual := len(board.GetTilePlacementsFor(tiletemplates.StraightRoads()))
	expected := 3
	if actual != expected {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}

	ptile := elements.ToPlacedTile(tiletemplates.StraightRoads())
	ptile.Position = elements.NewPosition(1, 0)
	_, err := board.PlaceTile(ptile)
	if err != nil {
		t.Fail()
	}

	actual = len(board.GetTilePlacementsFor(tiletemplates.StraightRoads()))
	expected = 5
	if actual != expected {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}

	ptile = elements.ToPlacedTile(tiletemplates.StraightRoads())
	ptile.Position = elements.NewPosition(2, 0)
	_, err = board.PlaceTile(ptile)
	if err != nil {
		t.Fail()
	}

	actual = len(board.GetTilePlacementsFor(tiletemplates.StraightRoads()))
	expected = 7
	if actual != expected {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}
}
