package elements

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/tiletemplates"
)

func TestTilePlacementRotate(t *testing.T) {
	move := ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected TilePlacement.Rotate() to panic")
		}
	}()

	move.Rotate(1)
}

func TestPlacedTileFeatureGet(t *testing.T) {
	move := ToPlacedTile(tiletemplates.MonasteryWithSingleRoad())
	move.Monastery().Meeple.Type = NormalMeeple
	move.Monastery().Meeple.PlayerID = 1

	expectedMonastery := tiletemplates.MonasteryWithSingleRoad().Monastery()

	if move.Monastery().Feature != *expectedMonastery {
		t.Fatalf("got\n %#v \nshould be \n%#v", move.Monastery().Feature, *expectedMonastery)
	}
	if move.Monastery().Meeple.Type != NormalMeeple {
		t.Fatalf("got\n %#v \nshould be \n%#v", move.Monastery().Meeple.Type, NormalMeeple)
	}
	if MeepleType(move.Monastery().Meeple.PlayerID) != 1 {
		t.Fatalf("got\n %#v \nshould be \n%#v", move.Monastery().Meeple.PlayerID, 1)
	}
}

func TestGetCityFeatures(t *testing.T) {
	var expectedLen = 1
	var expectedSide side.Side = side.Top

	tile := ToPlacedTile(tiletemplates.SingleCityEdgeNoRoads())

	cityFeatures := tile.GetFeaturesOfType(feature.City)

	if len(cityFeatures) != expectedLen {
		t.Fatalf("expected %#v, got %#v instead", expectedLen, len(cityFeatures))
	}

	actualSide := cityFeatures[0].Sides
	if actualSide != expectedSide {
		t.Fatalf("expected side %#v, got %#v instead", expectedSide, actualSide)
	}
}

func TestGetPlacedFeatureAtSide(t *testing.T) {
	tile := ToPlacedTile(tiletemplates.SingleCityEdgeStraightRoads())

	// city on top of the tile
	tileFeature := tile.GetPlacedFeatureAtSide(side.Top, feature.City)
	if tileFeature == nil || tileFeature.FeatureType != feature.City {
		t.Fatalf("expected a city feature, got %#v instead", tileFeature)
	}

	tileFeature = tile.GetPlacedFeatureAtSide(side.Right, feature.City)
	if tileFeature != nil {
		t.Fatalf("expected nil, got %#v instead", tileFeature)
	}

	tileFeature = tile.GetPlacedFeatureAtSide(side.Bottom, feature.City)
	if tileFeature != nil {
		t.Fatalf("expected nil, got %#v instead", tileFeature)
	}

	tileFeature = tile.GetPlacedFeatureAtSide(side.Left, feature.City)
	if tileFeature != nil {
		t.Fatalf("expected nil, got %#v instead", tileFeature)
	}

	// road on the bottom of the tile
	tileFeature = tile.GetPlacedFeatureAtSide(side.Left, feature.Road)
	if tileFeature == nil || tileFeature.FeatureType != feature.Road {
		t.Fatalf("expected a road feature, got %#v instead", tileFeature)
	}

	tileFeature = tile.GetPlacedFeatureAtSide(side.Right, feature.Road)
	if tileFeature == nil || tileFeature.FeatureType != feature.Road {
		t.Fatalf("expected a road feature, got %#v instead", tileFeature)
	}

	tileFeature = tile.GetPlacedFeatureAtSide(side.Top, feature.Road)
	if tileFeature != nil {
		t.Fatalf("expected nil, got %#v instead", tileFeature)
	}

	tileFeature = tile.GetPlacedFeatureAtSide(side.Bottom, feature.Road)
	if tileFeature != nil {
		t.Fatalf("expected nil, got %#v instead", tileFeature)
	}
}

