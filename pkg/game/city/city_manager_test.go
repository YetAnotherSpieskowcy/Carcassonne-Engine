package city

import (
	"reflect"
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/binarytiles"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/tiletemplates"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
)

func TestDeepClone(t *testing.T) {
	pa := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	pa.Position = position.New(1, 1)
	a := binarytiles.FromPlacedTile(pa) // todo binarytiles rewrite
	original := NewCityManager()
	original.UpdateCities(a)

	clone := original.DeepClone()

	pb := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads().Rotate(2))
	pb.Position = position.New(1, 2)
	b := binarytiles.FromPlacedTile(pb) // todo binarytiles rewrite
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

	pa := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	a := binarytiles.FromPlacedTile(pa) // todo binarytiles rewrite
	manager.UpdateCities(a)

	if len(manager.cities) != 1 {
		t.Fatalf("expected %#v, got %#v instead", 1, len(manager.cities))
	}
}

func TestUpdateCitiesWhenNoAddToExistingCity(t *testing.T) {
	pa := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	pa.Position = position.New(1, 1)
	a := binarytiles.FromPlacedTile(pa) // todo binarytiles rewrite
	manager := NewCityManager()
	manager.UpdateCities(a)

	pb := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	pb.Position = position.New(2, 1)
	b := binarytiles.FromPlacedTile(pb) // todo binarytiles rewrite
	manager.UpdateCities(b)

	if len(manager.cities) != 2 {
		t.Fatalf("expected %#v, got %#v instead", 2, len(manager.cities))
	}
}

func TestUpdateCitiesWhenAddToExistingCity(t *testing.T) {
	pa := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	pa.Position = position.New(1, 1)
	a := binarytiles.FromPlacedTile(pa) // todo binarytiles rewrite
	manager := NewCityManager()
	manager.UpdateCities(a)

	pb := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads().Rotate(2))
	pb.Position = position.New(1, 2)
	b := binarytiles.FromPlacedTile(pb) // todo binarytiles rewrite
	manager.UpdateCities(b)

	if len(manager.cities) != 1 {
		t.Fatalf("expected %#v, got %#v instead", 1, len(manager.cities))
	}
}

func TestUpdateCitiesWhenNoCityAdded(t *testing.T) {
	pa := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	pa.Position = position.New(1, 1)
	a := binarytiles.FromPlacedTile(pa) // todo binarytiles rewrite
	manager := NewCityManager()
	manager.UpdateCities(a)

	pb := elements.ToPlacedTile(tiletemplates.MonasteryWithSingleRoad())
	pb.Position = position.New(2, 1)
	b := binarytiles.FromPlacedTile(pb) // todo binarytiles rewrite
	manager.UpdateCities(b)

	if len(manager.cities) != 1 {
		t.Fatalf("expected %#v, got %#v instead", 1, len(manager.cities))
	}
}

func TestUpdateCitiesWhenOneCityClosedSeconedOpen(t *testing.T) {
	pa := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	pa.Position = position.New(1, 1)
	a := binarytiles.FromPlacedTile(pa) // todo binarytiles rewrite
	manager := NewCityManager()
	manager.UpdateCities(a)

	pb := elements.ToPlacedTile(tiletemplates.TwoCityEdgesUpAndDownNotConnected())
	pb.Position = position.New(1, 2)
	b := binarytiles.FromPlacedTile(pb) // todo binarytiles rewrite
	manager.UpdateCities(b)

	if len(manager.cities) != 2 {
		t.Fatalf("expected %#v, got %#v instead", 2, len(manager.cities))
	}
}

