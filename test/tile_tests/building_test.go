package tile_tests

import (
	"testing"

	buildings "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/buildings"
)

func TestBuildingToString(t *testing.T) {

	if buildings.NONE_BULDING.ToString() != "NONE_BUILDING" {
		t.Fatalf(`NONE_BULDING toString failed`)
	}

	if buildings.MONASTERY.ToString() != "MONASTERY" {
		t.Fatalf(`MONASTERY toString failed`)
	}

	if buildings.Bulding(100).ToString() != "ERROR" {
		t.Fatalf(`ERROR toString failed`)
	}
}
