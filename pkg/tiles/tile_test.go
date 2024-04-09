package tiles

import (
	"reflect"
	"testing"

	buildings "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/buildings"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
)

func TestTileRotate(t *testing.T) {
	var tile Tile
	tile.Features = append(tile.Features, feature.Feature{FeatureType: feature.City, Sides: []side.Side{side.Top, side.Left}})
	tile.Features = append(tile.Features, feature.Feature{FeatureType: feature.Road, Sides: []side.Side{side.Bottom, side.Right}})
	tile.Features = append(tile.Features, feature.Feature{FeatureType: feature.Field, Sides: []side.Side{side.BottomRightEdge, side.RightBottomEdge}})
	tile.HasShield = true
	tile.Building = buildings.None

	var rotated = tile.Rotate(1)

	var expected Tile
	expected.Features = append(expected.Features, feature.Feature{FeatureType: feature.City, Sides: []side.Side{side.Right, side.Top}})
	expected.Features = append(expected.Features, feature.Feature{FeatureType: feature.Road, Sides: []side.Side{side.Left, side.Bottom}})
	expected.Features = append(expected.Features, feature.Feature{FeatureType: feature.Field, Sides: []side.Side{side.LeftBottomEdge, side.BottomLeftEdge}})
	expected.HasShield = true
	expected.Building = buildings.None

	if !reflect.DeepEqual(rotated, expected) {
		t.Fatalf("got\n %#v \nshould be \n%#v", rotated, expected)
	}
}

func TestTileFeatureGet(t *testing.T) {
	var tile Tile
	tile.Features = append(tile.Features, feature.Feature{FeatureType: feature.City, Sides: []side.Side{side.Top, side.Left}})
	tile.Features = append(tile.Features, feature.Feature{FeatureType: feature.Road, Sides: []side.Side{side.Bottom, side.Right}})
	tile.Features = append(tile.Features, feature.Feature{FeatureType: feature.Field, Sides: []side.Side{side.BottomRightEdge, side.RightBottomEdge}})
	tile.HasShield = true
	tile.Building = buildings.None

	var expectedCities = []feature.Feature{
		{
			FeatureType: feature.City, Sides: []side.Side{side.Top, side.Left},
		},
	}

	var expectedRoads = []feature.Feature{
		{
			FeatureType: feature.Road, Sides: []side.Side{side.Bottom, side.Right},
		},
	}

	var expectedFields = []feature.Feature{
		{
			FeatureType: feature.Field, Sides: []side.Side{side.BottomRightEdge, side.RightBottomEdge},
		},
	}

	if !reflect.DeepEqual(tile.Cities(), expectedCities) {
		t.Fatalf("got\n %#v \nshould be \n%#v", tile.Cities(), expectedCities)
	}

	if !reflect.DeepEqual(tile.Roads(), expectedRoads) {
		t.Fatalf("got\n %#v \nshould be \n%#v", tile.Roads(), expectedRoads)
	}

	if !reflect.DeepEqual(tile.Fields(), expectedFields) {
		t.Fatalf("got\n %#v \nshould be \n%#v", tile.Fields(), expectedFields)
	}
}
