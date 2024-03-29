package tile_tests

import (
	"testing"

	farm_connection "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/farm_connection"
)

func TestFarmConnectionRotate(t *testing.T) {

	var connec farm_connection.FarmConnection
	connec.A = farm_connection.LEFT_TOP
	connec.B = farm_connection.RIGHT_BOTTOM

	var rotated = connec.Rotate(1)

	if rotated != farm_connection.NewFarmConnection(farm_connection.TOP_RIGHT, farm_connection.RIGHT_BOTTOM) {
		println("conncet ", rotated.A, " ", rotated.B)
		println("should be:")
		println("conncet ", farm_connection.TOP_RIGHT, " ", farm_connection.RIGHT_BOTTOM)
		t.Fatalf(`Connection rotation failed`)
	}
}
