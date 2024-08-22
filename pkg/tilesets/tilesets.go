package tilesets

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/tiletemplates"
)

type TileSet struct {
	StartingTile tiles.Tile
	Tiles        []tiles.Tile
}

func StandardTileSet() TileSet { //nolint:gocyclo // shallow loops for adding tiles
	var tiles []tiles.Tile
	// Source: https://en.wikipedia.org/w/index.php?title=Carcassonne_(board_game)&oldid=1214139777#Tiles
	// Code below appends the tiles sourced from the "Non-river terrain tiles" table
	// by taking tiles from each cell, column by column.

	// monastery without roads
	for range 4 {
		tiles = append(tiles, tiletemplates.MonasteryWithoutRoads())
	}

	// monastery with single road
	for range 2 {
		tiles = append(tiles, tiletemplates.MonasteryWithSingleRoad())
	}

	// straight roads
	for range 8 {
		tiles = append(tiles, tiletemplates.StraightRoads())
	}

	// roads turns
	for range 9 {
		tiles = append(tiles, tiletemplates.RoadsTurn())
	}

	// T cross
	for range 4 {
		tiles = append(tiles, tiletemplates.TCrossRoad())
	}

	// + cross
	for range 1 {
		tiles = append(tiles, tiletemplates.XCrossRoad())
	}

	// 1 city edge no roads
	for range 5 {
		tiles = append(tiles, tiletemplates.SingleCityEdgeNoRoads())
	}

	// 1 city edge straight road
	for range 3 {
		tiles = append(tiles, tiletemplates.SingleCityEdgeStraightRoads())
	}

	// 1 city edge -| turn
	for range 3 {
		tiles = append(tiles, tiletemplates.SingleCityEdgeLeftRoadTurn())
	}

	// 1 city edge |- turn
	for range 3 {
		tiles = append(tiles, tiletemplates.SingleCityEdgeRightRoadTurn())
	}

	// 1 city edge, road cross
	for range 3 {
		tiles = append(tiles, tiletemplates.SingleCityEdgeCrossRoad())
	}

	// 2 city edges (up and down)
	for range 3 {
		tiles = append(tiles, tiletemplates.TwoCityEdgesUpAndDownNotConnected())
	}

	// 2 city edges (up and right)
	for range 2 {
		tiles = append(tiles, tiletemplates.TwoCityEdgesCornerNotConnected())
	}

	// 2 city edges (up and down but connected)
	for range 1 {
		tiles = append(tiles, tiletemplates.TwoCityEdgesUpAndDownConnected())
	}

	// 2 city edges (up and down but connected but also shields)
	for range 2 {
		tiles = append(tiles, tiletemplates.TwoCityEdgesUpAndDownConnectedShield())
	}

	// 2 city edges (up and right but connected)
	for range 3 {
		tiles = append(tiles, tiletemplates.TwoCityEdgesCornerConnected())
	}

	// 2 city edges (up and right but connected but with shield)
	for range 2 {
		tiles = append(tiles, tiletemplates.TwoCityEdgesCornerConnectedShield())
	}

	// 2 city edges (up and right but connected but road)
	for range 3 {
		tiles = append(tiles, tiletemplates.TwoCityEdgesCornerConnectedRoadTurn())
	}

	// 2 city edges (up and right but connected but road but shield)
	for range 2 {
		tiles = append(tiles, tiletemplates.TwoCityEdgesCornerConnectedRoadTurnShield())
	}

	// 3 city edges ( but connected)
	for range 3 {
		tiles = append(tiles, tiletemplates.ThreeCityEdgesConnected())
	}

	// 3 city edges (but connected but shield)
	for range 1 {
		tiles = append(tiles, tiletemplates.ThreeCityEdgesConnectedShield())
	}

	// 3 city edges (but connected but road)
	for range 1 {
		tiles = append(tiles, tiletemplates.ThreeCityEdgesConnectedRoad())
	}

	// 3 city edges (but connected but road but shield)
	for range 2 {
		tiles = append(tiles, tiletemplates.ThreeCityEdgesConnectedRoadShield())
	}

	// 4 city edges (but shield)
	for range 1 {
		tiles = append(tiles, tiletemplates.FourCityEdgesConnectedShield())
	}

	return TileSet{
		StartingTile: tiletemplates.SingleCityEdgeStraightRoads(),
		Tiles:        tiles,
	}
}

