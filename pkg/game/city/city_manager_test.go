package city

import (
	"reflect"
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/tiletemplates"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
)

func TestDeepClone(t *testing.T) {
	a := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	a.Position = position.New(1, 1)
	original := NewCityManager()
	original.UpdateCities(a)

	clone := original.DeepClone()

	b := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads().Rotate(2))
	b.Position = position.New(1, 2)
	clone.UpdateCities(b)

	if reflect.DeepEqual(original.cities[0], clone.cities[0]) {
		t.Fatalf(
			"cities from original manager (%v) and cloned manager (%v) should not be equal",
			original.cities[0],
			clone.cities[0],
		)
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
	a.Position = position.New(1, 1)
	manager := NewCityManager()
	manager.UpdateCities(a)

	b := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	b.Position = position.New(2, 1)
	manager.UpdateCities(b)

	if len(manager.cities) != 2 {
		t.Fatalf("expected %#v, got %#v instead", 2, len(manager.cities))
	}
}

func TestUpdateCitiesWhenAddToExistingCity(t *testing.T) {
	a := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	a.Position = position.New(1, 1)
	manager := NewCityManager()
	manager.UpdateCities(a)

	b := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads().Rotate(2))
	b.Position = position.New(1, 2)
	manager.UpdateCities(b)

	if len(manager.cities) != 1 {
		t.Fatalf("expected %#v, got %#v instead", 1, len(manager.cities))
	}
}

func TestUpdateCitiesWhenNoCityAdded(t *testing.T) {
	a := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	a.Position = position.New(1, 1)
	manager := NewCityManager()
	manager.UpdateCities(a)

	b := elements.ToPlacedTile(tiletemplates.MonasteryWithSingleRoad())
	b.Position = position.New(2, 1)
	manager.UpdateCities(b)

	if len(manager.cities) != 1 {
		t.Fatalf("expected %#v, got %#v instead", 1, len(manager.cities))
	}
}

func TestUpdateCitiesWhenOneCityClosedSeconedOpen(t *testing.T) {
	a := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	a.Position = position.New(1, 1)
	manager := NewCityManager()
	manager.UpdateCities(a)

	b := elements.ToPlacedTile(tiletemplates.TwoCityEdgesUpAndDownNotConnected())
	b.Position = position.New(1, 2)
	manager.UpdateCities(b)

	if len(manager.cities) != 2 {
		t.Fatalf("expected %#v, got %#v instead", 2, len(manager.cities))
	}
}

func TestUpdateCityWhenTwoFeaturesAdded(t *testing.T) {
	manager := NewCityManager()

	a := elements.ToPlacedTile(tiletemplates.ThreeCityEdgesConnected().Rotate(1))
	a.Position = position.New(1, 0)
	manager.UpdateCities(a)

	b := elements.ToPlacedTile(tiletemplates.ThreeCityEdgesConnected())
	b.Position = position.New(2, 0)
	manager.UpdateCities(b)

	c := elements.ToPlacedTile(tiletemplates.FourCityEdgesConnectedShield())
	c.Position = position.New(2, 1)
	manager.UpdateCities(c)

	d := elements.ToPlacedTile(tiletemplates.TwoCityEdgesCornerNotConnected().Rotate(1))
	d.Position = position.New(1, 1)
	manager.UpdateCities(d)

	if len(manager.cities) != 1 {
		t.Fatalf("expected %#v, got %#v instead", 1, len(manager.cities))
	}

	for _, f := range d.GetFeaturesOfType(feature.City) {
		city, _ := manager.GetCity(d.Position, f)
		if city == nil {
			t.Fatalf("not found city feature at side %#v", f.Sides)
		}
	}
}

func TestJoinCitiesOnAdd(t *testing.T) {
	a := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	a.Position = position.New(1, 1)
	manager := NewCityManager()
	manager.UpdateCities(a)

	b := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads().Rotate(3))
	b.Position = position.New(2, 2)
	manager.UpdateCities(b)

	c := elements.ToPlacedTile(tiletemplates.TwoCityEdgesCornerConnected().Rotate(1))
	c.Position = position.New(1, 2)
	manager.UpdateCities(c)

	if len(manager.cities) != 1 {
		t.Fatalf("expected %#v, got %#v instead", 1, len(manager.cities))
	}
}

