package tileSets

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/tileTemplates"
)

func GetStandardTiles() []tiles.Tile {
	var tiles []tiles.Tile
	// Source: https://en.wikipedia.org/w/index.php?title=Carcassonne_(board_game)&oldid=1214139777#Tiles
	// Code below appends the tiles sourced from the "Non-river terrain tiles" table
	// by taking tiles from each cell, column by column.

	// monastery without roads
	for range 4 {
		tiles = append(tiles, tileTemplates.MonasteryWithoutRoads())
	}

	// monastery with single road
	for range 2 {
		tiles = append(tiles, tileTemplates.MonasteryWithSingleRoad())
	}

	// straight roads
	for range 8 {
		tiles = append(tiles, tileTemplates.StraightRoads())
	}

	// roads turns
	for range 9 {
		tiles = append(tiles, tileTemplates.RoadsTurn())
	}

	// T cross
	for range 4 {
		tiles = append(tiles, tileTemplates.TCrossRoad())
	}

	// + cross
	for range 1 {
		tiles = append(tiles, tileTemplates.XCrossRoad())
	}

	// 1 city edge no roads
	for range 5 {
		tiles = append(tiles, tileTemplates.SingleCityEdgeNoRoads())
	}

	// 1 city edge straight road
	for range 4 {
		tiles = append(tiles, tileTemplates.SingleCityEdgeStraightRoads())
	}

	// 1 city edge -| turn
	for range 3 {
		tiles = append(tiles, tileTemplates.SingleCityEdgeLeftRoadTurn())
	}

	// 1 city edge |- turn
	for range 3 {
		tiles = append(tiles, tileTemplates.SingleCityEdgeRightRoadTurn())
	}

	// 1 city edge, road cross
	for range 3 {
		tiles = append(tiles, tileTemplates.SingleCityEdgeCrossRoad())
	}

	// 2 city edges (up and down)
	for range 3 {
		tiles = append(tiles, tileTemplates.TwoCityEdgesUpAndDownNotConnected())
	}

	// 2 city edges (up and right)
	for range 2 {
		tiles = append(tiles, tileTemplates.TwoCityEdgesCornerNotConnected())
	}

	// 2 city edges (up and down but connected)
	for range 1 {
		tiles = append(tiles, tileTemplates.TwoCityEdgesUpAndDownConnected())
	}

	// 2 city edges (up and down but connected but also shields)
	for range 2 {
		tiles = append(tiles, tileTemplates.TwoCityEdgesUpAndDownConnectedShield())
	}

	// 2 city edges (up and right but connected)
	for range 3 {
		tiles = append(tiles, tileTemplates.TwoCityEdgesCornerConnected())
	}

	// 2 city edges (up and right but connected but with shield)
	for range 2 {
		tiles = append(tiles, tileTemplates.TwoCityEdgesCornerConnectedShield())
	}

	// 2 city edges (up and right but connected but road)
	for range 3 {
		tiles = append(tiles, tileTemplates.TwoCityEdgesCornerConnectedRoadTurn())
	}

	// 2 city edges (up and right but connected but road but shield)
	for range 2 {
		tiles = append(tiles, tileTemplates.TwoCityEdgesCornerConnectedRoadTurnShield())
	}

	// 3 city edges ( but connected)
	for range 3 {
		tiles = append(tiles, tileTemplates.ThreeCityEdgesConnected())
	}

	// 3 city edges (but connected but shield)
	for range 1 {
		tiles = append(tiles, tileTemplates.ThreeCityEdgesConnectedShield())
	}

	// 3 city edges (but connected but road)
	for range 1 {
		tiles = append(tiles, tileTemplates.ThreeCityEdgesConnectedRoad())
	}

	// 3 city edges (but connected but road but shield)
	for range 2 {
		tiles = append(tiles, tileTemplates.ThreeCityEdgesConnectedRoadShield())
	}

	// 4 city edges (but shield)
	for range 1 {
		tiles = append(tiles, tileTemplates.FourCityEdgesConnectedShield())
	}
	return tiles
}
