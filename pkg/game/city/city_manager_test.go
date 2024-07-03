package city

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/tiletemplates"
)

func TestGetNeighbouringPositions(t *testing.T) {
	positions := getNeighbouringPositions(elements.NewPosition(1, 1))

	topPosition := positions[side.Top]
	if topPosition.X() != 1 || topPosition.Y() != 2 {
		t.Fatalf("expected x=%#v y=%#v, got x=%#v y=%#v instead", 1, 2, topPosition.X(), topPosition.Y())
	}
}

func TestUpdateCitiesWhenNoCities(t *testing.T) {
	a := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	manager := NewCityManager()

	manager.UpdateCities(a)

	if len(manager.cities) != 1 {
		t.Fatalf("expected %#v, got %#v instead", 1, len(manager.cities))
	}
}
