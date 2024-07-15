package tiles_test

import (
	"reflect"
	"testing"

	//revive:disable-next-line:dot-imports Dot imports for package under test are fine.
	. "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature/modifier"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/tiletemplates"
)

func TestTileEqualsReturnsFalseWhenOnlyOneHasShield(t *testing.T) {
	a := tiletemplates.TwoCityEdgesUpAndDownConnectedShield()
	b := tiletemplates.TwoCityEdgesUpAndDownConnected()
	if a.Equals(b) {
		t.Fail()
	}
}

func TestTileEqualsReturnsFalseWhenOnlyOneHasBuilding(t *testing.T) {
	a := tiletemplates.MonasteryWithSingleRoad()
	b := tiletemplates.StraightRoads()
	if a.Equals(b) {
		t.Fail()
	}
}

func TestTileEqualsReturnsFalseWhenFeatureCountDiffers(t *testing.T) {
	a := tiletemplates.StraightRoads()
	b := tiletemplates.TCrossRoad()
	if a.Equals(b) {
		t.Fail()
	}
}

func TestTileEqualsReturnsFalseWhenFeatureSidesDiffer(t *testing.T) {
	a := tiletemplates.TwoCityEdgesUpAndDownNotConnected()
	b := tiletemplates.TwoCityEdgesCornerNotConnected()
	if a.Equals(b) {
		t.Fail()
	}
}

func TestTileEqualsReturnsTrueWhenEqual(t *testing.T) {
	a := tiletemplates.MonasteryWithoutRoads()
	b := tiletemplates.MonasteryWithoutRoads()
	if !a.Equals(b) {
		t.Fail()
	}
}

func TestGetTileRotationsStraightRoad(t *testing.T) {
	tile := tiletemplates.StraightRoads()
	actual := len(tile.GetTileRotations())
	expected := 2
	if actual != expected {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}
}

func TestGetTileRotationsStartingTile(t *testing.T) {
	tile := tiletemplates.SingleCityEdgeStraightRoads()
	actual := len(tile.GetTileRotations())
	expected := 4
	if actual != expected {
		t.Fatalf("expected %#v, got %#v instead", expected, actual)
	}
}

func TestTileEqualsReturnsTrueWhenEqualButRotated(t *testing.T) {
	a := tiletemplates.MonasteryWithoutRoads()
	b := tiletemplates.MonasteryWithoutRoads().Rotate(1)
	c := tiletemplates.MonasteryWithoutRoads().Rotate(2)
	d := tiletemplates.MonasteryWithoutRoads().Rotate(3)
	if !a.Equals(b) {
		t.Fatalf("a != b")
	}
	if !a.Equals(c) {
		t.Fatalf("a != c")
	}
	if !a.Equals(d) {
		t.Fatalf("a != d")
	}
}

func TestTileRotate(t *testing.T) {
	var tile Tile
	tile.Features = append(tile.Features, feature.Feature{FeatureType: feature.City, ModifierType: modifier.Shield, Sides: side.Top | side.Left})
	tile.Features = append(tile.Features, feature.Feature{FeatureType: feature.Road, Sides: side.Bottom | side.Right})
	tile.Features = append(tile.Features, feature.Feature{FeatureType: feature.Field, Sides: side.BottomRightEdge | side.RightBottomEdge})

	var rotated = tile.Rotate(1)

	var expected Tile
	expected.Features = append(expected.Features, feature.Feature{FeatureType: feature.City, ModifierType: modifier.Shield, Sides: side.Right | side.Top})
	expected.Features = append(expected.Features, feature.Feature{FeatureType: feature.Road, Sides: side.Left | side.Bottom})
	expected.Features = append(expected.Features, feature.Feature{FeatureType: feature.Field, Sides: side.LeftBottomEdge | side.BottomLeftEdge})

	if !reflect.DeepEqual(rotated, expected) {
		t.Fatalf("got\n %#v \nshould be \n%#v", rotated, expected)
	}
}

func TestTileFeatureGet(t *testing.T) {
	var tile Tile
	tile.Features = append(tile.Features, feature.Feature{FeatureType: feature.City, ModifierType: modifier.Shield, Sides: side.Top | side.Left})
	tile.Features = append(tile.Features, feature.Feature{FeatureType: feature.Road, Sides: side.Bottom | side.Right})
	tile.Features = append(tile.Features, feature.Feature{FeatureType: feature.Field, Sides: side.BottomRightEdge | side.RightBottomEdge})
	tile.Features = append(tile.Features, feature.Feature{FeatureType: feature.Monastery})

	var expectedCities = []feature.Feature{
		{
			FeatureType: feature.City, ModifierType: modifier.Shield, Sides: side.Top | side.Left,
		},
	}

	var expectedRoads = []feature.Feature{
		{
			FeatureType: feature.Road, Sides: side.Bottom | side.Right,
		},
	}

	var expectedFields = []feature.Feature{
		{
			FeatureType: feature.Field, Sides: side.BottomRightEdge | side.RightBottomEdge,
		},
	}

	var expectedMonastery = feature.Feature{FeatureType: feature.Monastery}

	if !reflect.DeepEqual(tile.Cities(), expectedCities) {
		t.Fatalf("got\n %#v \nshould be \n%#v", tile.Cities(), expectedCities)
	}

	if !reflect.DeepEqual(tile.Roads(), expectedRoads) {
		t.Fatalf("got\n %#v \nshould be \n%#v", tile.Roads(), expectedRoads)
	}

	if !reflect.DeepEqual(tile.Fields(), expectedFields) {
		t.Fatalf("got\n %#v \nshould be \n%#v", tile.Fields(), expectedFields)
	}

	if *tile.Monastery() != expectedMonastery {
		t.Fatalf("got\n %#v \nshould be \n%#v", *tile.Monastery(), expectedMonastery)
	}
}
