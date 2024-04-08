package game

import (
	"reflect"
	"slices"
	"testing"
)


func TestBoardGetLegalMovesForReturnsEmptySliceWhenCityCannotBePlaced(t *testing.T) {
	// starting tile has a city on top, we want to close it with a single city tile
	// and then try finding legal moves of a tile filled with a city terrain
	board := NewBoard(5)
	board.PlaceTile(
		PlacedTile{
			LegalMove: LegalMove{
				Tile: SingleCityEdgeNoRoads().Rotate(2),
				pos: Position{0, 1},
			},
			meeple: Meeple{side: None},
		},
	)
	expected := []LegalMove{}

	actual := board.GetLegalMovesFor(FourCityEdgesConnectedShield())

	if !slices.Equal(expected, actual) {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}
}

func TestBoardHasValidPlacementReturnsTrueWhenValidPlacementExists(t *testing.T) {
	board := NewBoard(5)

	expected := true
	actual := board.HasValidPlacement(SingleCityEdgeNoRoads())

	if expected != actual {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}
}

func TestBoardCanBePlacedReturnsTrueWhenPlacedTileCanBePlaced(t *testing.T) {
	board := NewBoard(5)

	expected := true
	actual := board.CanBePlaced(getTestPlacedTile())

	if expected != actual {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}
}

func TestBoardPlaceTileErrorsWhenCapacityIsExceeded(t *testing.T) {
	board := NewBoard(1)

	_, err := board.PlaceTile(getTestPlacedTile())
	if err == nil {
		t.Fatal("expected capacity exceeded error to be returned")
	}
}

func TestBoardPlaceTileUpdatesBoardFields(t *testing.T) {
	board := NewBoard(5)
	expected := getTestPlacedTile()

	_, err := board.PlaceTile(expected)
	if err != nil {
		t.Fatal(err.Error())
	}

	actual := board.Tiles()[1]
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}

	actual, ok := board.GetTileAt(expected.pos)
	if !ok || !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %#v, got %#v instead (ok = %#v)", expected, actual, ok)
	}
}
