package tile_tests

import (
	"testing"

	tile_sets "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tile_sets"
)

func TestSetStandardTiles(t *testing.T) {
	var set = tile_sets.GetStandardTiles()

	if len(set) == 71 {

		println("got")
		println(len(set), " tiles")
		println("should be")
		println(71)

		t.Fatalf(`tile ToString failed`)
	}

}