func MiniTileSet() TileSet {
	var tiles = []tiles.Tile{}
	// mini simple set containing (12 tiles in total):
	// 1 monastery with road
	// 2 straight roads
	// 1 straight road with city
	// 3 road turns
	// 2 T crossroads
	// 3 city edges up and down not connected

	// 1 monastery with road
	tiles = append(tiles, tiletemplates.MonasteryWithSingleRoad())

	// 2 straight roads
	tiles = append(tiles, tiletemplates.StraightRoads())
	tiles = append(tiles, tiletemplates.StraightRoads())

	// 1 straight road with city
	tiles = append(tiles, tiletemplates.SingleCityEdgeStraightRoads())

	// 3 road turns
	tiles = append(tiles, tiletemplates.RoadsTurn())
	tiles = append(tiles, tiletemplates.RoadsTurn())
	tiles = append(tiles, tiletemplates.RoadsTurn())

	// 2 T crossroads
	tiles = append(tiles, tiletemplates.TCrossRoad())
	tiles = append(tiles, tiletemplates.TCrossRoad())

	// 3 city edges up and down not connected
	tiles = append(tiles, tiletemplates.TwoCityEdgesUpAndDownNotConnected())
	tiles = append(tiles, tiletemplates.TwoCityEdgesUpAndDownNotConnected())
	tiles = append(tiles, tiletemplates.TwoCityEdgesUpAndDownNotConnected())
	return TileSet{
		StartingTile: tiletemplates.SingleCityEdgeStraightRoads(),
		Tiles:        tiles,
	}
}

func EveryTileOnceTileSet() TileSet { //nolint:gocyclo // shallow loops for adding tiles
	var tiles = []tiles.Tile{}

	tiles = append(tiles, tiletemplates.MonasteryWithoutRoads())
	tiles = append(tiles, tiletemplates.MonasteryWithSingleRoad())
	tiles = append(tiles, tiletemplates.StraightRoads())
	tiles = append(tiles, tiletemplates.RoadsTurn())
	tiles = append(tiles, tiletemplates.TCrossRoad())
	tiles = append(tiles, tiletemplates.XCrossRoad())
	tiles = append(tiles, tiletemplates.SingleCityEdgeNoRoads())
	tiles = append(tiles, tiletemplates.SingleCityEdgeStraightRoads())
	tiles = append(tiles, tiletemplates.SingleCityEdgeLeftRoadTurn())
	tiles = append(tiles, tiletemplates.SingleCityEdgeRightRoadTurn())
	tiles = append(tiles, tiletemplates.SingleCityEdgeCrossRoad())
	tiles = append(tiles, tiletemplates.TwoCityEdgesUpAndDownNotConnected())
	tiles = append(tiles, tiletemplates.TwoCityEdgesCornerNotConnected())
	tiles = append(tiles, tiletemplates.TwoCityEdgesUpAndDownConnected())
	tiles = append(tiles, tiletemplates.TwoCityEdgesUpAndDownConnectedShield())
	tiles = append(tiles, tiletemplates.TwoCityEdgesCornerConnected())
	tiles = append(tiles, tiletemplates.TwoCityEdgesCornerConnectedShield())
	tiles = append(tiles, tiletemplates.TwoCityEdgesCornerConnectedRoadTurn())
	tiles = append(tiles, tiletemplates.TwoCityEdgesCornerConnectedRoadTurnShield())
	tiles = append(tiles, tiletemplates.ThreeCityEdgesConnected())
	tiles = append(tiles, tiletemplates.ThreeCityEdgesConnectedShield())
	tiles = append(tiles, tiletemplates.ThreeCityEdgesConnectedRoad())
	tiles = append(tiles, tiletemplates.ThreeCityEdgesConnectedRoadShield())
	tiles = append(tiles, tiletemplates.FourCityEdgesConnectedShield())

	return TileSet{
		StartingTile: tiletemplates.SingleCityEdgeStraightRoads(),
		Tiles:        tiles,
	}
}
