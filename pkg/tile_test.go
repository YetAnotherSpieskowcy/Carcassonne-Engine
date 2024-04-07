package tile_tests

import (
	"reflect"
	"strconv"
	"testing"

	tiles "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	buildings "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/buildings"
	Connection "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/connection"
)

func TestTileRotate(t *testing.T) {
	var tile tiles.Tile
	//  \/ da się to lepiej napisać? \/
	tile.Cities.Cities = append(tile.Cities.Cities, Connection.Connection{Sides: []Connection.Side{Connection.TOP, Connection.LEFT}})
	tile.Roads.Roads = append(tile.Roads.Roads, Connection.Connection{Sides: []Connection.Side{Connection.BOTTOM, Connection.RIGHT}})
	tile.Fields.Fields = append(tile.Fields.Fields, Connection.Connection{Sides: []Connection.Side{Connection.BOTTOM_RIGHT_EDGE, Connection.RIGHT_BOTTOM_EDGE}})
	tile.HasShield = true
	tile.Building = buildings.NONE_BULDING

	var rotated = tile.Rotate(1)

	var expected tiles.Tile
	expected.Cities.Cities = append(expected.Cities.Cities, Connection.Connection{Sides: []Connection.Side{Connection.RIGHT, Connection.TOP}})
	expected.Roads.Roads = append(expected.Roads.Roads, Connection.Connection{Sides: []Connection.Side{Connection.LEFT, Connection.BOTTOM}})
	expected.Fields.Fields = append(expected.Fields.Fields, Connection.Connection{Sides: []Connection.Side{Connection.LEFT_BOTTOM_EDGE, Connection.BOTTOM_LEFT_EDGE}})
	expected.HasShield = true
	expected.Building = buildings.NONE_BULDING

	if !reflect.DeepEqual(rotated, expected) {
		t.Fatalf("got %#v should be %#v", rotated.String(), expected.String())
	}
}

func TestTileToString(t *testing.T) {
	var tile tiles.Tile
	tile.Cities.Cities = append(tile.Cities.Cities, Connection.Connection{[]Connection.Side{Connection.TOP, Connection.LEFT}})
	tile.Roads.Roads = append(tile.Roads.Roads, Connection.Connection{[]Connection.Side{Connection.BOTTOM, Connection.RIGHT}})
	tile.Fields.Fields = append(tile.Fields.Fields, Connection.Connection{[]Connection.Side{Connection.BOTTOM_RIGHT_EDGE, Connection.RIGHT_BOTTOM_EDGE}})
	tile.HasShield = true
	tile.Building = buildings.NONE_BULDING

	var expected string
	expected = ""
	expected += "Cities\n"
	expected += Connection.Connection{[]Connection.Side{Connection.TOP, Connection.LEFT}}.String() + "\n"
	expected += "Roads\n"
	expected += Connection.Connection{[]Connection.Side{Connection.BOTTOM, Connection.RIGHT}}.String() + "\n"
	expected += "Fields\n"
	expected += Connection.Connection{[]Connection.Side{Connection.BOTTOM_RIGHT_EDGE, Connection.RIGHT_BOTTOM_EDGE}}.String() + "\n"
	expected += "Has shields: " + strconv.FormatBool(true) + "\n"
	expected += "Building: " + buildings.NONE_BULDING.String() + "\n"

	if tile.String() != expected {
		t.Fatalf("got %#v should be %#v", tile.String(), expected)
	}

}