func TestUpdateCityWhenTwoFeaturesAdded(t *testing.T) {
	manager := NewCityManager()

	pa := elements.ToPlacedTile(tiletemplates.ThreeCityEdgesConnected().Rotate(1))
	pa.Position = position.New(1, 0)
	a := binarytiles.FromPlacedTile(pa) // todo binarytiles rewrite
	manager.UpdateCities(a)

	pb := elements.ToPlacedTile(tiletemplates.ThreeCityEdgesConnected())
	pb.Position = position.New(2, 0)
	b := binarytiles.FromPlacedTile(pb) // todo binarytiles rewrite
	manager.UpdateCities(b)

	pc := elements.ToPlacedTile(tiletemplates.FourCityEdgesConnectedShield())
	pc.Position = position.New(2, 1)
	c := binarytiles.FromPlacedTile(pc) // todo binarytiles rewrite
	manager.UpdateCities(c)

	pd := elements.ToPlacedTile(tiletemplates.TwoCityEdgesCornerNotConnected().Rotate(1))
	pd.Position = position.New(1, 1)
	d := binarytiles.FromPlacedTile(pd) // todo binarytiles rewrite
	manager.UpdateCities(d)

	if len(manager.cities) != 1 {
		t.Fatalf("expected %#v, got %#v instead", 1, len(manager.cities))
	}

	for _, fSides := range d.GetFeaturesOfType(feature.City) {
		city, _ := manager.GetCity(d.Position(), fSides)
		if city == nil {
			t.Fatalf("not found city feature at side %#v", fSides)
		}
	}
}

func TestJoinCitiesOnAdd(t *testing.T) {
	pa := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	pa.Position = position.New(1, 1)
	a := binarytiles.FromPlacedTile(pa) // todo binarytiles rewrite
	manager := NewCityManager()
	manager.UpdateCities(a)

	pb := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads().Rotate(3))
	pb.Position = position.New(2, 2)
	b := binarytiles.FromPlacedTile(pb) // todo binarytiles rewrite
	manager.UpdateCities(b)

	pc := elements.ToPlacedTile(tiletemplates.TwoCityEdgesCornerConnected().Rotate(1))
	pc.Position = position.New(1, 2)
	c := binarytiles.FromPlacedTile(pc) // todo binarytiles rewrite
	manager.UpdateCities(c)

	if len(manager.cities) != 1 {
		t.Fatalf("expected %#v, got %#v instead", 1, len(manager.cities))
	}
}

func TestJoinCitiesOnAddCityNotClosed(t *testing.T) {
	pa := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	pa.Position = position.New(1, 1)
	a := binarytiles.FromPlacedTile(pa) // todo binarytiles rewrite
	manager := NewCityManager()
	manager.UpdateCities(a)

	pb := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads().Rotate(3))
	pb.Position = position.New(2, 2)
	b := binarytiles.FromPlacedTile(pb) // todo binarytiles rewrite
	manager.UpdateCities(b)

	pc := elements.ToPlacedTile(tiletemplates.FourCityEdgesConnectedShield())
	pc.Position = position.New(1, 2)
	c := binarytiles.FromPlacedTile(pc) // todo binarytiles rewrite
	manager.UpdateCities(c)

	if len(manager.cities) != 1 {
		t.Fatalf("expected %#v, got %#v instead", 1, len(manager.cities))
	}
}

func TestJoinCitiesFourEdgeCity(t *testing.T) {
	pa1 := elements.ToPlacedTile(tiletemplates.TwoCityEdgesCornerConnected())
	pa1.Position = position.New(1, 1)
	a1 := binarytiles.FromPlacedTile(pa1) // todo binarytiles rewrite
	manager := NewCityManager()
	manager.UpdateCities(a1)

	pa2 := elements.ToPlacedTile(tiletemplates.TwoCityEdgesCornerConnectedShield().Rotate(1))
	pa2.Position = position.New(1, 2)
	a2 := binarytiles.FromPlacedTile(pa2) // todo binarytiles rewrite
	manager.UpdateCities(a2)

	pa3 := elements.ToPlacedTile(tiletemplates.ThreeCityEdgesConnectedShield())
	pa3.Position = position.New(2, 1)
	a3 := binarytiles.FromPlacedTile(pa3) // todo binarytiles rewrite
	manager.UpdateCities(a3)

	pd := elements.ToPlacedTile(tiletemplates.ThreeCityEdgesConnected())
	pd.Position = position.New(1, 3)
	d := binarytiles.FromPlacedTile(pd) // todo binarytiles rewrite
	manager.UpdateCities(d)

	pe := elements.ToPlacedTile(tiletemplates.FourCityEdgesConnectedShield())
	pe.Position = position.New(2, 2)
	e := binarytiles.FromPlacedTile(pe) // todo binarytiles rewrite
	manager.UpdateCities(e)

	if len(manager.cities) != 2 {
		t.Fatalf("expected %#v, got %#v instead", 2, len(manager.cities))
	}
}

