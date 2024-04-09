package elements

import "testing"

// coverage on mocks go brrrr

func TestTileRotate(t *testing.T) {
	tile := Tile{ID: 1}
	actual := tile.Rotate(4)
	if tile != actual {
		t.Fatalf("expected %#v, got %#v instead", tile, actual)
	}
}

func TestSingleCityEdgeNoRoads(t *testing.T) {
	expected := Tile{ID: 1}
	actual := SingleCityEdgeNoRoads()
	if expected != actual {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}
}

func TestFourCityEdgesConnectedShield(t *testing.T) {
	expected := Tile{ID: 2}
	actual := FourCityEdgesConnectedShield()
	if expected != actual {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}
}

func TestGetStandardTiles(t *testing.T) {
	expected := 71
	actual := len(GetStandardTiles())
	if expected != actual {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}
}
