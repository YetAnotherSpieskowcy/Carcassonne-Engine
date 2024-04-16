package city

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/tiletemplates"
)

func TestNewAndGetCompleted(t *testing.T) {
	a := tiletemplates.SingleCityEdgeNoRoads()
	pos := elements.NewPosition(1, 1)
	city := New(pos, a.Cities())

	completed := city.GetCompleted()
	if completed {
		t.Fatalf("expected %#v, got %#v instead", false, completed)
	}
}

func TestNewAndGetFeaturesFromTile(t *testing.T) {
	a := tiletemplates.SingleCityEdgeNoRoads()
	expectedFeatures := a.Cities()
	pos := elements.NewPosition(1, 1)
	city := New(pos, expectedFeatures)

	features, ok := city.GetFeaturesFromTile(pos)

	if ok == false {
		t.Fatalf("expected %#v, got %#v instead", true, ok)
	}
	if len(features) != len(expectedFeatures) {
		t.Fatalf("expected %#v, got %#v instead", expectedFeatures, features)
	}
	featureEqual := features[0].Equals(expectedFeatures[0])
	if !featureEqual {
		t.Fatalf("expected %#v, got %#v instead", true, featureEqual)
	}
}

func TestAddTileAndGetFeaturesFromTile(t *testing.T) {
	a := tiletemplates.SingleCityEdgeNoRoads()
	b := tiletemplates.SingleCityEdgeNoRoads()
	b.Rotate(2)
	pos := elements.NewPosition(1, 2)
	city := New(elements.NewPosition(1, 1), a.Cities())
	city.AddTile(pos, b.Cities())

	expectedFeatures := b.Cities()
	features, ok := city.GetFeaturesFromTile(pos)

	if !ok {
		t.Fatalf("expected %#v, got %#v instead", true, ok)
	}
	if len(features) != len(expectedFeatures) {
		t.Fatalf("expected %#v, got %#v instead", expectedFeatures, features)
	}
	featureEqual := features[0].Equals(expectedFeatures[0])
	if !featureEqual {
		t.Fatalf("expected %#v, got %#v instead", true, featureEqual)
	}

}