func TestJoinCitiesFourEdgeCityTwoCitiesConnected(t *testing.T) {
	pa1 := elements.ToPlacedTile(tiletemplates.TwoCityEdgesCornerConnected())
	pa1.Position = position.New(1, 1)
	a1 := binarytiles.FromPlacedTile(pa1) // todo binarytiles rewrite
	manager := NewCityManager()
	manager.UpdateCities(a1)

	pa2 := elements.ToPlacedTile(tiletemplates.TwoCityEdgesCornerConnectedShield().Rotate(1))
	pa2.Position = position.New(1, 2)
	a2 := binarytiles.FromPlacedTile(pa2) // todo binarytiles rewrite
	manager.UpdateCities(a2)

	pa3 := elements.ToPlacedTile(tiletemplates.ThreeCityEdgesConnectedShield())
	pa3.Position = position.New(2, 1)
	a3 := binarytiles.FromPlacedTile(pa3) // todo binarytiles rewrite
	manager.UpdateCities(a3)

	pb1 := elements.ToPlacedTile(tiletemplates.TwoCityEdgesCornerConnected().Rotate(1))
	pb1.Position = position.New(2, 3)
	b1 := binarytiles.FromPlacedTile(pb1) // todo binarytiles rewrite
	manager.UpdateCities(b1)

	pb2 := elements.ToPlacedTile(tiletemplates.TwoCityEdgesCornerConnected().Rotate(2))
	pb2.Position = position.New(3, 3)
	b2 := binarytiles.FromPlacedTile(pb2) // todo binarytiles rewrite
	manager.UpdateCities(b2)

	pb3 := elements.ToPlacedTile(tiletemplates.TwoCityEdgesCornerConnected().Rotate(3))
	pb3.Position = position.New(3, 2)
	b3 := binarytiles.FromPlacedTile(pb3) // todo binarytiles rewrite
	manager.UpdateCities(b3)

	pe := elements.ToPlacedTile(tiletemplates.FourCityEdgesConnectedShield())
	pe.Position = position.New(2, 2)
	e := binarytiles.FromPlacedTile(pe) // todo binarytiles rewrite
	manager.UpdateCities(e)

	if len(manager.cities) != 1 {
		t.Fatalf("expected %#v, got %#v instead", 1, len(manager.cities))
	}
}

