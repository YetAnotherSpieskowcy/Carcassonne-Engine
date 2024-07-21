package elements

import (
	"testing"

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
	move.Monastery().Meeple.MeepleType = NormalMeeple
	move.Monastery().Meeple.PlayerID = 1

	expectedMonastery := tiletemplates.MonasteryWithSingleRoad().Monastery()

	if move.Monastery().Feature != *expectedMonastery {
		t.Fatalf("got\n %#v \nshould be \n%#v", move.Monastery().Feature, *expectedMonastery)
	}
	if move.Monastery().Meeple.MeepleType != NormalMeeple {
		t.Fatalf("got\n %#v \nshould be \n%#v", move.Monastery().Meeple.MeepleType, NormalMeeple)
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
