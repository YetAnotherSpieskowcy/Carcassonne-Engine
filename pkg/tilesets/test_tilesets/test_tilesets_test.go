package test_tilesets

import (
	"testing"
)

func TestMiniTileSet1(t *testing.T) {
	var set = OrderedMiniTileSet1()
	expected := 12

	actual := len(set.Tiles)

	if expected != actual {
		t.Fatalf("got %#v tiles, should be %#v", actual, expected)
	}
}

func TestMiniTileSet2(t *testing.T) {
	var set = OrderedMiniTileSet2()
	expected := 12

	actual := len(set.Tiles)

	if expected != actual {
		t.Fatalf("got %#v tiles, should be %#v", actual, expected)
	}
}

func TestEveryTileOnceTileSet(t *testing.T) {
	var set = EveryTileOnceTileSet()
	expected := 24

	actual := len(set.Tiles)

	if expected != actual {
		t.Fatalf("got %#v tiles, should be %#v", actual, expected)
	}
}
