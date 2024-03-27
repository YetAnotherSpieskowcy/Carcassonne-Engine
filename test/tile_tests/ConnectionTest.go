package tile_tests

import (
	"testing"

	tiles "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
)

func TestSideRotate(t *testing.T) {

	if tiles.NONE.Rotate(1) != tiles.NONE {
		t.Fatalf(`Hello("Gladys")`)
	}
}
