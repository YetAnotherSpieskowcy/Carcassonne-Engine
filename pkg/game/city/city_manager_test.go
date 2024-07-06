package city

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/tiletemplates"
)

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

func TestJoinCitiesOnAdd(t *testing.T) {
	a := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	a.Position = elements.NewPosition(1, 1)
	manager := NewCityManager()
	manager.UpdateCities(a)

	b := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads().Rotate(3))
	b.Position = elements.NewPosition(2, 2)
	manager.UpdateCities(b)

	c := elements.ToPlacedTile(tiletemplates.TwoCityEdgesCornerConnected().Rotate(1))
	c.Position = elements.NewPosition(1, 2)
	manager.UpdateCities(c)

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

	if len(manager.cities) != 1 {
		t.Fatalf("expected %#v, got %#v instead", 1, len(manager.cities))
	}
}

func TestScoreTwice(t *testing.T) {
	var expectedScore uint32 = 4
	var expectedScore2 uint32 = 0
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

	score := report.ReceivedPoints[uint8(expectedPlayerID)]
	if score != expectedScore {
		t.Fatalf("expected %#v, got %#v instead", expectedScore, score)
	}

	if len(manager.cities) != 1 {
		t.Fatalf("expected %#v, got %#v instead", 1, len(manager.cities))
	}

	report2 := manager.ScoreCities(false)
	score2 := report2.ReceivedPoints[uint8(expectedPlayerID)]
	if score2 != expectedScore2 {
		t.Fatalf("expected %#v, got %#v instead", expectedScore2, score2)
	}
}

func TestScoreAfterJoin(t *testing.T) {
	var expectedScore uint32 = 6
	var expectedMeepleType elements.MeepleType = elements.NormalMeeple
	var expectedPlayerID elements.ID = 1

	a := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	a.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple.PlayerID = expectedPlayerID
	a.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple.MeepleType = expectedMeepleType
	a.Position = elements.NewPosition(1, 1)
	manager := NewCityManager()
	manager.UpdateCities(a)

	b := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads().Rotate(3))
	b.Position = elements.NewPosition(2, 2)
	manager.UpdateCities(b)

	c := elements.ToPlacedTile(tiletemplates.TwoCityEdgesCornerConnected().Rotate(1))
	c.Position = elements.NewPosition(1, 2)
	manager.UpdateCities(c)

	report := manager.ScoreCities(false)

	if report.ReceivedPoints[uint8(expectedPlayerID)] != expectedScore {
		t.Fatalf("expected %#v, got %#v instead", expectedScore, report.ReceivedPoints[uint8(expectedPlayerID)])
	}

	if len(manager.cities) != 1 {
		t.Fatalf("expected %#v, got %#v instead", 1, len(manager.cities))
	}
}

func TestScoreTwoCitiesNotConnected(t *testing.T) {
	var expectedScore uint32 = 4
	var expectedMeepleType elements.MeepleType = elements.NormalMeeple
	var expectedPlayerID1 elements.ID = 1
	var expectedPlayerID2 elements.ID = 2

	manager := NewCityManager()

	a := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	a.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple.PlayerID = expectedPlayerID1
	a.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple.MeepleType = expectedMeepleType
	a.Position = elements.NewPosition(1, 1)
	manager.UpdateCities(a)

	b := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads().Rotate(2))
	b.GetPlacedFeatureAtSide(side.Bottom, feature.City).Meeple.PlayerID = expectedPlayerID2
	b.GetPlacedFeatureAtSide(side.Bottom, feature.City).Meeple.MeepleType = expectedMeepleType
	b.Position = elements.NewPosition(1, 3)
	manager.UpdateCities(b)

	c := elements.ToPlacedTile(tiletemplates.TwoCityEdgesUpAndDownNotConnected())
	c.Position = elements.NewPosition(1, 2)
	manager.UpdateCities(c)

	report := manager.ScoreCities(false)
	if report.ReceivedPoints[uint8(expectedPlayerID1)] != expectedScore {
		t.Fatalf("expected %#v for player %#v, got %#v instead", expectedScore, expectedPlayerID1, report.ReceivedPoints[uint8(expectedPlayerID1)])
	}

	if report.ReceivedPoints[uint8(expectedPlayerID2)] != expectedScore {
		t.Fatalf("expected %#v for player %#v, got %#v instead", expectedScore, expectedPlayerID2, report.ReceivedPoints[uint8(expectedPlayerID2)])
	}

	if len(manager.cities) != 2 {
		t.Fatalf("expected %#v, got %#v instead", 2, len(manager.cities))
	}
}

func TestScoreTwoPlayersCityConnected(t *testing.T) {
	var expectedScore uint32 = 6
	var expectedMeepleType elements.MeepleType = elements.NormalMeeple
	var expectedPlayerID1 elements.ID = 1
	var expectedPlayerID2 elements.ID = 2

	manager := NewCityManager()

	a := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	a.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple.PlayerID = expectedPlayerID1
	a.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple.MeepleType = expectedMeepleType
	a.Position = elements.NewPosition(1, 1)
	manager.UpdateCities(a)

	b := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads().Rotate(2))
	b.GetPlacedFeatureAtSide(side.Bottom, feature.City).Meeple.PlayerID = expectedPlayerID2
	b.GetPlacedFeatureAtSide(side.Bottom, feature.City).Meeple.MeepleType = expectedMeepleType
	b.Position = elements.NewPosition(1, 3)
	manager.UpdateCities(b)

	c := elements.ToPlacedTile(tiletemplates.TwoCityEdgesUpAndDownConnected())
	c.Position = elements.NewPosition(1, 2)
	manager.UpdateCities(c)

	report := manager.ScoreCities(false)
	if report.ReceivedPoints[uint8(expectedPlayerID1)] != expectedScore {
		t.Fatalf("expected %#v for player %#v, got %#v instead", expectedScore, expectedPlayerID1, report.ReceivedPoints[uint8(expectedPlayerID1)])
	}

	if report.ReceivedPoints[uint8(expectedPlayerID2)] != expectedScore {
		t.Fatalf("expected %#v for player %#v, got %#v instead", expectedScore, expectedPlayerID2, report.ReceivedPoints[uint8(expectedPlayerID2)])
	}

	if len(manager.cities) != 1 {
		t.Fatalf("expected %#v, got %#v instead", 1, len(manager.cities))
	}
}
