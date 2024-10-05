package testtilesets

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/tiletemplates"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
)

func OrderedMiniTileSet1() tilesets.TileSet {
	// mini simple set containing (12 tiles in total):
	// 1 monastery with road
	// 2 straight roads
	// 1 straight road with city
	// 3 road turns
	// 2 T crossroads
	// 3 city edges up and down not connected

	var tiles = []tiles.Tile{
		tiletemplates.SingleCityEdgeStraightRoads().Rotate(2), // turn 1
		tiletemplates.RoadsTurn(),
		tiletemplates.RoadsTurn().Rotate(1), // turn 3
		tiletemplates.TCrossRoad().Rotate(3),
		tiletemplates.MonasteryWithSingleRoad().Rotate(2), // turn 5
		tiletemplates.TwoCityEdgesUpAndDownNotConnected().Rotate(1),
		tiletemplates.TwoCityEdgesUpAndDownNotConnected().Rotate(1), // turn 7
		tiletemplates.StraightRoads().Rotate(1),
		tiletemplates.TCrossRoad().Rotate(3), // turn 9
		tiletemplates.TwoCityEdgesUpAndDownNotConnected().Rotate(1),
		tiletemplates.RoadsTurn().Rotate(2), // turn 11
		tiletemplates.StraightRoads(),
	}

	return tilesets.TileSet{
		StartingTile: tiletemplates.SingleCityEdgeStraightRoads(),
		Tiles:        tiles,
	}
}

func OrderedMiniTileSet2() tilesets.TileSet {
	// mini simple set containing (12 tiles in total):
	// 1 monastery with road
	// 2 straight roads
	// 1 straight road with city
	// 3 road turns
	// 2 T crossroads
	// 3 city edges up and down not connected

	var tiles = []tiles.Tile{
		tiletemplates.TCrossRoad().Rotate(1), // 1 turn
		tiletemplates.TwoCityEdgesUpAndDownNotConnected(),
		tiletemplates.TwoCityEdgesUpAndDownNotConnected(),
		tiletemplates.RoadsTurn().Rotate(3),
		tiletemplates.RoadsTurn().Rotate(1), // 5 turn
		tiletemplates.StraightRoads(),
		tiletemplates.TwoCityEdgesUpAndDownNotConnected(),
		tiletemplates.RoadsTurn().Rotate(3),
		tiletemplates.SingleCityEdgeStraightRoads().Rotate(2),
		tiletemplates.MonasteryWithSingleRoad().Rotate(1), // 10 turn
		tiletemplates.TCrossRoad().Rotate(3),
		tiletemplates.StraightRoads(),
	}

	return tilesets.TileSet{
		StartingTile: tiletemplates.SingleCityEdgeStraightRoads(),
		Tiles:        tiles,
	}
}

func EveryTileOnceTileSet() tilesets.TileSet {
	var tiles = []tiles.Tile{
		tiletemplates.MonasteryWithoutRoads(), // 1
		tiletemplates.MonasteryWithSingleRoad().Rotate(1),
		tiletemplates.StraightRoads(), // 3
		tiletemplates.RoadsTurn().Rotate(2),
		tiletemplates.TCrossRoad().Rotate(1), // 5
		tiletemplates.XCrossRoad(),
		tiletemplates.SingleCityEdgeNoRoads().Rotate(1), // 7
		tiletemplates.SingleCityEdgeStraightRoads().Rotate(2),
		tiletemplates.SingleCityEdgeLeftRoadTurn().Rotate(3), // 9
		tiletemplates.SingleCityEdgeRightRoadTurn().Rotate(2),
		tiletemplates.SingleCityEdgeCrossRoad(), // B
		tiletemplates.TwoCityEdgesUpAndDownNotConnected(),
		tiletemplates.TwoCityEdgesUpAndDownConnected(), // D
		tiletemplates.TwoCityEdgesUpAndDownConnectedShield(),
		tiletemplates.TwoCityEdgesCornerNotConnected(), // F
		tiletemplates.TwoCityEdgesCornerConnected().Rotate(3),
		tiletemplates.TwoCityEdgesCornerConnectedShield().Rotate(2), // H
		tiletemplates.TwoCityEdgesCornerConnectedRoadTurn().Rotate(1),
		tiletemplates.TwoCityEdgesCornerConnectedRoadTurnShield().Rotate(2), // J
		tiletemplates.ThreeCityEdgesConnected().Rotate(2),
		tiletemplates.ThreeCityEdgesConnectedShield(), // L
		tiletemplates.ThreeCityEdgesConnectedRoad(),
		tiletemplates.ThreeCityEdgesConnectedRoadShield().Rotate(2), // M
		tiletemplates.FourCityEdgesConnectedShield(),
	}

	return tilesets.TileSet{
		StartingTile: tiletemplates.SingleCityEdgeStraightRoads(),
		Tiles:        tiles,
	}
}