func TestJoinCitiesOnAddCityNotClosed(t *testing.T) {
	a := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	a.Position = position.New(1, 1)
	manager := NewCityManager()
	manager.UpdateCities(a)

	b := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads().Rotate(3))
	b.Position = position.New(2, 2)
	manager.UpdateCities(b)

	c := elements.ToPlacedTile(tiletemplates.FourCityEdgesConnectedShield())
	c.Position = position.New(1, 2)
	manager.UpdateCities(c)

	if len(manager.cities) != 1 {
		t.Fatalf("expected %#v, got %#v instead", 1, len(manager.cities))
	}
}

func TestJoinCitiesFourEdgeCity(t *testing.T) {
	a1 := elements.ToPlacedTile(tiletemplates.TwoCityEdgesCornerConnected())
	a1.Position = position.New(1, 1)
	manager := NewCityManager()
	manager.UpdateCities(a1)

	a2 := elements.ToPlacedTile(tiletemplates.TwoCityEdgesCornerConnectedShield().Rotate(1))
	a2.Position = position.New(1, 2)
	manager.UpdateCities(a2)

	a3 := elements.ToPlacedTile(tiletemplates.ThreeCityEdgesConnectedShield())
	a3.Position = position.New(2, 1)
	manager.UpdateCities(a3)

	d := elements.ToPlacedTile(tiletemplates.ThreeCityEdgesConnected())
	d.Position = position.New(1, 3)
	manager.UpdateCities(d)

	e := elements.ToPlacedTile(tiletemplates.FourCityEdgesConnectedShield())
	e.Position = position.New(2, 2)
	manager.UpdateCities(e)

	if len(manager.cities) != 2 {
		t.Fatalf("expected %#v, got %#v instead", 2, len(manager.cities))
	}
}

func TestJoinCitiesFourEdgeCityTwoCitiesConnected(t *testing.T) {
	a1 := elements.ToPlacedTile(tiletemplates.TwoCityEdgesCornerConnected())
	a1.Position = position.New(1, 1)
	manager := NewCityManager()
	manager.UpdateCities(a1)

	a2 := elements.ToPlacedTile(tiletemplates.TwoCityEdgesCornerConnectedShield().Rotate(1))
	a2.Position = position.New(1, 2)
	manager.UpdateCities(a2)

	a3 := elements.ToPlacedTile(tiletemplates.ThreeCityEdgesConnectedShield())
	a3.Position = position.New(2, 1)
	manager.UpdateCities(a3)

	b1 := elements.ToPlacedTile(tiletemplates.TwoCityEdgesCornerConnected().Rotate(1))
	b1.Position = position.New(2, 3)
	manager.UpdateCities(b1)

	b2 := elements.ToPlacedTile(tiletemplates.TwoCityEdgesCornerConnected().Rotate(2))
	b2.Position = position.New(3, 3)
	manager.UpdateCities(b2)

	b3 := elements.ToPlacedTile(tiletemplates.TwoCityEdgesCornerConnected().Rotate(3))
	b3.Position = position.New(3, 2)
	manager.UpdateCities(b3)

	e := elements.ToPlacedTile(tiletemplates.FourCityEdgesConnectedShield())
	e.Position = position.New(2, 2)
	manager.UpdateCities(e)

	if len(manager.cities) != 1 {
		t.Fatalf("expected %#v, got %#v instead", 1, len(manager.cities))
	}
}

func TestForceScore(t *testing.T) {
	var expectedScore uint32 = 1
	var expectedMeepleType elements.MeepleType = elements.NormalMeeple
	var expectedPlayerID elements.ID = 1
	a := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	a.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple.PlayerID = expectedPlayerID
	a.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple.Type = expectedMeepleType

	manager := NewCityManager()
	manager.UpdateCities(a)
	report := manager.ScoreCities(true)
	meeples, ok := report.ReturnedMeeples[expectedPlayerID]
	if !ok {
		t.Fatalf("expected player id not in the map")
	}

	numMeeples := len(meeples)
	if numMeeples != 1 {
		t.Fatalf("expected %#v meeple, got %#v meeples instead", 1, numMeeples)
	}

	score := report.ReceivedPoints[expectedPlayerID]
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
	a.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple.Type = expectedMeepleType
	a.Position = position.New(1, 1)
	manager := NewCityManager()
	manager.UpdateCities(a)

	b := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads().Rotate(2))
	b.Position = position.New(1, 2)
	manager.UpdateCities(b)

	report := manager.ScoreCities(false)
	meeples, ok := report.ReturnedMeeples[expectedPlayerID]
	if !ok {
		t.Fatalf("expected player id not in the map")
	}

	numMeeples := len(meeples)
	if numMeeples != 1 {
		t.Fatalf("expected %#v meeple, got %#v meeples instead", 1, numMeeples)
	}

	score := report.ReceivedPoints[expectedPlayerID]
	if score != expectedScore {
		t.Fatalf("expected %#v, got %#v instead", expectedScore, score)
	}

	if len(manager.cities) != 1 {
		t.Fatalf("expected %#v, got %#v instead", 1, len(manager.cities))
	}
}

