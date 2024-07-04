package city

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
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
	manager := NewCityManager()

	a := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	manager.UpdateCities(a)

	if len(manager.cities) != 1 {
		t.Fatalf("expected %#v, got %#v instead", 1, len(manager.cities))
	}
}

func TestUpdateCitiesWhenNoAddToExistingCity(t *testing.T) {
	a := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	a.Position = elements.NewPosition(1, 1)
	manager := NewCityManager()
	manager.UpdateCities(a)

	b := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	b.Position = elements.NewPosition(2, 1)
	manager.UpdateCities(b)

	if len(manager.cities) != 2 {
		t.Fatalf("expected %#v, got %#v instead", 2, len(manager.cities))
	}
}

func TestUpdateCitiesWhenAddToExistingCity(t *testing.T) {
	a := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	a.Position = elements.NewPosition(1, 1)
	manager := NewCityManager()
	manager.UpdateCities(a)

	b := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads().Rotate(2))
	b.Position = elements.NewPosition(1, 2)
	manager.UpdateCities(b)

	if len(manager.cities) != 1 {
		t.Fatalf("expected %#v, got %#v instead", 1, len(manager.cities))
	}
}

func TestUpdateCitiesWhenNoCityAdded(t *testing.T) {
	a := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	a.Position = elements.NewPosition(1, 1)
	manager := NewCityManager()
	manager.UpdateCities(a)

	b := elements.ToPlacedTile(tiletemplates.MonasteryWithSingleRoad())
	b.Position = elements.NewPosition(2, 1)
	manager.UpdateCities(b)

	if len(manager.cities) != 1 {
		t.Fatalf("expected %#v, got %#v instead", 1, len(manager.cities))
	}
}

func TestForceScore(t *testing.T) {
	var expectedScore uint32 = 2
	var expectedMeepleType elements.MeepleType = elements.NormalMeeple
	var expectedPlayerID elements.ID = 1
	a := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	a.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple.PlayerID = expectedPlayerID
	a.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple.MeepleType = expectedMeepleType

	manager := NewCityManager()
	manager.UpdateCities(a)
	report := manager.ScoreCities(true)
	meeples, ok := report.ReturnedMeeples[uint8(expectedPlayerID)]
	if !ok {
		t.Fatalf("expected player id not in the map")
	}

	numMeeples := meeples[expectedMeepleType]
	if numMeeples != 1 {
		t.Fatalf("expected %#v meeple, got %#v meeples instead", 1, numMeeples)
	}

	score := report.ReceivedPoints[uint8(expectedPlayerID)]
	if score != expectedScore {
		t.Fatalf("expected %#v, got %#v instead", expectedScore, score)
	}
}

func TestScore(t *testing.T) {
	var expectedScore uint32 = 4
	var expectedMeepleType elements.MeepleType = elements.NormalMeeple
	var expectedPlayerID elements.ID = 1
	a := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	a.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple.PlayerID = expectedPlayerID
	a.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple.MeepleType = expectedMeepleType
	a.Position = elements.NewPosition(1, 1)
	manager := NewCityManager()
	manager.UpdateCities(a)

	b := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads().Rotate(2))
	b.Position = elements.NewPosition(1, 2)
	manager.UpdateCities(b)

	report := manager.ScoreCities(false)
	meeples, ok := report.ReturnedMeeples[uint8(expectedPlayerID)]
	if !ok {
		t.Fatalf("expected player id not in the map")
	}

	numMeeples := meeples[expectedMeepleType]
	if numMeeples != 1 {
		t.Fatalf("expected %#v meeple, got %#v meeples instead", 1, numMeeples)
	}

	score := report.ReceivedPoints[uint8(expectedPlayerID)]
	if score != expectedScore {
		t.Fatalf("expected %#v, got %#v instead", expectedScore, score)
	}

	if len(manager.cities) != 0 {
		t.Fatalf("expected %#v, got %#v instead", 0, len(manager.cities))
	}
}
