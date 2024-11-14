package elements

import (
	"slices"
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature/modifier"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/tiletemplates"
)

func TestBinaryTileFromPlacedTileCityWithShield(t *testing.T) {
	// tile with city on top and right, with shield in the city and meeple belonging to player 2
	tile := ToPlacedTile(tiletemplates.TwoCityEdgesCornerConnectedRoadTurn())
	tile.GetPlacedFeatureAtSide(side.Top, feature.City).Meeple =
		Meeple{PlayerID: 2, Type: NormalMeeple}
	tile.GetPlacedFeatureAtSide(side.Top, feature.City).ModifierType = modifier.Shield
	tile.Position = position.New(85, 42)

	expected := BinaryTile(0b01010101_00101010_1_10_000000011_00_0011_0000010011_0001001100_1000001110)
	actual := BinaryTileFromPlacedTile(tile)

	if expected != actual {
		t.Fatalf("expected: %064b\ngot: %064b", expected, actual)
	}
}

func TestBinaryTileFromPlacedTileUnconnectedField(t *testing.T) {
	// tile with cities on all sides, the left one having a shield, and a field in the middle.
	//      On the middle field is a meeple belonging to player 1
	tile := ToPlacedTile(tiles.Tile{
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
		Meeple{PlayerID: 1, Type: NormalMeeple}
	tile.Position = position.New(-21, -37)

	expected := BinaryTile(0b11101011_11011011_1_01_100000000_10_1000_0000001111_0000000000_0000000000)
	actual := BinaryTileFromPlacedTile(tile)

	if expected != actual {
		t.Fatalf("expected: %064b\ngot: %064b", expected, actual)
	}
}

func TestBinaryTileFromPlacedTileMonastery(t *testing.T) {
	// monastery with a single road, with a meeple in the monastery belonging to player 2
	tile := ToPlacedTile(tiletemplates.MonasteryWithSingleRoad())
	tile.Monastery().Meeple = Meeple{PlayerID: 2, Type: NormalMeeple}
	tile.Position = position.New(-128, 127)

	expected := BinaryTile(0b10000000_01111111_1_10_100000000_01_0000_0000000000_0000000100_1111111111)
	actual := BinaryTileFromPlacedTile(tile)

	if expected != actual {
		t.Fatalf("expected: %064b\ngot: %064b", expected, actual)
	}
}

func TestBinaryTileFromPlacedTileEmptyTile(t *testing.T) {
	var tile PlacedTile

	expected := BinaryTile(0b00000000_00000000_0_00_000000000_00_0000_0000000000_0000000000_0000000000)
	actual := BinaryTileFromPlacedTile(tile)

	if expected != actual {
		t.Fatalf("expected: %064b\ngot: %064b", expected, actual)
	}
}

func TestPosition(t *testing.T) {
	expectedPos := position.New(-127, 126)

	tile := ToPlacedTile(tiletemplates.MonasteryWithSingleRoad())
	tile.Position = expectedPos

	binaryTile := BinaryTileFromPlacedTile(tile)
	actualPos := binaryTile.Position()

	if expectedPos != actualPos {
		t.Fatalf("expected: %#v\ngot: %#v", expectedPos, actualPos)
	}

	expectedPos = position.New(85, -42)

	tile = ToPlacedTile(tiletemplates.MonasteryWithSingleRoad())
	tile.Position = expectedPos

	binaryTile = BinaryTileFromPlacedTile(tile)
	actualPos = binaryTile.Position()

	if expectedPos != actualPos {
		t.Fatalf("expected: %#v\ngot: %#v", expectedPos, actualPos)
	}
}

func TestHasMonastery(t *testing.T) {
	tile := ToPlacedTile(tiletemplates.MonasteryWithSingleRoad())
	binaryTile := BinaryTileFromPlacedTile(tile)

	if !binaryTile.HasMonastery() {
		t.Fatalf("expected: %#v\ngot: %#v", true, binaryTile.HasMonastery())
	}

	tile = ToPlacedTile(tiletemplates.RoadsTurn())
	binaryTile = BinaryTileFromPlacedTile(tile)

	if binaryTile.HasMonastery() {
		t.Fatalf("expected: %#v\ngot: %#v", false, binaryTile.HasMonastery())
	}
}

func TestGetMeepleIDAtSide(t *testing.T) {
	tile := ToPlacedTile(tiletemplates.RoadsTurn())
	expectedID := ID(1)

	tile.GetPlacedFeatureAtSide(side.BottomLeftEdge, feature.Field).Meeple =
		Meeple{PlayerID: expectedID, Type: NormalMeeple}

	binaryTile := BinaryTileFromPlacedTile(tile)

	actualID := binaryTile.GetMeepleIDAtSide(SideBottomLeftCorner, feature.Field)
	if expectedID != actualID {
		t.Fatalf("expected: %#v\ngot: %#v", expectedID, actualID)
	}

	expectedID = ID(0)
	actualID = binaryTile.GetMeepleIDAtSide(SideTopRightCorner, feature.Field)
	if expectedID != actualID {
		t.Fatalf("%064b\n%016b\n\nexpected: %#v\ngot: %#v", binaryTile, SideTopRightCorner, expectedID, actualID)
	}
	actualID = binaryTile.GetMeepleIDAtSide(SideTopRightCorner|SideTopLeftCorner|SideBottomRightCorner, feature.Field)
	if expectedID != actualID {
		t.Fatalf("expected: %#v\ngot: %#v", expectedID, actualID)
	}
}

func TestGetMeepleIDAtCenter(t *testing.T) {
	// tile with cities on all sides, the left one having a shield, and a field in the middle.
	//      On the middle field is a meeple belonging to player 1
	tile := ToPlacedTile(tiles.Tile{
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

	expectedID := ID(1)

	tile.GetPlacedFeatureAtSide(side.NoSide, feature.Field).Meeple =
		Meeple{PlayerID: expectedID, Type: NormalMeeple}

	binaryTile := BinaryTileFromPlacedTile(tile)

	actualID := binaryTile.GetMeepleIDAtCenter(feature.Field)

	if expectedID != actualID {
		t.Fatalf("expected: %#v\ngot: %#v", expectedID, actualID)
	}

	expectedID = ID(0)
	actualID = binaryTile.GetMeepleIDAtCenter(feature.Monastery)
	if expectedID != actualID {
		t.Fatalf("expected: %#v\ngot: %#v", expectedID, actualID)
	}

	// monastery with a single road, with a meeple in the monastery belonging to player 2
	tile = ToPlacedTile(tiletemplates.MonasteryWithSingleRoad())

	expectedID = ID(2)

	tile.Monastery().Meeple = Meeple{PlayerID: expectedID, Type: NormalMeeple}

	binaryTile = BinaryTileFromPlacedTile(tile)

	actualID = binaryTile.GetMeepleIDAtCenter(feature.Monastery)

	if expectedID != actualID {
		t.Fatalf("expected: %#v\ngot: %#v", expectedID, actualID)
	}

	expectedID = ID(0)
	actualID = binaryTile.GetMeepleIDAtCenter(feature.Field)
	if expectedID != actualID {
		t.Fatalf("expected: %#v\ngot: %#v", expectedID, actualID)
	}
}

func TestGetConnectedSides(t *testing.T) {
	tile := BinaryTileFromTile(tiletemplates.TwoCityEdgesCornerConnectedRoadTurn())
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

func TestGetConnectedSidesWithNoConnectedSides(t *testing.T) {
	tile := BinaryTileFromTile(tiletemplates.SingleCityEdgeCrossRoad())
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

func TestGetConnectedSidesWithNonexistentSides(t *testing.T) {
	tile := BinaryTileFromTile(tiletemplates.TwoCityEdgesCornerConnectedRoadTurn())
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

func TestGetConnectedSidesWithMultipleFeaturesSides(t *testing.T) {
	tile := BinaryTileFromTile(tiletemplates.StraightRoads())
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

func TestGetFeaturesOfType(t *testing.T) {
	tile := BinaryTileFromTile(tiletemplates.SingleCityEdgeCrossRoad())
	expectedFields := []BinaryTileSide{
		SideTopRightCorner | SideTopLeftCorner,
		SideBottomRightCorner,
		SideBottomLeftCorner,
	}
	expectedCities := []BinaryTileSide{
		SideTop,
	}
	expectedRoads := []BinaryTileSide{
		SideRight,
		SideBottom,
		SideLeft,
	}

	actualFields := tile.GetFeaturesOfType(feature.Field)
	actualCities := tile.GetFeaturesOfType(feature.City)
	actualRoads := tile.GetFeaturesOfType(feature.Road)

	if !slices.Equal(actualFields, expectedFields) {
		t.Fatalf("expected %#v, got %#v", expectedFields, actualFields)
	}

	if !slices.Equal(actualCities, expectedCities) {
		t.Fatalf("expected %#v, got %#v", expectedCities, actualCities)
	}

	if !slices.Equal(actualRoads, expectedRoads) {
		t.Fatalf("expected %#v, got %#v", expectedRoads, actualRoads)
	}
}

func TestGetFeatureSides(t *testing.T) {
	tile := BinaryTileFromTile(tiletemplates.SingleCityEdgeCrossRoad())
	expectedFieldSides := SideTopRightCorner | SideTopLeftCorner | SideBottomRightCorner | SideBottomLeftCorner
	expectedCitySides := SideTop
	expectedRoadSides := SideRight | SideBottom | SideLeft

	actualFieldSides := tile.GetFeatureSides(feature.Field)
	actualCitySides := tile.GetFeatureSides(feature.City)
	actualRoadSides := tile.GetFeatureSides(feature.Road)

	if expectedFieldSides != actualFieldSides {
		t.Fatalf("expected %016b, got %016b", expectedFieldSides, actualFieldSides)
	}

	if expectedCitySides != actualCitySides {
		t.Fatalf("expected %016b, got %016b", expectedCitySides, actualCitySides)
	}

	if expectedRoadSides != actualRoadSides {
		t.Fatalf("expected %016b, got %016b", expectedRoadSides, actualRoadSides)
	}
}
