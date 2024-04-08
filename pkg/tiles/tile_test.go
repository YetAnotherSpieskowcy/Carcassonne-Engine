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
	tile.CitiesAppendConnection(connection.Connection{Sides: []connection.Side{connection.Top, connection.Left}})
	tile.RoadsAppendConnection(connection.Connection{Sides: []connection.Side{connection.Bottom, connection.Right}})
	tile.FieldsAppendConnection(connection.Connection{Sides: []connection.Side{connection.BottomRightEdge, connection.RightBottomEdge}})
	tile.HasShield = true
	tile.Building = buildings.NoneBuilding

	var rotated = tile.Rotate(1)

	var expected Tile
	expected.CitiesAppendConnection(connection.Connection{Sides: []connection.Side{connection.Right, connection.Top}})
	expected.RoadsAppendConnection(connection.Connection{Sides: []connection.Side{connection.Left, connection.Bottom}})
	expected.FieldsAppendConnection(connection.Connection{Sides: []connection.Side{connection.LeftBottomEdge, connection.BottomLeftEdge}})
	expected.HasShield = true
	expected.Building = buildings.NoneBuilding

	if !reflect.DeepEqual(rotated, expected) {
		t.Fatalf("got %#v should be %#v", rotated.String(), expected.String())
	}
}

func TestTileToString(t *testing.T) {
	var tile Tile
	tile.CitiesAppendConnection(connection.Connection{Sides: []connection.Side{connection.Top, connection.Left}})
	tile.RoadsAppendConnection(connection.Connection{Sides: []connection.Side{connection.Bottom, connection.Right}})
	tile.FieldsAppendConnection(connection.Connection{Sides: []connection.Side{connection.BottomRightEdge, connection.RightBottomEdge}})
	tile.HasShield = true
	tile.Building = buildings.NoneBuilding

	var expected string
	expected = ""
	expected += "Cities\n"
	expected += connection.Connection{Sides: []connection.Side{connection.Top, connection.Left}}.String() + "\n"
	expected += "Roads\n"
	expected += connection.Connection{Sides: []connection.Side{connection.Bottom, connection.Right}}.String() + "\n"
	expected += "Fields\n"
	expected += connection.Connection{Sides: []connection.Side{connection.BottomRightEdge, connection.RightBottomEdge}}.String() + "\n"
	expected += "Has shields: " + strconv.FormatBool(true) + "\n"
	expected += "Building: " + buildings.NoneBuilding.String() + "\n"

	if tile.String() != expected {
		t.Fatalf("got %s \nshould be\n %s", tile.String(), expected)
	}

}
