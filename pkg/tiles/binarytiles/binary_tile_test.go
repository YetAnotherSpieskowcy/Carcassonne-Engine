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

	tile := elements.ToPlacedTile(tiles.Tile{
		Features: []feature.Feature{
			feature.New(
				feature.City,
				side.Top|
					side.Right,
				modifier.Shield,
			),
			feature.New(
				feature.Road,
				side.Left|
					side.Bottom,
			),
			feature.New(
				feature.Field,
				side.LeftBottomEdge|
					side.BottomLeftEdge,
			),
			feature.New(
				feature.Field,
				side.LeftTopEdge|
					side.BottomRightEdge,
			),
		},
	})

	tile.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple =
		elements.Meeple{PlayerID: 2, Type: elements.NormalMeeple}
	tile.Position = position.New(85, 42)

	expected := BinaryTile(0b01010101_00101010_1_10_000000011_00_0011_0000010011_0001001100_1000001110)
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
			feature.New(
				feature.Field,
				side.NoSide,
			),
			feature.New(
				feature.City,
				side.Top,
			),
			feature.New(
				feature.City,
				side.Right,
			),
			feature.New(
				feature.City,
				side.Bottom,
			),
			feature.New(
				feature.City,

				side.Left,
				modifier.Shield,
			),
		},
	})
	tile.GetPlacedFeatureAtSide(side.NoSide, feature.Field).Meeple =
		elements.Meeple{PlayerID: 1, Type: elements.NormalMeeple}
	tile.Position = position.New(-21, -37)

	expected := BinaryTile(0b11101011_11011011_1_01_100000000_10_1000_0000001111_0000000000_0000000000)
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

	expected := BinaryTile(0b10000000_01111111_1_10_100000000_01_0000_0000000000_0000000100_1111111111)
	actual := FromPlacedTile(tile)

	if expected != actual {
		t.Fatalf("expected: %064b\ngot: %064b", expected, actual)
	}
}

func TestFromPlacedTileEmptyTile(t *testing.T) {
	var tile elements.PlacedTile

	expected := BinaryTile(0b00000000_00000000_0_00_000000000_00_0000_0000000000_0000000000_0000000000)
	actual := FromPlacedTile(tile)

	if expected != actual {
		t.Fatalf("expected: %064b\ngot: %064b", expected, actual)
	}
}

func TestPosition(t *testing.T) {
	expectedPos := position.New(-127, 126)

	tile := elements.ToPlacedTile(tiletemplates.MonasteryWithSingleRoad())
	tile.Position = expectedPos

	binaryTile := FromPlacedTile(tile)
	actualPos := binaryTile.Position()

	if expectedPos != actualPos {
		t.Fatalf("expected: %#v\ngot: %#v", expectedPos, actualPos)
	}

	expectedPos = position.New(85, -42)

	tile = elements.ToPlacedTile(tiletemplates.MonasteryWithSingleRoad())
	tile.Position = expectedPos

	binaryTile = FromPlacedTile(tile)
	actualPos = binaryTile.Position()

	if expectedPos != actualPos {
		t.Fatalf("expected: %#v\ngot: %#v", expectedPos, actualPos)
	}
}