func TestForceScore(t *testing.T) {
	var expectedScore uint32 = 1
	var expectedMeepleType elements.MeepleType = elements.NormalMeeple
	var expectedPlayerID elements.ID = 1
	pa := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	pa.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple.PlayerID = expectedPlayerID
	pa.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple.Type = expectedMeepleType
	a := binarytiles.FromPlacedTile(pa) // todo binarytiles rewrite

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
	pa := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	pa.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple.PlayerID = expectedPlayerID
	pa.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple.Type = expectedMeepleType
	pa.Position = position.New(1, 1)
	a := binarytiles.FromPlacedTile(pa) // todo binarytiles rewrite
	manager := NewCityManager()
	manager.UpdateCities(a)

	pb := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads().Rotate(2))
	pb.Position = position.New(1, 2)
	b := binarytiles.FromPlacedTile(pb) // todo binarytiles rewrite
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
	pa := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	pa.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple.PlayerID = expectedPlayerID
	pa.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple.Type = expectedMeepleType
	pa.Position = position.New(1, 1)
	a := binarytiles.FromPlacedTile(pa) // todo binarytiles rewrite
	manager := NewCityManager()
	manager.UpdateCities(a)

	pb := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads().Rotate(2))
	pb.Position = position.New(1, 2)
	b := binarytiles.FromPlacedTile(pb) // todo binarytiles rewrite
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

	pa := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	pa.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple.PlayerID = expectedPlayerID
	pa.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple.Type = expectedMeepleType
	pa.Position = position.New(1, 1)
	a := binarytiles.FromPlacedTile(pa) // todo binarytiles rewrite
	manager := NewCityManager()
	manager.UpdateCities(a)

	pb := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads().Rotate(3))
	pb.Position = position.New(2, 2)
	b := binarytiles.FromPlacedTile(pb) // todo binarytiles rewrite
	manager.UpdateCities(b)

	pc := elements.ToPlacedTile(tiletemplates.TwoCityEdgesCornerConnected().Rotate(1))
	pc.Position = position.New(1, 2)
	c := binarytiles.FromPlacedTile(pc) // todo binarytiles rewrite
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

	pa := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	pa.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple.PlayerID = expectedPlayerID
	pa.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple.Type = expectedMeepleType
	pa.Position = position.New(1, 1)
	a := binarytiles.FromPlacedTile(pa) // todo binarytiles rewrite
	manager := NewCityManager()
	manager.UpdateCities(a)

	pb := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads().Rotate(3))
	pb.Position = position.New(2, 2)
	b := binarytiles.FromPlacedTile(pb) // todo binarytiles rewrite
	manager.UpdateCities(b)

	pc := elements.ToPlacedTile(tiletemplates.FourCityEdgesConnectedShield())
	pc.Position = position.New(1, 2)
	c := binarytiles.FromPlacedTile(pc) // todo binarytiles rewrite
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

	pa := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	pa.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple.PlayerID = expectedPlayerID
	pa.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple.Type = expectedMeepleType
	pa.Position = position.New(1, 1)
	a := binarytiles.FromPlacedTile(pa) // todo binarytiles rewrite
	manager := NewCityManager()
	manager.UpdateCities(a)

	pb := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads().Rotate(3))
	pb.Position = position.New(2, 2)
	b := binarytiles.FromPlacedTile(pb) // todo binarytiles rewrite
	manager.UpdateCities(b)

	pc := elements.ToPlacedTile(tiletemplates.FourCityEdgesConnectedShield())
	pc.Position = position.New(1, 2)
	c := binarytiles.FromPlacedTile(pc) // todo binarytiles rewrite
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

	pa := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	pa.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple.PlayerID = expectedPlayerID1
	pa.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple.Type = expectedMeepleType
	pa.Position = position.New(1, 1)
	a := binarytiles.FromPlacedTile(pa) // todo binarytiles rewrite
	manager.UpdateCities(a)

	pb := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads().Rotate(2))
	pb.GetPlacedFeatureAtSide(side.Bottom, feature.City).Meeple.PlayerID = expectedPlayerID2
	pb.GetPlacedFeatureAtSide(side.Bottom, feature.City).Meeple.Type = expectedMeepleType
	pb.Position = position.New(1, 3)
	b := binarytiles.FromPlacedTile(pb) // todo binarytiles rewrite
	manager.UpdateCities(b)

	pc := elements.ToPlacedTile(tiletemplates.TwoCityEdgesUpAndDownNotConnected())
	pc.Position = position.New(1, 2)
	c := binarytiles.FromPlacedTile(pc) // todo binarytiles rewrite
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

	pa := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	pa.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple.PlayerID = expectedPlayerID1
	pa.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple.Type = expectedMeepleType
	pa.Position = position.New(1, 1)
	a := binarytiles.FromPlacedTile(pa) // todo binarytiles rewrite
	manager.UpdateCities(a)

	pb := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads().Rotate(2))
	pb.GetPlacedFeatureAtSide(side.Bottom, feature.City).Meeple.PlayerID = expectedPlayerID2
	pb.GetPlacedFeatureAtSide(side.Bottom, feature.City).Meeple.Type = expectedMeepleType
	pb.Position = position.New(1, 3)
	b := binarytiles.FromPlacedTile(pb) // todo binarytiles rewrite
	manager.UpdateCities(b)

	pc := elements.ToPlacedTile(tiletemplates.TwoCityEdgesUpAndDownConnected())
	pc.Position = position.New(1, 2)
	c := binarytiles.FromPlacedTile(pc) // todo binarytiles rewrite
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

	pa := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	pa.Position = position.New(1, 1)
	a := binarytiles.FromPlacedTile(pa) // todo binarytiles rewrite
	manager.UpdateCities(a)

	pb := elements.ToPlacedTile(tiletemplates.TwoCityEdgesUpAndDownNotConnected())
	pb.Position = position.New(2, 1)
	b := binarytiles.FromPlacedTile(pb) // todo binarytiles rewrite
	manager.UpdateCities(b)

	aCity, aCityID := manager.GetCity(position.New(1, 1), a.GetConnectedSides(binarytiles.SideTop, feature.City))
	if !reflect.DeepEqual(*aCity, manager.cities[aCityID]) {
		t.Fatalf("expected %#v, got %#v instead", *aCity, manager.cities[aCityID])
	}

	bTopCity, bTopCityID := manager.GetCity(position.New(2, 1), b.GetConnectedSides(binarytiles.SideTop, feature.City))

	if !reflect.DeepEqual(*bTopCity, manager.cities[bTopCityID]) {
		t.Fatalf("expected %#v, got %#v instead", *bTopCity, manager.cities[bTopCityID])
	}

	bBottomCity, bBottomCityID := manager.GetCity(position.New(2, 1), b.GetConnectedSides(binarytiles.SideBottom, feature.City))

	if !reflect.DeepEqual(*bBottomCity, manager.cities[bBottomCityID]) {
		t.Fatalf("expected %#v, got %#v instead", *bBottomCity, manager.cities[bBottomCityID])
	}

	if aCityID == bTopCityID || aCityID == bBottomCityID || bTopCityID == bBottomCityID {
		t.Fatalf("expected all city IDs to be different. Got: %#v, %#v, %#v", aCityID, bTopCityID, bBottomCityID)
	}

	if len(manager.cities) != 3 {
		t.Fatalf("expected %#v, got %#v instead", 3, len(manager.cities))
	}

	nilCity, nilCityID := manager.GetCity(position.New(21, 37), a.GetConnectedSides(binarytiles.SideTop, feature.City))
	if nilCity != nil {
		t.Fatalf("expected %#v, got %#v instead", nil, nilCity)
	}
	if nilCityID != -1 {
		t.Fatalf("expected %#v, got %#v instead", -1, nilCityID)
	}
}

