package test

import (
	"reflect"
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
)

func TestBoardMockTileCountWithoutFunc(t *testing.T) {
	board := BoardMock{}
	actual := board.TileCount()
	expected := 0
	if actual != expected {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}
}

func TestBoardMockTileCountWithFunc(t *testing.T) {
	wasCalled := false
	expected := 2
	board := BoardMock{TileCountFunc: func() int {
		wasCalled = true
		return expected
	}}
	actual := board.TileCount()
	if actual != expected {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}
	if !wasCalled {
		t.Fatal("expected TileCount() to be called")
	}
}

func TestBoardMockTiles(t *testing.T) {
	board := BoardMock{}
	actual := board.Tiles()
	if len(actual) != 0 {
		t.Fatalf("expected Tiles() output to be empty, got %#v instead", actual)
	}
}

func TestBoardMockGetTileAt(t *testing.T) {
	board := BoardMock{}
	_, ok := board.GetTileAt(elements.NewPosition(0, 0))
	if !ok {
		t.Fatalf("expected GetTileAt() output to be ok")
	}
}

func TestBoardMockGetLegalMovesFor(t *testing.T) {
	board := BoardMock{}
	actual := board.GetLegalMovesFor(GetTestTile())
	if len(actual) != 0 {
		t.Fatalf("expected GetLegalMovesFor() output to be empty, got %#v instead", actual)
	}
}

func TestBoardMockHasValidPlacement(t *testing.T) {
	board := BoardMock{}
	actual := board.HasValidPlacement(GetTestTile())
	expected := true
	if actual != expected {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}
}

func TestBoardMockCanBePlaced(t *testing.T) {
	board := BoardMock{}
	actual := board.CanBePlaced(
		GetTestPlacedTileWithMeeple(elements.Meeple{Side: elements.None}),
	)
	expected := true
	if actual != expected {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}
}

func TestBoardMockPlaceTileWithoutFunc(t *testing.T) {
	board := BoardMock{}
	actual, err := board.PlaceTile(GetTestPlacedTile())
	if err != nil {
		t.Fatal(err.Error())
	}
	expected := GetTestScoreReport()
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}
}

func TestBoardMockPlaceTileWithFunc(t *testing.T) {
	wasCalled := false
	expected := elements.ScoreReport{}
	board := BoardMock{
		PlaceTileFunc: func(_ elements.PlacedTile) (elements.ScoreReport, error) {
			wasCalled = true
			return expected, nil
		},
	}
	actual, err := board.PlaceTile(GetTestPlacedTile())
	if err != nil {
		t.Fatal(err.Error())
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}
	if !wasCalled {
		t.Fatal("expected PlaceTile() to be called")
	}
}
