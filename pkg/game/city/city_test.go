package city

import (
	"reflect"
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/tiletemplates"
)

func TestNewAndGetCompleted(t *testing.T) {
	a := tiletemplates.SingleCityEdgeNoRoads()
	cities := a.Cities()
	cityFeatures := []elements.PlacedFeature{}
	for _, c := range cities {
		cityFeatures = append(cityFeatures,
			elements.PlacedFeature{c, elements.NoneMeeple, elements.NonePlayer})
	}
	pos := elements.NewPosition(1, 1)
	city := New(pos, cityFeatures)

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
			elements.PlacedFeature{c, elements.NoneMeeple, elements.NonePlayer})
	}
	pos := elements.NewPosition(1, 1)
	city := New(pos, expectedFeatures)

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
			elements.PlacedFeature{c, elements.NoneMeeple, elements.NonePlayer})
	}
	city := New(elements.NewPosition(1, 1), aFeatures)

	bRotated := b.Rotate(2)
	bFeatures := []elements.PlacedFeature{}
	for _, c := range bRotated.Cities() {
		bFeatures = append(bFeatures,
			elements.PlacedFeature{c, elements.NoneMeeple, elements.NonePlayer})
	}

	pos := elements.NewPosition(1, 2)
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
	a := tiletemplates.SingleCityEdgeNoRoads()
	b := tiletemplates.SingleCityEdgeNoRoads()

	aFeatures := []elements.PlacedFeature{}
	for _, c := range a.Cities() {
		aFeatures = append(aFeatures,
			elements.PlacedFeature{c, elements.NoneMeeple, elements.NonePlayer})
	}
	city := New(elements.NewPosition(1, 1), aFeatures)

	bRotated := b.Rotate(2)
	bFeatures := []elements.PlacedFeature{}
	for _, c := range bRotated.Cities() {
		bFeatures = append(bFeatures,
			elements.PlacedFeature{c, elements.NoneMeeple, elements.NonePlayer})
	}

	pos := elements.NewPosition(1, 2)
	city.AddTile(pos, bFeatures)

	var expected bool = true
	var actual bool = city.GetCompleted()

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
			elements.PlacedFeature{c, elements.NoneMeeple, elements.NonePlayer})
	}
	city := New(elements.NewPosition(1, 1), aFeatures)

	bRotated := b.Rotate(2)
	bFeatures := []elements.PlacedFeature{}
	for _, c := range bRotated.Cities() {
		bFeatures = append(bFeatures,
			elements.PlacedFeature{c, elements.NoneMeeple, elements.NonePlayer})
	}

	pos := elements.NewPosition(1, 2)
	city.AddTile(pos, bFeatures)

	var expected bool = false
	var actual bool = city.GetCompleted()

	if actual != expected {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}
}
