package city

import (
	"reflect"
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature/modifier"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/tiletemplates"
)

func TestNewAndGetCompleted(t *testing.T) {
	a := tiletemplates.SingleCityEdgeNoRoads()
	cities := a.Cities()
	cityFeatures := []elements.PlacedFeature{}
	for _, c := range cities {
		cityFeatures = append(cityFeatures,
			elements.PlacedFeature{c, elements.Meeple{elements.NoneMeeple, elements.NonePlayer}})
	}
	pos := elements.NewPosition(1, 1)
	city := NewCity(pos, cityFeatures, false)

	completed := city.GetCompleted()
	if completed {
		t.Fatalf("expected %#v, got %#v instead", false, completed)
	}
}

func TestNewAndGetFeaturesFromTile(t *testing.T) {
	a := tiletemplates.SingleCityEdgeNoRoads()
	cities := a.Cities()
	expectedFeatures := []elements.PlacedFeature{}
	for _, c := range cities {
		expectedFeatures = append(expectedFeatures,
			elements.PlacedFeature{c, elements.Meeple{elements.NoneMeeple, elements.NonePlayer}})
	}
	pos := elements.NewPosition(1, 1)
	city := NewCity(pos, expectedFeatures, false)

	features, ok := city.GetFeaturesFromTile(pos)

	if ok == false {
		t.Fatalf("expected %#v, got %#v instead", true, ok)
	}
	if len(features) != len(expectedFeatures) {
		t.Fatalf("expected %#v, got %#v instead", expectedFeatures, features)
	}
	featureEqual := reflect.DeepEqual(expectedFeatures[0], features[0])
	if !featureEqual {
		t.Fatalf("expected %#v, got %#v instead", true, featureEqual)
	}
}

func TestAddTileAndGetFeaturesFromTile(t *testing.T) {
	a := tiletemplates.SingleCityEdgeNoRoads()
	b := tiletemplates.SingleCityEdgeNoRoads()

	aFeatures := []elements.PlacedFeature{}
	for _, c := range a.Cities() {
		aFeatures = append(aFeatures,
			elements.PlacedFeature{c, elements.Meeple{elements.NoneMeeple, elements.NonePlayer}})
	}
	city := NewCity(elements.NewPosition(1, 1), aFeatures, false)

	bRotated := b.Rotate(2)
	bFeatures := []elements.PlacedFeature{}
	for _, c := range bRotated.Cities() {
		bFeatures = append(bFeatures,
			elements.PlacedFeature{c, elements.Meeple{elements.NoneMeeple, elements.NonePlayer}})
	}

	pos := elements.NewPosition(1, 2)
	city.AddTile(pos, bFeatures, false)

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
	a := tiletemplates.SingleCityEdgeNoRoads()
	b := tiletemplates.SingleCityEdgeNoRoads()

	aFeatures := []elements.PlacedFeature{}
	for _, c := range a.Cities() {
		aFeatures = append(aFeatures,
			elements.PlacedFeature{c, elements.Meeple{elements.NoneMeeple, elements.NonePlayer}})
	}
	city := NewCity(elements.NewPosition(1, 1), aFeatures, false)

	bRotated := b.Rotate(2)
	bFeatures := []elements.PlacedFeature{}
	for _, c := range bRotated.Cities() {
		bFeatures = append(bFeatures,
			elements.PlacedFeature{c, elements.Meeple{elements.NoneMeeple, elements.NonePlayer}})
	}

	pos := elements.NewPosition(1, 2)
	city.AddTile(pos, bFeatures, false)

	var expected = true
	var actual = city.GetCompleted()

	if actual != expected {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}
}

func TestCheckCompletedWhenOpen(t *testing.T) {
	a := tiletemplates.SingleCityEdgeNoRoads()
	b := tiletemplates.TwoCityEdgesCornerConnected()

	aFeatures := []elements.PlacedFeature{}
	for _, c := range a.Cities() {
		aFeatures = append(aFeatures,
			elements.PlacedFeature{c, elements.Meeple{elements.NoneMeeple, elements.NonePlayer}})
	}
	city := NewCity(elements.NewPosition(1, 1), aFeatures, false)

	bRotated := b.Rotate(2)
	bFeatures := []elements.PlacedFeature{}
	for _, c := range bRotated.Cities() {
		bFeatures = append(bFeatures,
			elements.PlacedFeature{c, elements.Meeple{elements.NoneMeeple, elements.NonePlayer}})
	}

	pos := elements.NewPosition(1, 2)
	city.AddTile(pos, bFeatures, false)

	var expected = false
	var actual = city.GetCompleted()

	if actual != expected {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}
}

func TestScoreOneTileCity(t *testing.T) {
	var expectedPlayerID elements.ID = 1
	var expectedMeepleType elements.MeepleType = elements.NormalMeeple
	var expectedScore uint32 = 2

	a := tiletemplates.SingleCityEdgeNoRoads()
	aPlaced := elements.ToPlacedTile(a)

	aFeatures := []elements.PlacedFeature{}
	for _, tmp := range aPlaced.Features {
		if tmp.FeatureType == feature.City {
			tmp.PlayerID = expectedPlayerID
			tmp.MeepleType = expectedMeepleType
			aFeatures = append(aFeatures, tmp)
		}
	}
	city := NewCity(elements.NewPosition(1, 1), aFeatures, aPlaced.TileWithMeeple.HasShield)

	scoreReport := city.GetScoreReport()
	meeples, ok := scoreReport.ReturnedMeeples[uint8(expectedPlayerID)]
	if !ok {
		t.Fatalf("expected player id not in the map")
	}

	numMeeples := meeples[expectedMeepleType]
	if numMeeples != 1 {
		t.Fatalf("expected %#v meeple, got %#v meeples instead", 1, numMeeples)
	}

	score := scoreReport.ReceivedPoints[uint8(expectedPlayerID)]
	if score != expectedScore {
		t.Fatalf("expected %#v, got %#v instead", expectedScore, score)
	}
}

func TestScoreOneTileCityWithShield(t *testing.T) {
	var expectedPlayerID elements.ID = 1
	var expectedMeepleType elements.MeepleType = elements.NormalMeeple
	var expectedScore uint32 = 4

	a := tiletemplates.TwoCityEdgesCornerConnectedShield()
	aPlaced := elements.ToPlacedTile(a)

	aFeatures := []elements.PlacedFeature{}
	shield := false
	for _, tmp := range aPlaced.Features {
		if tmp.FeatureType == feature.City {
			tmp.PlayerID = expectedPlayerID
			tmp.MeepleType = expectedMeepleType
			aFeatures = append(aFeatures, tmp)
			if tmp.ModifierType == modifier.Shield {
				shield = true
			}
		}
	}
	city := NewCity(elements.NewPosition(1, 1), aFeatures, shield)

	scoreReport := city.GetScoreReport()
	meeples, ok := scoreReport.ReturnedMeeples[uint8(expectedPlayerID)]
	if !ok {
		t.Fatalf("expected player id not in the map")
	}

	numMeeples := meeples[expectedMeepleType]
	if numMeeples != 1 {
		t.Fatalf("expected %#v meeple, got %#v meeples instead", 1, numMeeples)
	}

	score := scoreReport.ReceivedPoints[uint8(expectedPlayerID)]
	if score != expectedScore {
		t.Fatalf("expected %#v, got %#v instead", expectedScore, score)
	}
}