func TestScoreTwice(t *testing.T) {
	var expectedScore uint32 = 4
	var expectedScore2 uint32
	var expectedMeepleType elements.MeepleType = elements.NormalMeeple
	var expectedPlayerID elements.ID = 1
	a := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	a.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple.PlayerID = expectedPlayerID
	a.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple.Type = expectedMeepleType
	a.Position = position.New(1, 1)
	manager := NewCityManager()
	manager.UpdateCities(a)

	b := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads().Rotate(2))
	b.Position = position.New(1, 2)
	manager.UpdateCities(b)

	report := manager.ScoreCities(false)

	score := report.ReceivedPoints[expectedPlayerID]
	if score != expectedScore {
		t.Fatalf("expected %#v, got %#v instead", expectedScore, score)
	}

	if len(manager.cities) != 1 {
		t.Fatalf("expected %#v, got %#v instead", 1, len(manager.cities))
	}

	report2 := manager.ScoreCities(false)
	score2 := report2.ReceivedPoints[expectedPlayerID]
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
	a.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple.Type = expectedMeepleType
	a.Position = position.New(1, 1)
	manager := NewCityManager()
	manager.UpdateCities(a)

	b := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads().Rotate(3))
	b.Position = position.New(2, 2)
	manager.UpdateCities(b)

	c := elements.ToPlacedTile(tiletemplates.TwoCityEdgesCornerConnected().Rotate(1))
	c.Position = position.New(1, 2)
	manager.UpdateCities(c)

	report := manager.ScoreCities(false)

	if report.ReceivedPoints[expectedPlayerID] != expectedScore {
		t.Fatalf("expected %#v, got %#v instead", expectedScore, report.ReceivedPoints[expectedPlayerID])
	}

	if len(manager.cities) != 1 {
		t.Fatalf("expected %#v, got %#v instead", 1, len(manager.cities))
	}
}

func TestScoreAfterJoinNotClosed(t *testing.T) {
	var expectedScore uint32
	var expectedMeepleType elements.MeepleType = elements.NormalMeeple
	var expectedPlayerID elements.ID = 1

	a := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	a.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple.PlayerID = expectedPlayerID
	a.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple.Type = expectedMeepleType
	a.Position = position.New(1, 1)
	manager := NewCityManager()
	manager.UpdateCities(a)

	b := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads().Rotate(3))
	b.Position = position.New(2, 2)
	manager.UpdateCities(b)

	c := elements.ToPlacedTile(tiletemplates.FourCityEdgesConnectedShield())
	c.Position = position.New(1, 2)
	manager.UpdateCities(c)

	report := manager.ScoreCities(false)

	if report.ReceivedPoints[expectedPlayerID] != expectedScore {
		t.Fatalf("expected %#v, got %#v instead", expectedScore, report.ReceivedPoints[expectedPlayerID])
	}

	if len(manager.cities) != 1 {
		t.Fatalf("expected %#v, got %#v instead", 1, len(manager.cities))
	}
}

