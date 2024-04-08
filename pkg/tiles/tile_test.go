package tiles

import (
	"reflect"
	"strconv"
	"testing"

	buildings "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/buildings"
	connection "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/connection"
)

func TestTileRotate(t *testing.T) {
	var tile Tile
	//  \/ da się to lepiej napisać? \/
	tile.CitiesAppendConnection(connection.Connection{Sides: []connection.Side{connection.TOP, connection.LEFT}})
	tile.RoadsAppendConnection(connection.Connection{Sides: []connection.Side{connection.BOTTOM, connection.RIGHT}})
	tile.FieldsAppendConnection(connection.Connection{Sides: []connection.Side{connection.BOTTOM_RIGHT_EDGE, connection.RIGHT_BOTTOM_EDGE}})
	tile.HasShield = true
	tile.Building = buildings.NONE_BULDING

	var rotated = tile.Rotate(1)

	var expected Tile
	expected.CitiesAppendConnection(connection.Connection{Sides: []connection.Side{connection.RIGHT, connection.TOP}})
	expected.RoadsAppendConnection(connection.Connection{Sides: []connection.Side{connection.LEFT, connection.BOTTOM}})
	expected.FieldsAppendConnection(connection.Connection{Sides: []connection.Side{connection.LEFT_BOTTOM_EDGE, connection.BOTTOM_LEFT_EDGE}})
	expected.HasShield = true
	expected.Building = buildings.NONE_BULDING

	if !reflect.DeepEqual(rotated, expected) {
		t.Fatalf("got %#v should be %#v", rotated.String(), expected.String())
	}
}

func TestTileToString(t *testing.T) {
	var tile Tile
	tile.CitiesAppendConnection(connection.Connection{[]connection.Side{connection.TOP, connection.LEFT}})
	tile.RoadsAppendConnection(connection.Connection{[]connection.Side{connection.BOTTOM, connection.RIGHT}})
	tile.FieldsAppendConnection(connection.Connection{[]connection.Side{connection.BOTTOM_RIGHT_EDGE, connection.RIGHT_BOTTOM_EDGE}})
	tile.HasShield = true
	tile.Building = buildings.NONE_BULDING

	var expected string
	expected = ""
	expected += "Cities\n"
	expected += connection.Connection{[]connection.Side{connection.TOP, connection.LEFT}}.String() + "\n"
	expected += "Roads\n"
	expected += connection.Connection{[]connection.Side{connection.BOTTOM, connection.RIGHT}}.String() + "\n"
	expected += "Fields\n"
	expected += connection.Connection{[]connection.Side{connection.BOTTOM_RIGHT_EDGE, connection.RIGHT_BOTTOM_EDGE}}.String() + "\n"
	expected += "Has shields: " + strconv.FormatBool(true) + "\n"
	expected += "Building: " + buildings.NONE_BULDING.String() + "\n"

	if tile.String() != expected {
		t.Fatalf("got %s \nshould be\n %s", tile.String(), expected)
	}

}
