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

func TestHasMonastery(t *testing.T) {
	tile := elements.ToPlacedTile(tiletemplates.MonasteryWithSingleRoad())
	binaryTile := FromPlacedTile(tile)

	if !binaryTile.HasMonastery() {
		t.Fatalf("expected: %#v\ngot: %#v", true, binaryTile.HasMonastery())
	}

	tile = elements.ToPlacedTile(tiletemplates.RoadsTurn())
	binaryTile = FromPlacedTile(tile)

	if binaryTile.HasMonastery() {
		t.Fatalf("expected: %#v\ngot: %#v", false, binaryTile.HasMonastery())
	}
}

func TestGetMeepleIDAtSideCenter(t *testing.T) {
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

	expectedID := elements.ID(1)

	tile.GetPlacedFeatureAtSide(side.NoSide, feature.Field).Meeple =
		elements.Meeple{PlayerID: expectedID, Type: elements.NormalMeeple}

	binaryTile := FromPlacedTile(tile)

	actualID := binaryTile.GetMeepleIDAtSide(SideCenter, feature.Field)

	if expectedID != actualID {
		t.Fatalf("expected: %#v\ngot: %#v", expectedID, actualID)
	}

	expectedID = elements.ID(0)
	actualID = binaryTile.GetMeepleIDAtSide(SideCenter, feature.Monastery)
	if expectedID != actualID {
		t.Fatalf("expected: %#v\ngot: %#v", expectedID, actualID)
	}

	// monastery with a single road, with a meeple in the monastery belonging to player 2
	tile = elements.ToPlacedTile(tiletemplates.MonasteryWithSingleRoad())

	expectedID = elements.ID(2)

	tile.Monastery().Meeple = elements.Meeple{PlayerID: expectedID, Type: elements.NormalMeeple}

	binaryTile = FromPlacedTile(tile)

	actualID = binaryTile.GetMeepleIDAtSide(SideCenter, feature.Monastery)

	if expectedID != actualID {
		t.Fatalf("expected: %#v\ngot: %#v", expectedID, actualID)
	}

	expectedID = elements.ID(0)
	actualID = binaryTile.GetMeepleIDAtSide(SideCenter, feature.Field)
	if expectedID != actualID {
		t.Fatalf("expected: %#v\ngot: %#v", expectedID, actualID)
	}
}

func TestGetConnectedFeatures(t *testing.T) {
	tile := FromTile(tiletemplates.TwoCityEdgesCornerConnectedRoadTurn())
	sides := []BinaryTileSide{
		SideTop,
		SideRight,

		SideBottom,
		SideLeft,

		SideBottomRightCorner,
		SideTopLeftCorner,
	}
	features := []feature.Type{
		feature.City,
		feature.City,

		feature.Road,
		feature.Road,

		feature.Field,
		feature.Field,
	}
	expectedResults := []BinaryTileSide{
		SideTop | SideRight,
		SideTop | SideRight,

		SideBottom | SideLeft,
		SideBottom | SideLeft,

		SideBottomRightCorner | SideTopLeftCorner,
		SideBottomRightCorner | SideTopLeftCorner,
	}

	for i := range sides {
		actualResult := tile.GetConnectedSides(sides[i], features[i])
		if actualResult != expectedResults[i] {
			t.Fatalf("tile.GetConnectedSides(%#v, %#v) expected: %016b\ngot: %016b", sides[i], features[i], expectedResults[i], actualResult)
		}
	}
}

func TestGetConnectedFeaturesWithNoConnectedFeatures(t *testing.T) {
	tile := FromTile(tiletemplates.SingleCityEdgeCrossRoad())
	sides := []BinaryTileSide{
		SideTop,
		SideRight,
		SideBottom,
		SideLeft,

		SideBottomRightCorner,
		SideBottomLeftCorner,
	}
	features := []feature.Type{
		feature.City,
		feature.Road,
		feature.Road,
		feature.Road,

		feature.Field,
		feature.Field,
	}
	expectedResults := sides // in this test, no sides have any connections, so the expected output should be the same as input

	for i := range sides {
		actualResult := tile.GetConnectedSides(sides[i], features[i])
		if actualResult != expectedResults[i] {
			t.Fatalf("tile.GetConnectedSides(%#v, %#v) expected: %016b\ngot: %016b", sides[i], features[i], expectedResults[i], actualResult)
		}
	}
}

func TestGetConnectedFeaturesWithNonexistentSides(t *testing.T) {
	tile := FromTile(tiletemplates.TwoCityEdgesCornerConnectedRoadTurn())
	sides := []BinaryTileSide{
		SideTop,
		SideRight,
		SideBottom,
		SideLeft,

		SideTopRightCorner,
		SideTopRightCorner,
		SideBottomLeftCorner,
	}
	features := []feature.Type{
		feature.Field,
		feature.Road,
		feature.City,
		feature.Field,

		feature.Field,
		feature.City,
		feature.Road,
	}
	expectedResult := SideNone // none of the features tested has any side at the tested side, so result should always be SideNone

	for i := range sides {
		actualResult := tile.GetConnectedSides(sides[i], features[i])
		if actualResult != expectedResult {
			t.Fatalf("tile.GetConnectedSides(%#v, %#v) expected: %016b\ngot: %016b", sides[i], features[i], expectedResult, actualResult)
		}
	}
}

func TestGetConnectedFeaturesWithMultipleFeaturesSides(t *testing.T) {
	tile := FromTile(tiletemplates.StraightRoads())
	sides := []BinaryTileSide{
		SideTopRightCorner,
		SideTopRightCorner | SideBottomRightCorner,
	}
	features := []feature.Type{
		feature.Field,
		feature.Field,
	}
	expectedResults := []BinaryTileSide{
		SideTopRightCorner | SideTopLeftCorner,
		SideTopRightCorner | SideBottomRightCorner | SideBottomLeftCorner | SideTopLeftCorner,
	}

	for i := range sides {
		actualResult := tile.GetConnectedSides(sides[i], features[i])
		if actualResult != expectedResults[i] {
			t.Fatalf("tile.GetConnectedSides(%#v, %#v) expected: %016b\ngot: %016b", sides[i], features[i], expectedResults[i], actualResult)
		}
	}
}
