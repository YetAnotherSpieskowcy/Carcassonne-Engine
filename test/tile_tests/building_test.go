package tile_tests

import (
	"testing"

	tiles "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
)

func TestBuildingToString(t *testing.T) {

	if tiles.NONE_BULDING.ToString() != "NONE_BUILDING" {
		t.Fatalf(`NONE_BULDING toString failed`)
	}

	if tiles.MONASTERY.ToString() != "MONASTERY" {
		t.Fatalf(`MONASTERY toString failed`)
	}

	if tiles.Bulding(100).ToString() != "ERROR" {
		t.Fatalf(`ERROR toString failed`)
	}
}