func TestForceScoreAfterJoinNotClosedWithShield(t *testing.T) {
	var expectedScore uint32 = 4
	var expectedMeepleType elements.MeepleType = elements.NormalMeeple
	var expectedPlayerID elements.ID = 1

	a := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	a.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple.PlayerID = expectedPlayerID
	a.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple.Type = expectedMeepleType
	a.Position = position.New(1, 1)
	manager := NewCityManager()
	manager.UpdateCities(a)

	b := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads().Rotate(3))
	b.Position = position.New(2, 2)
	manager.UpdateCities(b)

	c := elements.ToPlacedTile(tiletemplates.FourCityEdgesConnectedShield())
	c.Position = position.New(1, 2)
	manager.UpdateCities(c)

	report := manager.ScoreCities(true)

	if report.ReceivedPoints[expectedPlayerID] != expectedScore {
		t.Fatalf("expected %#v, got %#v instead", expectedScore, report.ReceivedPoints[expectedPlayerID])
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
	a.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple.Type = expectedMeepleType
	a.Position = position.New(1, 1)
	manager.UpdateCities(a)

	b := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads().Rotate(2))
	b.GetPlacedFeatureAtSide(side.Bottom, feature.City).Meeple.PlayerID = expectedPlayerID2
	b.GetPlacedFeatureAtSide(side.Bottom, feature.City).Meeple.Type = expectedMeepleType
	b.Position = position.New(1, 3)
	manager.UpdateCities(b)

	c := elements.ToPlacedTile(tiletemplates.TwoCityEdgesUpAndDownNotConnected())
	c.Position = position.New(1, 2)
	manager.UpdateCities(c)

	report := manager.ScoreCities(false)
	if report.ReceivedPoints[expectedPlayerID1] != expectedScore {
		t.Fatalf("expected %#v for player %#v, got %#v instead", expectedScore, expectedPlayerID1, report.ReceivedPoints[expectedPlayerID1])
	}

	if report.ReceivedPoints[expectedPlayerID2] != expectedScore {
		t.Fatalf("expected %#v for player %#v, got %#v instead", expectedScore, expectedPlayerID2, report.ReceivedPoints[expectedPlayerID2])
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
	a.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple.Type = expectedMeepleType
	a.Position = position.New(1, 1)
	manager.UpdateCities(a)

	b := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads().Rotate(2))
	b.GetPlacedFeatureAtSide(side.Bottom, feature.City).Meeple.PlayerID = expectedPlayerID2
	b.GetPlacedFeatureAtSide(side.Bottom, feature.City).Meeple.Type = expectedMeepleType
	b.Position = position.New(1, 3)
	manager.UpdateCities(b)

	c := elements.ToPlacedTile(tiletemplates.TwoCityEdgesUpAndDownConnected())
	c.Position = position.New(1, 2)
	manager.UpdateCities(c)

	report := manager.ScoreCities(false)
	if report.ReceivedPoints[expectedPlayerID1] != expectedScore {
		t.Fatalf("expected %#v for player %#v, got %#v instead", expectedScore, expectedPlayerID1, report.ReceivedPoints[expectedPlayerID1])
	}

	if report.ReceivedPoints[expectedPlayerID2] != expectedScore {
		t.Fatalf("expected %#v for player %#v, got %#v instead", expectedScore, expectedPlayerID2, report.ReceivedPoints[expectedPlayerID2])
	}

	if len(manager.cities) != 1 {
		t.Fatalf("expected %#v, got %#v instead", 1, len(manager.cities))
	}
}

func TestGetCity(t *testing.T) {
	manager := NewCityManager()

	a := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	a.Position = position.New(1, 1)
	manager.UpdateCities(a)

	b := elements.ToPlacedTile(tiletemplates.TwoCityEdgesUpAndDownNotConnected())
	b.Position = position.New(2, 1)
	manager.UpdateCities(b)

	aCity, aCityID := manager.GetCity(position.New(1, 1), *a.GetPlacedFeatureAtSide(side.Top, feature.City))
	if !reflect.DeepEqual(*aCity, manager.cities[aCityID]) {
		t.Fatalf("expected %#v, got %#v instead", *aCity, manager.cities[aCityID])
	}

	bTopCity, bTopCityID := manager.GetCity(position.New(2, 1), *b.GetPlacedFeatureAtSide(side.Top, feature.City))

	if !reflect.DeepEqual(*bTopCity, manager.cities[bTopCityID]) {
		t.Fatalf("expected %#v, got %#v instead", *bTopCity, manager.cities[bTopCityID])
	}

	bBottomCity, bBottomCityID := manager.GetCity(position.New(2, 1), *b.GetPlacedFeatureAtSide(side.Bottom, feature.City))

	if !reflect.DeepEqual(*bBottomCity, manager.cities[bBottomCityID]) {
		t.Fatalf("expected %#v, got %#v instead", *bBottomCity, manager.cities[bBottomCityID])
	}

	if aCityID == bTopCityID || aCityID == bBottomCityID || bTopCityID == bBottomCityID {
		t.Fatalf("expected all city IDs to be different. Got: %#v, %#v, %#v", aCityID, bTopCityID, bBottomCityID)
	}

	if len(manager.cities) != 3 {
		t.Fatalf("expected %#v, got %#v instead", 3, len(manager.cities))
	}

	nilCity, nilCityID := manager.GetCity(position.New(21, 37), *a.GetPlacedFeatureAtSide(side.Top, feature.City))
	if nilCity != nil {
		t.Fatalf("expected %#v, got %#v instead", nil, nilCity)
	}
	if nilCityID != -1 {
		t.Fatalf("expected %#v, got %#v instead", -1, nilCityID)
	}
}

