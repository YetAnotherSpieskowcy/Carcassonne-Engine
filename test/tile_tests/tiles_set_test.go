package tile_tests

import (
	"testing"

	tiles "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
)

func TestSetStandardTiles(t *testing.T) {
	var set []tiles.Tile
	set = tiles.GetStandardTiles()

	if len(set) == 71 {

		println("got")
		println(len(set), " tiles")
		println("should be")
		println(71)

		t.Fatalf(`tile ToString failed`)
	}

}
