package city

import (
	"reflect"
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/tiletemplates"
)

func TestNewAndIsCompleted(t *testing.T) {
	a := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	cities := a.GetFeaturesOfType(feature.City)
	pos := position.New(1, 1)

	city := NewCity(pos, cities)

	completed := city.IsCompleted()
	if completed {
		t.Fatalf("expected %#v, got %#v instead", false, completed)
	}
}

func TestNewAndGetFeaturesFromTile(t *testing.T) {
	a := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	cities := a.GetFeaturesOfType(feature.City)
	pos := position.New(1, 1)

	city := NewCity(pos, cities)

	features, ok := city.GetFeaturesFromTile(pos)

	if ok == false {
		t.Fatalf("expected %#v, got %#v instead", true, ok)
	}
	if len(features) != len(cities) {
		t.Fatalf("expected %#v, got %#v instead", cities, features)
	}
	featureEqual := reflect.DeepEqual(cities[0], features[0])
	if !featureEqual {
		t.Fatalf("expected %#v, got %#v instead", true, featureEqual)
	}
}

func TestAddTileAndGetFeaturesFromTile(t *testing.T) {
	a := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	b := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads().Rotate(2))

	aFeatures := a.GetFeaturesOfType(feature.City)
	city := NewCity(position.New(1, 1), aFeatures)

	bFeatures := b.GetFeaturesOfType(feature.City)

	pos := position.New(1, 2)
	city.AddTile(pos, bFeatures)

	features, ok := city.GetFeaturesFromTile(pos)

	if !ok {
		t.Fatalf("expected %#v, got %#v instead", true, ok)
	}
	if len(features) != len(bFeatures) {
		t.Fatalf("expected %#v, got %#v instead", len(bFeatures), len(features))
	}
	featureEqual := reflect.DeepEqual(bFeatures[0], features[0])
	if !featureEqual {
		t.Fatalf("expected %#v, got %#v instead", true, featureEqual)
	}

}

func TestCheckCompletedWhenClosed(t *testing.T) {
	a := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	b := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads().Rotate(2))

	aFeatures := a.GetFeaturesOfType(feature.City)
	city := NewCity(position.New(1, 1), aFeatures)

	bFeatures := b.GetFeaturesOfType(feature.City)

	pos := position.New(1, 2)
	city.AddTile(pos, bFeatures)

	var expected = true
	var actual = city.IsCompleted()

	if actual != expected {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}
}

func TestCheckCompletedWhenOpen(t *testing.T) {
	a := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	b := elements.ToPlacedTile(tiletemplates.TwoCityEdgesCornerConnected())

	aFeatures := a.GetFeaturesOfType(feature.City)
	city := NewCity(position.New(1, 1), aFeatures)

	bFeatures := b.GetFeaturesOfType(feature.City)

	pos := position.New(1, 2)
	city.AddTile(pos, bFeatures)

	var expected = false
	var actual = city.IsCompleted()

	if actual != expected {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}
}

func TestScoreOneTileCity(t *testing.T) {
	var expectedPlayerID elements.ID = 1
	var expectedMeepleType elements.MeepleType = elements.NormalMeeple
	var expectedScore uint32 = 1

	a := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())

	aFeatures := []elements.PlacedFeature{}
	for _, tmp := range a.Features {
		if tmp.Type() == feature.City {
			tmp.Meeple.PlayerID = expectedPlayerID
			tmp.Meeple.Type = expectedMeepleType
			aFeatures = append(aFeatures, tmp)
		}
	}
	city := NewCity(position.New(1, 1), aFeatures)

	scoreReport := city.GetScoreReport()
	meeples, ok := scoreReport.ReturnedMeeples[expectedPlayerID]
	if !ok {
		t.Fatalf("expected player id not in the map")
	}

	numMeeples := len(meeples)
	if numMeeples != 1 {
		t.Fatalf("expected %#v meeple, got %#v meeples instead", 1, numMeeples)
	}

	score := scoreReport.ReceivedPoints[expectedPlayerID]
	if score != expectedScore {
		t.Fatalf("expected %#v, got %#v instead", expectedScore, score)
	}
}

func TestScoreOneTileCityWithShield(t *testing.T) {
	var expectedPlayerID elements.ID = 1
	var expectedMeepleType elements.MeepleType = elements.NormalMeeple
	var expectedScore uint32 = 2

	a := elements.ToPlacedTile(tiletemplates.TwoCityEdgesCornerConnectedShield())

	aFeatures := a.GetFeaturesOfType(feature.City)
	for i := range aFeatures {
		aFeatures[i].Meeple.PlayerID = expectedPlayerID
		aFeatures[i].Meeple.Type = expectedMeepleType
	}
	city := NewCity(position.New(1, 1), aFeatures)

	scoreReport := city.GetScoreReport()
	meeples, ok := scoreReport.ReturnedMeeples[expectedPlayerID]
	if !ok {
		t.Fatalf("expected player id not in the map")
	}

	numMeeples := len(meeples)
	if numMeeples != 1 {
		t.Fatalf("expected %#v meeple, got %#v meeples instead", 1, numMeeples)
	}

	score := scoreReport.ReceivedPoints[expectedPlayerID]
	if score != expectedScore {
		t.Fatalf("expected %#v, got %#v instead", expectedScore, score)
	}
}

func TestScoreThreeTilesPlusShield(t *testing.T) {
	var expectedScore uint32 = 4
	var expectedMeepleType elements.MeepleType = elements.NormalMeeple
	var expectedPlayerID elements.ID = 1

	a := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	aFeatures := a.GetFeaturesOfType(feature.City)
	for i := range aFeatures {
		aFeatures[i].Meeple.PlayerID = expectedPlayerID
		aFeatures[i].Meeple.Type = expectedMeepleType
	}
	city := NewCity(position.New(1, 1), aFeatures)

	b := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads().Rotate(3))
	bFeatures := b.GetFeaturesOfType(feature.City)
	city.AddTile(position.New(2, 2), bFeatures)

	c := elements.ToPlacedTile(tiletemplates.FourCityEdgesConnectedShield())
	cFeatures := c.GetFeaturesOfType(feature.City)
	city.AddTile(position.New(1, 2), cFeatures)

	report := city.GetScoreReport()
	meeples, ok := report.ReturnedMeeples[expectedPlayerID]
	if !ok {
		t.Fatalf("expected player id not in the map")
	}

	numMeeples := len(meeples)
	if numMeeples != 1 {
		t.Fatalf("expected %#v meeple, got %#v meeples instead", 1, numMeeples)
	}

	if report.ReceivedPoints[expectedPlayerID] != expectedScore {
		t.Fatalf("expected %#v, got %#v instead", expectedScore, report.ReceivedPoints[expectedPlayerID])
	}
}