func TestCanBePlacedReturnsTrueWhenOpeningNewCity(t *testing.T) {
	manager := NewCityManager()

	startingTile := elements.NewStartingTile(tilesets.StandardTileSet())
	manager.UpdateCities(startingTile)

	ptile := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads().Rotate(2))
	ptile.Position = position.New(0, -1)

	expected := true
	actual := manager.CanBePlaced(ptile, *ptile.GetPlacedFeatureAtSide(side.Bottom, feature.City))

	if expected != actual {
		t.Fatalf("expected %v, got %v instead", expected, actual)
	}
}

func TestCanBePlacedReturnsTrueWhenClosingExistingCityAndOpeningNewCityWithMeeple(t *testing.T) {
	manager := NewCityManager()

	startingTile := elements.NewStartingTile(tilesets.StandardTileSet())
	manager.UpdateCities(startingTile)

	ptile := elements.ToPlacedTile(tiletemplates.TwoCityEdgesUpAndDownNotConnected())
	ptile.Position = position.New(0, 1)
	feat := ptile.GetPlacedFeatureAtSide(side.Top, feature.City)
	feat.Meeple = elements.Meeple{Type: elements.NormalMeeple, PlayerID: 1}

	expected := true
	actual := manager.CanBePlaced(ptile, *feat)

	if expected != actual {
		t.Fatalf("expected %v, got %v instead", expected, actual)
	}
}

func TestCanBePlacedReturnsTrueWhenClosingExistingCityAndPlacingFirstMeeple(t *testing.T) {
	manager := NewCityManager()

	startingTile := elements.NewStartingTile(tilesets.StandardTileSet())
	manager.UpdateCities(startingTile)

	ptile := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads().Rotate(2))
	ptile.Position = position.New(0, 1)
	feat := ptile.GetPlacedFeatureAtSide(side.Bottom, feature.City)
	feat.Meeple = elements.Meeple{Type: elements.NormalMeeple, PlayerID: 1}

	expected := true
	actual := manager.CanBePlaced(ptile, *feat)

	if expected != actual {
		t.Fatalf("expected %v, got %v instead", expected, actual)
	}
}

func TestCanBePlacedReturnsFalseWhenClosingExistingCityAndTryingToPlaceSecondMeeple(t *testing.T) {
	manager := NewCityManager()

	a := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	a.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple = elements.Meeple{
		Type: elements.NormalMeeple, PlayerID: 1,
	}
	manager.UpdateCities(a)

	b := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads().Rotate(2))
	b.Position = position.New(0, 1)
	feat := b.GetPlacedFeatureAtSide(side.Bottom, feature.City)
	feat.Meeple = elements.Meeple{Type: elements.NormalMeeple, PlayerID: 2}

	expected := false
	actual := manager.CanBePlaced(b, *feat)

	if expected != actual {
		t.Fatalf("expected %v, got %v instead", expected, actual)
	}
}

func TestCanBePlacedReturnsTrueWhenExpandingExistingCityAndPlacingFirstMeeple(t *testing.T) {
	manager := NewCityManager()

	startingTile := elements.NewStartingTile(tilesets.StandardTileSet())
	manager.UpdateCities(startingTile)

	b := elements.ToPlacedTile(tiletemplates.TwoCityEdgesUpAndDownConnected())
	b.Position = position.New(0, 1)
	feat := b.GetPlacedFeatureAtSide(side.Bottom, feature.City)
	feat.Meeple = elements.Meeple{Type: elements.NormalMeeple, PlayerID: 2}

	expected := true
	actual := manager.CanBePlaced(b, *feat)

	if expected != actual {
		t.Fatalf("expected %v, got %v instead", expected, actual)
	}
}

func TestCanBePlacedReturnsFalseWhenExpandingExistingCityAndTryingToPlaceSecondMeeple(t *testing.T) {
	manager := NewCityManager()

	a := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	a.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple = elements.Meeple{
		Type: elements.NormalMeeple, PlayerID: 1,
	}
	manager.UpdateCities(a)

	b := elements.ToPlacedTile(tiletemplates.TwoCityEdgesUpAndDownConnected())
	b.Position = position.New(0, 1)
	feat := b.GetPlacedFeatureAtSide(side.Bottom, feature.City)
	feat.Meeple = elements.Meeple{Type: elements.NormalMeeple, PlayerID: 2}

	expected := false
	actual := manager.CanBePlaced(b, *feat)

	if expected != actual {
		t.Fatalf("expected %v, got %v instead", expected, actual)
	}
}
