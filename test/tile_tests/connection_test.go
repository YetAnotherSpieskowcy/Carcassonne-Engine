package tile_tests

import (
	"testing"

	connection "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/connection"
)

func TestConnectionRotate(t *testing.T) {

	var connec connection.Connection
	connec.A = connection.TOP
	connec.B = connection.RIGHT

	var rotated = connec.Rotate(1)

	if rotated != connection.NewConnection(connection.RIGHT, connection.BOTTOM) {
		println("conncet ", rotated.A, " ", rotated.B)
		println("should be:")
		println("conncet ", connection.RIGHT, " ", connection.BOTTOM)
		t.Fatalf(`Connection rotation failed`)
	}
}
