package requestperformancetests

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/tiletemplates"
)

func TestCreateEarlyGameEngine(t *testing.T) {
	eng, game := CreateEarlyGameEngine(t.TempDir())
	if eng == nil {
		t.Fatalf("Engine is null")
	}

	expectedTile := tiletemplates.SingleCityEdgeCrossRoad() // 11th turn is B
	if !game.Game.CurrentTile.Equals(expectedTile) {
		t.Fatalf("Wrong current tile. Actual current tile: %#v \n expected: %#v", game.Game.CurrentTile, expectedTile)
	}
}

func TestCreateLateGameEngine(t *testing.T) {
	eng, game := CreateLateGameEngine(t.TempDir())
	if eng == nil {
		t.Fatalf("Engine is null")
	}

	expectedTile := tiletemplates.ThreeCityEdgesConnectedShield() // 21st turn is L
	if !game.Game.CurrentTile.Equals(expectedTile) {
		t.Fatalf("Wrong current tile. Actual current tile: %#v \n expected: %#v", game.Game.CurrentTile, expectedTile)
	}
}
