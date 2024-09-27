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

func EveryTileOnceTileSet() TileSet { //nolint:gocyclo // shallow loops for adding tiles
	var tiles = []tiles.Tile{}

	tiles = append(tiles, tiletemplates.MonasteryWithoutRoads()) // 1
	tiles = append(tiles, tiletemplates.MonasteryWithSingleRoad().Rotate(1))
	tiles = append(tiles, tiletemplates.StraightRoads()) // 3
	tiles = append(tiles, tiletemplates.RoadsTurn().Rotate(2))
	tiles = append(tiles, tiletemplates.TCrossRoad().Rotate(1)) // 5
	tiles = append(tiles, tiletemplates.XCrossRoad())
	tiles = append(tiles, tiletemplates.SingleCityEdgeNoRoads().Rotate(1)) // 7
	tiles = append(tiles, tiletemplates.SingleCityEdgeStraightRoads().Rotate(2))
	tiles = append(tiles, tiletemplates.SingleCityEdgeLeftRoadTurn().Rotate(3)) // 9
	tiles = append(tiles, tiletemplates.SingleCityEdgeRightRoadTurn().Rotate(2))
	tiles = append(tiles, tiletemplates.SingleCityEdgeCrossRoad()) // B
	tiles = append(tiles, tiletemplates.TwoCityEdgesUpAndDownNotConnected())
	tiles = append(tiles, tiletemplates.TwoCityEdgesUpAndDownConnected()) // D
	tiles = append(tiles, tiletemplates.TwoCityEdgesUpAndDownConnectedShield())
	tiles = append(tiles, tiletemplates.TwoCityEdgesCornerNotConnected()) // F
	tiles = append(tiles, tiletemplates.TwoCityEdgesCornerConnected().Rotate(3))
	tiles = append(tiles, tiletemplates.TwoCityEdgesCornerConnectedShield().Rotate(2)) // H
	tiles = append(tiles, tiletemplates.TwoCityEdgesCornerConnectedRoadTurn().Rotate(1))
	tiles = append(tiles, tiletemplates.TwoCityEdgesCornerConnectedRoadTurnShield().Rotate(2)) // J
	tiles = append(tiles, tiletemplates.ThreeCityEdgesConnected().Rotate(2))
	tiles = append(tiles, tiletemplates.ThreeCityEdgesConnectedShield()) // L
	tiles = append(tiles, tiletemplates.ThreeCityEdgesConnectedRoad())
	tiles = append(tiles, tiletemplates.ThreeCityEdgesConnectedRoadShield().Rotate(2)) // N
	tiles = append(tiles, tiletemplates.FourCityEdgesConnectedShield())

	return TileSet{
		StartingTile: tiletemplates.SingleCityEdgeStraightRoads(),
		Tiles:        tiles,
	}
}
