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
	for range 4 {
		tiles = append(tiles, tile_templates.MonasteryWithoutRoads())
	}

	//monastery with single road
	for range 2 {
		tiles = append(tiles, tile_templates.MonasteryWithSingleRoad())
	}

	//straight roads
	for range 8 {
		tiles = append(tiles, tile_templates.StraightRoads())
	}

	//roads turns
	for range 9 {
		tiles = append(tiles, tile_templates.RoadsTurn())
	}

	//T cross
	for range 4 {
		tiles = append(tiles, tile_templates.TCrossRoad())
	}

	// + cross
	for range 1 {
		tiles = append(tiles, tile_templates.XCrossRoad())
	}

	//1 city edge no roads
	for range 5 {
		tiles = append(tiles, tile_templates.SingleCityEdgeNoRoads())
	}

	//1 city edge straight road
	for range 4 {
		tiles = append(tiles, tile_templates.SingleCityEdgeStraightRoads())
	}

	//1 city edge -| turn
	for range 3 {
		tiles = append(tiles, tile_templates.SingleCityEdgeLeftRoadTurn())
	}

	//1 city edge |- turn
	for range 3 {
		tiles = append(tiles, tile_templates.SingleCityEdgeRightRoadTurn())
	}

	//1 city edge, road cross
	for range 3 {
		tiles = append(tiles, tile_templates.SingleCityEdgeCrossRoad())
	}

	//2 city edges (up and down)
	for range 3 {
		tiles = append(tiles, tile_templates.TwoCityEdgesUpAndDownNotConnected())
	}

	//2 city edges (up and right)
	for range 2 {
		tiles = append(tiles, tile_templates.TwoCityEdgesCornerNotConnected())
	}

	//2 city edges (up and down but connected)
	for range 1 {
		tiles = append(tiles, tile_templates.TwoCityEdgesUpAndDownConnected())
	}

	//2 city edges (up and down but connected but also shields)
	for range 2 {
		tiles = append(tiles, tile_templates.TwoCityEdgesUpAndDownConnectedShield())
	}

	//2 city edges (up and right but connected)
	for range 3 {
		tiles = append(tiles, tile_templates.TwoCityEdgesCornerConnected())
	}

	//2 city edges (up and right but connected but with shield)
	for range 2 {
		tiles = append(tiles, tile_templates.TwoCityEdgesCornerConnectedShield())
	}

	//2 city edges (up and right but connected but road)
	for range 3 {
		tiles = append(tiles, tile_templates.TwoCityEdgesCornerConnectedRoadTurn())
	}

	//2 city edges (up and right but connected but road but shield)
	for range 2 {
		tiles = append(tiles, tile_templates.TwoCityEdgesCornerConnectedRoadTurnShield())
	}

	//3 city edges ( but connected)
	for range 3 {
		tiles = append(tiles, tile_templates.ThreeCityEdgesConnected())
	}

	//3 city edges (but connected but shield)
	for range 1 {
		tiles = append(tiles, tile_templates.ThreeCityEdgesConnectedShield())
	}

	//3 city edges (but connected but road)
	for range 1 {
		tiles = append(tiles, tile_templates.ThreeCityEdgesConnectedRoad())
	}

	//3 city edges (but connected but road but shield)
	for range 2 {
		tiles = append(tiles, tile_templates.ThreeCityEdgesConnectedRoadShield())
	}

	//4 city edges (but shield)
	for range 1 {
		tiles = append(tiles, tile_templates.FourCityEdgesConnectedShield())
	}
	return tiles
}
