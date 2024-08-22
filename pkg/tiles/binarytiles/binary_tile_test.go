package binarytiles

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature/modifier"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/tiletemplates"
)

func TestFromPlacedTileCityWithShield(t *testing.T) {
	// tile with city on top and right, with shield in the city and meeple belonging to player 2
	tile := elements.ToPlacedTile(tiletemplates.TwoCityEdgesCornerConnectedRoadTurn())
	tile.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple =
		elements.Meeple{PlayerID: 2, Type: elements.NormalMeeple}
	tile.GetPlacedFeatureAtSide(side.Top, feature.City).ModifierType = modifier.Shield
	tile.Position = position.New(85, 42)

	expected := BinaryTile(0b11010101_10101010_1_10_000000011_00_0011_0000010011_0001001100_1000001110)
	actual := FromPlacedTile(tile)

	if expected != actual {
		t.Fatalf("expected: %064b\ngot: %064b", expected, actual)
	}
}

func TestFromPlacedTileUnconnectedField(t *testing.T) {
	// tile with cities on all sides, the left one having a shield, and a field in the middle.
	//      On the middle field is a meeple belonging to player 1
	tile := elements.ToPlacedTile(tiles.Tile{
		Features: []feature.Feature{
			{
				FeatureType: feature.Field,
				Sides:       side.NoSide,
			},
			{
				FeatureType: feature.City,
				Sides:       side.Top,
			},
			{
				FeatureType: feature.City,
				Sides:       side.Right,
			},
			{
				FeatureType: feature.City,
				Sides:       side.Bottom,
			},
			{
				FeatureType:  feature.City,
				ModifierType: modifier.Shield,
				Sides:        side.Left,
			},
		},
	})
	tile.GetPlacedFeatureAtSide(side.NoSide, feature.Field).Meeple =
		elements.Meeple{PlayerID: 1, Type: elements.NormalMeeple}
	tile.Position = position.New(-21, -37)

	expected := BinaryTile(0b01101011_01011011_1_01_100000000_10_1000_0000001111_0000000000_0000000000)
	actual := FromPlacedTile(tile)

	if expected != actual {
		t.Fatalf("expected: %064b\ngot: %064b", expected, actual)
	}
}

func TestFromPlacedTileMonastery(t *testing.T) {
	// monastery with a single road, with a meeple in the monastery belonging to player 2
	tile := elements.ToPlacedTile(tiletemplates.MonasteryWithSingleRoad())
	tile.Monastery().Meeple = elements.Meeple{PlayerID: 2, Type: elements.NormalMeeple}
	tile.Position = position.New(-128, 127)

	expected := BinaryTile(0b00000000_11111111_1_10_100000000_01_0000_0000000000_0000000100_1111111111)
	actual := FromPlacedTile(tile)

	if expected != actual {
		t.Fatalf("expected: %064b\ngot: %064b", expected, actual)
	}
}
