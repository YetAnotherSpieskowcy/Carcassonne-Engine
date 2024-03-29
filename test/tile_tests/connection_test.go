package tile_tests

import (
	"reflect"
	"testing"

	connection "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/connection"
)

func TestConnectionRotate(t *testing.T) {

	var connec connection.Connection
	connec.Sides = []connection.Side{connection.RIGHT, connection.BOTTOM}

	var result connection.Connection
	result.Sides = []connection.Side{connection.RIGHT, connection.BOTTOM}

	if !reflect.DeepEqual(connec, result) {
		println("conncet ", connec.ToString())
		println("should be:")
		println("conncet ", result.ToString())
		t.Fatalf(`Connection rotation failed`)
	}
}
func TestConnectionToString(t *testing.T) {
	var connec connection.Connection
	connec.Sides = []connection.Side{connection.TOP, connection.RIGHT}

	if connec.ToString() != connection.TOP.ToString()+" "+connection.RIGHT.ToString()+" " {
		t.Fatalf(`TestConnectionToString failed`)
	}
}