func TestGetPlacedFeaturesOverlappingSides(t *testing.T) {
	tile := ToPlacedTile(tiletemplates.SingleCityEdgeCrossRoad())

	// city on top of the tile
	tileFeatures := tile.GetPlacedFeaturesOverlappingSide(side.Top, feature.City)
	if len(tileFeatures) != 1 {
		t.Fatalf("expected 1 feature, got %#v features instead", len(tileFeatures))
	}
	for _, f := range tileFeatures {
		if f.FeatureType != feature.City {
			t.Fatalf("expected a city feature, got %#v instead", f.FeatureType)
		}
	}

	// roads on right, bottom and left of the tile
	tileFeatures = tile.GetPlacedFeaturesOverlappingSide(side.Left|side.RightTopEdge|side.BottomLeftEdge, feature.Road)
	if len(tileFeatures) != 3 {
		t.Fatalf("expected 3 features, got %#v features instead", len(tileFeatures))
	}
	for _, f := range tileFeatures {
		if f.FeatureType != feature.Road {
			t.Fatalf("expected a road feature, got %#v instead", f.FeatureType)
		}
	}

	// fields
	tileFeatures = tile.GetPlacedFeaturesOverlappingSide(side.All, feature.Field)
	if len(tileFeatures) != 3 {
		t.Fatalf("expected 3 features, got %#v features instead", len(tileFeatures))
	}
	for _, f := range tileFeatures {
		if f.FeatureType != feature.Field {
			t.Fatalf("expected a road feature, got %#v instead", f.FeatureType)
		}
	}

	// fields on none side (should be 0)
	tileFeatures = tile.GetPlacedFeaturesOverlappingSide(side.NoSide, feature.Field)
	if len(tileFeatures) != 0 {
		t.Fatalf("expected 0 features, got %#v features instead", len(tileFeatures))
	}
	for _, f := range tileFeatures {
		t.Fatalf("expected nothing, got %#v instead", f)
	}
}

func TestEqualsTile(t *testing.T) {
	// equal tiles
	tile := tiletemplates.MonasteryWithSingleRoad()
	placedTile := ToPlacedTile(tile)
	placedTile.Position = position.New(12, 34)
	placedTile.GetPlacedFeatureAtSide(side.Bottom, feature.Road).Meeple = Meeple{NormalMeeple, ID(5)}
	if !placedTile.EqualsTile(tile) {
		t.Fatalf("expected %#v, got %#v instead", true, false)
	}

	// equal tiles, but rotated
	tile = tiletemplates.SingleCityEdgeCrossRoad()
	placedTile = ToPlacedTile(tile.Rotate(1))
	placedTile.Position = position.New(-43, -21)
	placedTile.GetPlacedFeatureAtSide(side.Right, feature.City).Meeple = Meeple{NormalMeeple, ID(3)}
	if !placedTile.EqualsTile(tile) {
		t.Fatalf("expected %#v, got %#v instead", true, false)
	}

	// non-equal tiles with same number of features
	tile = tiletemplates.StraightRoads()
	placedTile = ToPlacedTile(tiletemplates.RoadsTurn())
	placedTile.Position = position.New(-43, -21)
	placedTile.GetPlacedFeatureAtSide(side.Left, feature.Road).Meeple = Meeple{NormalMeeple, ID(3)}
	if placedTile.EqualsTile(tile) {
		t.Fatalf("expected %#v, got %#v instead", false, true)
	}

	// non-equal tiles with different number of features
	tile = tiletemplates.TestOnlyField()
	placedTile = ToPlacedTile(tiletemplates.MonasteryWithoutRoads())
	placedTile.Position = position.New(1, 0)
	if placedTile.EqualsTile(tile) {
		t.Fatalf("expected %#v, got %#v instead", false, true)
	}
}

func TestExactEqualsTile(t *testing.T) {
	// equal tiles
	tile := tiletemplates.MonasteryWithSingleRoad()
	placedTile := ToPlacedTile(tile)
	placedTile.Position = position.New(12, 34)
	placedTile.GetPlacedFeatureAtSide(side.Bottom, feature.Road).Meeple = Meeple{NormalMeeple, ID(5)}
	if !placedTile.ExactEqualsTile(tile) {
		t.Fatalf("expected %#v, got %#v instead", true, false)
	}

	// equal tiles, but rotated
	tile = tiletemplates.SingleCityEdgeCrossRoad()
	placedTile = ToPlacedTile(tile.Rotate(1))
	placedTile.Position = position.New(-43, -21)
	placedTile.GetPlacedFeatureAtSide(side.Right, feature.City).Meeple = Meeple{NormalMeeple, ID(3)}
	if placedTile.ExactEqualsTile(tile) {
		t.Fatalf("expected %#v, got %#v instead", false, true)
	}

	// non-equal tiles with same number of features
	tile = tiletemplates.StraightRoads()
	placedTile = ToPlacedTile(tiletemplates.RoadsTurn())
	placedTile.Position = position.New(-43, -21)
	placedTile.GetPlacedFeatureAtSide(side.Left, feature.Road).Meeple = Meeple{NormalMeeple, ID(3)}
	if placedTile.ExactEqualsTile(tile) {
		t.Fatalf("expected %#v, got %#v instead", false, true)
	}

	// non-equal tiles with different number of features
	tile = tiletemplates.TestOnlyField()
	placedTile = ToPlacedTile(tiletemplates.MonasteryWithoutRoads())
	placedTile.Position = position.New(1, 0)
	if placedTile.ExactEqualsTile(tile) {
		t.Fatalf("expected %#v, got %#v instead", false, true)
	}
}
