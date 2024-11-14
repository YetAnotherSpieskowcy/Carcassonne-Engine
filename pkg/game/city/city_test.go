package city

import (
	"reflect"
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/tiletemplates"
)

func TestNewAndIsCompleted(t *testing.T) {
	pa := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	pa.Position = position.New(1, 1)
	a := elements.BinaryTileFromPlacedTile(pa) // todo binarytiles rewrite
	cities := a.GetFeaturesOfType(feature.City)

	city := NewCity(a, cities)

	completed := city.IsCompleted()
	if completed {
		t.Fatalf("expected %#v, got %#v instead", false, completed)
	}
}

func TestNewAndGetFeaturesFromTile(t *testing.T) {
	pa := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	pa.Position = position.New(1, 1)
	a := elements.BinaryTileFromPlacedTile(pa) // todo binarytiles rewrite
	cities := a.GetFeaturesOfType(feature.City)

	city := NewCity(a, cities)

	features, ok := city.GetFeaturesFromTile(a.Position())

	if ok == false {
		t.Fatalf("expected %#v, got %#v instead", true, ok)
	}
	if len(features) != len(cities) {
		t.Fatalf("expected %#v, got %#v instead", cities, features)
	}
	featureEqual := reflect.DeepEqual(cities[0], features[0].Side)
	if !featureEqual {
		t.Fatalf("expected %#v, got %#v instead", true, featureEqual)
	}
}

func TestAddTileAndGetFeaturesFromTile(t *testing.T) {
	pa := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	pa.Position = position.New(1, 1)
	a := elements.BinaryTileFromPlacedTile(pa) // todo binarytiles rewrite

	pb := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads().Rotate(2))
	pb.Position = position.New(1, 2)
	b := elements.BinaryTileFromPlacedTile(pb) // todo binarytiles rewrite

	aFeatures := a.GetFeaturesOfType(feature.City)
	city := NewCity(a, aFeatures)

	bFeatures := b.GetFeaturesOfType(feature.City)

	city.AddTile(b, bFeatures)

	features, ok := city.GetFeaturesFromTile(b.Position())

	if !ok {
		t.Fatalf("expected %#v, got %#v instead", true, ok)
	}
	if len(features) != len(bFeatures) {
		t.Fatalf("expected %#v, got %#v instead", len(bFeatures), len(features))
	}
	featureEqual := reflect.DeepEqual(bFeatures[0], features[0].Side)
	if !featureEqual {
		t.Fatalf("expected %#v, got %#v instead", true, featureEqual)
	}

}

func TestCheckCompletedWhenClosed(t *testing.T) {
	pa := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	pa.Position = position.New(1, 1)
	a := elements.BinaryTileFromPlacedTile(pa) // todo binarytiles rewrite

	pb := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads().Rotate(2))
	pb.Position = position.New(1, 2)
	b := elements.BinaryTileFromPlacedTile(pb) // todo binarytiles rewrite

	aFeatures := a.GetFeaturesOfType(feature.City)
	city := NewCity(a, aFeatures)

	bFeatures := b.GetFeaturesOfType(feature.City)

	city.AddTile(b, bFeatures)

	var expected = true
	var actual = city.IsCompleted()

	if actual != expected {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}
}

func TestCheckCompletedWhenOpen(t *testing.T) {
	pa := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	pa.Position = position.New(1, 1)
	a := elements.BinaryTileFromPlacedTile(pa) // todo binarytiles rewrite
	pb := elements.ToPlacedTile(tiletemplates.TwoCityEdgesCornerConnected())
	pb.Position = position.New(1, 2)
	b := elements.BinaryTileFromPlacedTile(pb) // todo binarytiles rewrite

	aFeatures := a.GetFeaturesOfType(feature.City)
	city := NewCity(a, aFeatures)

	bFeatures := b.GetFeaturesOfType(feature.City)

	city.AddTile(b, bFeatures)

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

	pa := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	pa.Position = position.New(1, 1)
	pa.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple = elements.Meeple{PlayerID: expectedPlayerID, Type: expectedMeepleType}

	a := elements.BinaryTileFromPlacedTile(pa) // todo binarytiles rewrite
	aFeatures := a.GetFeaturesOfType(feature.City)

	city := NewCity(a, aFeatures)

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

	pa := elements.ToPlacedTile(tiletemplates.TwoCityEdgesCornerConnectedShield())
	pa.Position = position.New(1, 1)
	pa.GetPlacedFeatureAtSide(side.Top|side.Right, feature.City).Meeple = elements.Meeple{PlayerID: expectedPlayerID, Type: expectedMeepleType}
	a := elements.BinaryTileFromPlacedTile(pa) // todo binarytiles rewrite

	aFeatures := a.GetFeaturesOfType(feature.City)

	city := NewCity(a, aFeatures)

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

	pa := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	pa.Position = position.New(1, 1)
	pa.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple = elements.Meeple{PlayerID: expectedPlayerID, Type: expectedMeepleType}

	a := elements.BinaryTileFromPlacedTile(pa) // todo binarytiles rewrite
	aFeatures := a.GetFeaturesOfType(feature.City)

	city := NewCity(a, aFeatures)

	pb := elements.ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads().Rotate(3))
	pb.Position = position.New(2, 2)
	b := elements.BinaryTileFromPlacedTile(pb) // todo binarytiles rewrite

	bFeatures := b.GetFeaturesOfType(feature.City)
	city.AddTile(b, bFeatures)

	pc := elements.ToPlacedTile(tiletemplates.FourCityEdgesConnectedShield())
	pc.Position = position.New(1, 2)
	c := elements.BinaryTileFromPlacedTile(pc) // todo binarytiles rewrite

	cFeatures := c.GetFeaturesOfType(feature.City)
	city.AddTile(c, cFeatures)

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
