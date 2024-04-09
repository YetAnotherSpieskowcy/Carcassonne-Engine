package tiles

import (
	"reflect"
	"testing"

	buildings "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/buildings"
	connection "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/connection"
)

func TestTileRotate(t *testing.T) {
	var tile Tile
	//  \/ da się to lepiej napisać? \/
	tile.CitiesAppendConnection([]connection.Side{connection.Top, connection.Left})
	tile.RoadsAppendConnection([]connection.Side{connection.Bottom, connection.Right})
	tile.FieldsAppendConnection([]connection.Side{connection.BottomRightEdge, connection.RightBottomEdge})
	tile.HasShield = true
	tile.Building = buildings.None

	//var rotated = tile.Rotate(1)
	tile.Rotate(1)

	var expected Tile
	expected.CitiesAppendConnection([]connection.Side{connection.Right, connection.Top})
	expected.RoadsAppendConnection([]connection.Side{connection.Left, connection.Bottom})
	expected.FieldsAppendConnection([]connection.Side{connection.LeftBottomEdge, connection.BottomLeftEdge})
	expected.HasShield = true
	expected.Building = buildings.None

	if !reflect.DeepEqual(tile, expected) {
		t.Fatalf("got\n %#v \nshould be \n%#v", tile, expected)
	}
}
