package tilesets

import (
	"testing"
)

// reference for sets tiles amount https://docs.google.com/spreadsheets/d/1TnPvB6oyisNGs7GZ0xpu-3LPp1V5-t0xH4vocCUPvsY/edit#gid=0

func TestStandardTileSet(t *testing.T) {
	var set = StandardTileSet()
	expected := 71

	actual := len(set.Tiles)

	if expected != actual {
		t.Fatalf("got %#v tiles, should be %#v", actual, expected)
	}
}

func TestMiniTileSet(t *testing.T) {
	var set = MiniTileSet()
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