func TestGetCityWhenMeepleWasOnTile(t *testing.T) {
	manager := NewCityManager()

	meeple := elements.Meeple{Type: elements.NormalMeeple, PlayerID: elements.ID(1)}

	placedTile := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	placedTile.Position = position.New(1, 1)
	placedTile.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple = meeple
	tile := binarytiles.FromPlacedTile(placedTile) // todo binarytiles rewrite

	manager.UpdateCities(tile)

	city1, _ := manager.GetCity(position.New(1, 1), tile.GetConnectedSides(binarytiles.SideTop, feature.City))

	tile.RemoveMeeple()

	city, _ := manager.GetCity(position.New(1, 1), tile.GetConnectedSides(binarytiles.SideTop, feature.City))
	if !reflect.DeepEqual(city, city1) {
		t.Fatalf("expected %#v, got %#v instead", city1, city)
	}
}

func TestCanBePlacedReturnsTrueWhenOpeningNewCity(t *testing.T) {
	manager := NewCityManager()

	startingTile := elements.NewStartingTile(tilesets.StandardTileSet())
	binaryStartingTile := binarytiles.FromPlacedTile(startingTile) // todo binarytiles rewrite
	manager.UpdateCities(binaryStartingTile)

	ptile := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads().Rotate(2))
	ptile.Position = position.New(0, -1)
	tile := binarytiles.FromPlacedTile(ptile) // todo binarytiles rewrite

	expected := true
	actual := manager.CanBePlaced(tile, tile.GetConnectedSides(binarytiles.SideBottom, feature.City))

	if expected != actual {
		t.Fatalf("expected %v, got %v instead", expected, actual)
	}
}

func TestCanBePlacedReturnsTrueWhenClosingExistingCityAndOpeningNewCityWithMeeple(t *testing.T) {
	manager := NewCityManager()

	startingTile := elements.NewStartingTile(tilesets.StandardTileSet())
	binaryStartingTile := binarytiles.FromPlacedTile(startingTile) // todo binarytiles rewrite
	manager.UpdateCities(binaryStartingTile)

	ptile := elements.ToPlacedTile(tiletemplates.TwoCityEdgesUpAndDownNotConnected())
	ptile.Position = position.New(0, 1)
	feat := ptile.GetPlacedFeatureAtSide(side.Top, feature.City)
	feat.Meeple = elements.Meeple{Type: elements.NormalMeeple, PlayerID: 1}
	tile := binarytiles.FromPlacedTile(ptile) // todo binarytiles rewrite

	expected := true
	actual := manager.CanBePlaced(tile, tile.GetConnectedSides(binarytiles.SideTop, feature.City))

	if expected != actual {
		t.Fatalf("expected %v, got %v instead", expected, actual)
	}
}

