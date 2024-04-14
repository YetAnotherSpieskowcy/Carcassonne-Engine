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

	_, err := board.PlaceTile(elements.PlacedTile{
		LegalMove: elements.LegalMove{
			TilePlacement: elements.TilePlacement{Tile: tiletemplates.StraightRoads(), Pos: elements.NewPosition(1, 0)},
		},
	})
	if err != nil {
		t.Fail()
	}

	actual = len(board.GetTilePlacementsFor(tiletemplates.StraightRoads()))
	expected = 5
	if actual != expected {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}

	_, err = board.PlaceTile(elements.PlacedTile{LegalMove: elements.LegalMove{
		TilePlacement: elements.TilePlacement{Tile: tiletemplates.StraightRoads(), Pos: elements.NewPosition(2, 0)},
	},
	})
	if err != nil {
		t.Fail()
	}

	actual = len(board.GetTilePlacementsFor(tiletemplates.StraightRoads()))
	expected = 7
	if actual != expected {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}
}
