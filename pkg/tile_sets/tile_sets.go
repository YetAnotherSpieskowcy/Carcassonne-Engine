package tile_sets

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/tile_templates"
)

func GetStandardTiles() []tiles.Tile {
	var tiles []tiles.Tile
	// Source: https://en.wikipedia.org/w/index.php?title=Carcassonne_(board_game)&oldid=1214139777#Tiles
	// Code below appends the tiles sourced from the "Non-river terrain tiles" table
	// by taking tiles from each cell, column by column.

	//monastery without roads
	for i := 0; i < 5; i++ {
		tiles = append(tiles, tile_templates.MonasteryWithoutRoads())
	}

	//monastery with single road
	for i := 0; i < 2; i++ {
		tiles = append(tiles, tile_templates.MonasteryWithSingleRoad())
	}

	//straight roads
	for i := 0; i < 8; i++ {
		tiles = append(tiles, tile_templates.StraightRoads())
	}

	//roads turns
	for i := 0; i < 9; i++ {
		tiles = append(tiles, tile_templates.RoadsTurn())
	}

	//T cross
	for i := 0; i < 9; i++ {
		tiles = append(tiles, tile_templates.TCrossRoad())
	}

	// + cross
	for i := 0; i < 1; i++ {
		tiles = append(tiles, tile_templates.XCrossRoad())
	}

	//1 city edge no roads
	for i := 0; i < 5; i++ {
		tiles = append(tiles, tile_templates.SingleCityEdgeNoRoads())
	}

	//1 city edge straight road
	for i := 0; i < 4; i++ {
		tiles = append(tiles, tile_templates.SingleCityEdgeStraightRoads())
	}

	//1 city edge -| turn
	for i := 0; i < 4; i++ {
		tiles = append(tiles, tile_templates.SingleCityEdgeLeftRoadTurn())
	}

	//1 city edge |- turn
	for i := 0; i < 4; i++ {
		tiles = append(tiles, tile_templates.SingleCityEdgeRightRoadTurn())
	}

	//1 city edge, road cross
	for i := 0; i < 4; i++ {
		tiles = append(tiles, tile_templates.SingleCityEdgeCrossRoad())
	}

	//2 city edges (up and down)
	for i := 0; i < 3; i++ {
		tiles = append(tiles, tile_templates.TwoCityEdgesUpAndDownNotConnected())
	}

	//2 city edges (up and right)
	for i := 0; i < 2; i++ {
		tiles = append(tiles, tile_templates.TwoCityEdgesCornerNotConnected())
	}

	//2 city edges (up and down but connected)
	for i := 0; i < 1; i++ {
		tiles = append(tiles, tile_templates.TwoCityEdgesUpAndDownConnected())
	}

	//2 city edges (up and down but connected but also shields)
	for i := 0; i < 2; i++ {
		tiles = append(tiles, tile_templates.TwoCityEdgesUpAndDownConnectedShield())
	}

	//2 city edges (up and right but connected)
	for i := 0; i < 3; i++ {
		tiles = append(tiles, tile_templates.TwoCityEdgesCornerConnected())
	}

	//2 city edges (up and right but connected but with shield)
	for i := 0; i < 2; i++ {
		tiles = append(tiles, tile_templates.TwoCityEdgesCornerConnectedShield())
	}

	//2 city edges (up and right but connected but road)
	for i := 0; i < 3; i++ {
		tiles = append(tiles, tile_templates.TwoCityEdgesCornerConnectedRoadTurn())
	}

	//2 city edges (up and right but connected but road but shield)
	for i := 0; i < 2; i++ {
		tiles = append(tiles, tile_templates.TwoCityEdgesCornerConnectedRoadTurnShield())
	}

	//3 city edges ( but connected)
	for i := 0; i < 3; i++ {
		tiles = append(tiles, tile_templates.ThreeCityEdgesConnected())
	}

	//3 city edges (but connected but shield)
	for i := 0; i < 1; i++ {
		tiles = append(tiles, tile_templates.ThreeCityEdgesConnectedShield())
	}

	//3 city edges (but connected but road)
	for i := 0; i < 1; i++ {
		tiles = append(tiles, tile_templates.ThreeCityEdgesConnectedRoad())
	}

	//3 city edges (but connected but road but shield)
	for i := 0; i < 2; i++ {
		tiles = append(tiles, tile_templates.ThreeCityEdgesConnectedRoadShield())
	}

	//4 city edges (but shield)
	for i := 0; i < 1; i++ {
		tiles = append(tiles, tile_templates.FourCityEdgesConnectedShield())
	}
	return tiles
}
