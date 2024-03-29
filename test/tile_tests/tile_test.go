package tile_tests

import (
	"reflect"
	"testing"

	tiles "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	Connection "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/connection"
	farm_connection "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/farm_connection"
)

func TestTileRotate(t *testing.T) {
	var tile tiles.Tile
	//  \/ da się to lepiej napisać? \/
	tile.Cities.Cities = append(tile.Cities.Cities, Connection.Connection{Connection.TOP, Connection.LEFT})
	tile.Roads.Roads = append(tile.Roads.Roads, Connection.Connection{Connection.BOTTOM, Connection.RIGHT})
	tile.Fields.Fields = append(tile.Fields.Fields, farm_connection.FarmConnection{farm_connection.BOTTOM_RIGHT, farm_connection.RIGHT_BOTTOM})
	tile.HasShield = true
	tile.Building = tiles.NONE_BULDING

	var rotated = tile.Rotate(1)

	var result tiles.Tile
	result.Cities.Cities = append(result.Cities.Cities, Connection.Connection{Connection.RIGHT, Connection.TOP})
	result.Roads.Roads = append(result.Roads.Roads, Connection.Connection{Connection.LEFT, Connection.BOTTOM})
	result.Fields.Fields = append(result.Fields.Fields, farm_connection.FarmConnection{farm_connection.LEFT_BOTTOM, farm_connection.BOTTOM_LEFT})
	result.HasShield = true
	result.Building = tiles.NONE_BULDING

	if !reflect.DeepEqual(rotated, result) {

		println("got")
		println(rotated.ToString())
		println("should be")
		println(result.ToString())

		t.Fatalf(`tile rotation failed`)
	}

}