func TestCanBePlacedReturnsTrueWhenClosingExistingCityAndPlacingFirstMeeple(t *testing.T) {
	manager := NewCityManager()

	startingTile := elements.NewStartingTile(tilesets.StandardTileSet())
	binaryStartingTile := binarytiles.FromPlacedTile(startingTile) // todo binarytiles rewrite
	manager.UpdateCities(binaryStartingTile)

	ptile := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads().Rotate(2))
	ptile.Position = position.New(0, 1)
	feat := ptile.GetPlacedFeatureAtSide(side.Bottom, feature.City)
	feat.Meeple = elements.Meeple{Type: elements.NormalMeeple, PlayerID: 1}
	tile := binarytiles.FromPlacedTile(ptile) // todo binarytiles rewrite

	expected := true
	actual := manager.CanBePlaced(tile, tile.GetConnectedSides(binarytiles.SideBottom, feature.City))

	if expected != actual {
		t.Fatalf("expected %v, got %v instead", expected, actual)
	}
}

func TestCanBePlacedReturnsFalseWhenClosingExistingCityAndTryingToPlaceSecondMeeple(t *testing.T) {
	manager := NewCityManager()

	pa := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	pa.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple = elements.Meeple{
		Type: elements.NormalMeeple, PlayerID: 1,
	}
	a := binarytiles.FromPlacedTile(pa) // todo binarytiles rewrite
	manager.UpdateCities(a)

	pb := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads().Rotate(2))
	pb.Position = position.New(0, 1)
	feat := pb.GetPlacedFeatureAtSide(side.Bottom, feature.City)
	feat.Meeple = elements.Meeple{Type: elements.NormalMeeple, PlayerID: 2}
	b := binarytiles.FromPlacedTile(pb) // todo binarytiles rewrite

	expected := false
	actual := manager.CanBePlaced(b, b.GetConnectedSides(binarytiles.SideBottom, feature.City))

	if expected != actual {
		t.Fatalf("expected %v, got %v instead", expected, actual)
	}
}

func TestCanBePlacedReturnsTrueWhenExpandingExistingCityAndPlacingFirstMeeple(t *testing.T) {
	manager := NewCityManager()

	startingTile := elements.NewStartingTile(tilesets.StandardTileSet())
	binaryStartingTile := binarytiles.FromPlacedTile(startingTile) // todo binarytiles rewrite
	manager.UpdateCities(binaryStartingTile)

	pb := elements.ToPlacedTile(tiletemplates.TwoCityEdgesUpAndDownConnected())
	pb.Position = position.New(0, 1)
	feat := pb.GetPlacedFeatureAtSide(side.Bottom, feature.City)
	feat.Meeple = elements.Meeple{Type: elements.NormalMeeple, PlayerID: 2}
	b := binarytiles.FromPlacedTile(pb) // todo binarytiles rewrite

	expected := true
	actual := manager.CanBePlaced(b, b.GetConnectedSides(binarytiles.SideBottom, feature.City))

	if expected != actual {
		t.Fatalf("expected %v, got %v instead", expected, actual)
	}
}

func TestCanBePlacedReturnsFalseWhenExpandingExistingCityAndTryingToPlaceSecondMeeple(t *testing.T) {
	manager := NewCityManager()

	pa := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	pa.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple = elements.Meeple{
		Type: elements.NormalMeeple, PlayerID: 1,
	}
	a := binarytiles.FromPlacedTile(pa) // todo binarytiles rewrite
	manager.UpdateCities(a)

	pb := elements.ToPlacedTile(tiletemplates.TwoCityEdgesUpAndDownConnected())
	pb.Position = position.New(0, 1)
	feat := pb.GetPlacedFeatureAtSide(side.Bottom, feature.City)
	feat.Meeple = elements.Meeple{Type: elements.NormalMeeple, PlayerID: 2}
	b := binarytiles.FromPlacedTile(pb) // todo binarytiles rewrite

	expected := false
	actual := manager.CanBePlaced(b, b.GetConnectedSides(binarytiles.SideBottom, feature.City))

	if expected != actual {
		t.Fatalf("expected %v, got %v instead", expected, actual)
	}
}
