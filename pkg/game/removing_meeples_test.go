package game

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/deck"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/stack"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/feature"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/side"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles/tiletemplates"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
)

/*
Simple road using two monasteries.

	M1 - 0 - M2
*/
func TestRemoveSingleMeeple(t *testing.T) {
	var tile tiles.Tile

	// ------ create tileset --------
	var tiles []tiles.Tile

	// 2 monastery with road
	tiles = append(tiles, tiletemplates.MonasteryWithSingleRoad())
	tiles = append(tiles, tiletemplates.MonasteryWithSingleRoad())

	tileset := tilesets.TileSet{
		StartingTile: tiletemplates.SingleCityEdgeStraightRoads(),
		Tiles:        tiles,
	}

	// ------ create game --------
	game, err := NewFromTileSet(tileset, nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	// ------ first turn --------

	tile, err = game.GetCurrentTile()
	if err != nil {
		t.Fatal(err.Error())
	}
	var player = game.CurrentPlayer()
	var ptile = elements.ToPlacedTile(tile.Rotate(3)) // make road to right

	ptile.Position = position.New(-1, 0)
	ptile.GetPlacedFeatureAtSide(side.Right, feature.Road).Meeple = elements.Meeple{Type: elements.NormalMeeple, PlayerID: player.ID()}
	err = game.PlayTurn(ptile)
	if err != nil {
		t.Fatal(err.Error())
	}

	// ------ second turn --------

	tile, err = game.GetCurrentTile()
	if err != nil {
		t.Fatal(err.Error())
	}
	ptile = elements.ToPlacedTile(tile.Rotate(1)) // make road to left

	ptile.Position = position.New(1, 0)
	err = game.PlayTurn(ptile)
	if err != nil {
		t.Fatal(err.Error())
	}

	// ------ Check if meeple was removed --------
	ptile, _ = game.board.GetTileAt(position.New(-1, 0))
	meeple := ptile.GetPlacedFeatureAtSide(side.Right, feature.Road).Meeple
	var expectedMeeple = elements.Meeple{elements.NoneMeeple, elements.ID(0)}
	if meeple != expectedMeeple {
		t.Fatalf("Removing meeple failed! \nFound: \n%#v,\nexpected:\n%#v", meeple, expectedMeeple)
	}

}

/*
Simple road using two monasteries and two road turns.

	Roads:

	4 -	0 -	1
	|
	3 -	2
*/
func TestRemoveTwoMeeples(t *testing.T) {
	var tile tiles.Tile
	// ------ create tileset --------
	var tiles []tiles.Tile

	// 2 monastery with road and two road turns
	tiles = append(tiles, tiletemplates.MonasteryWithSingleRoad())
	tiles = append(tiles, tiletemplates.MonasteryWithSingleRoad())
	tiles = append(tiles, tiletemplates.RoadsTurn())
	tiles = append(tiles, tiletemplates.RoadsTurn())

	tileSet := tilesets.TileSet{
		StartingTile: tiletemplates.SingleCityEdgeStraightRoads(),
		Tiles:        tiles,
	}

	deckStack := stack.NewOrdered(tileSet.Tiles)
	deck := deck.Deck{Stack: &deckStack, StartingTile: tileSet.StartingTile}

	// ------ create game --------
	game, err := NewFromDeck(deck, nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	// ------ first turn --------
	tile, err = game.GetCurrentTile()
	if err != nil {
		t.Fatal(err.Error())
	}
	var player = game.CurrentPlayer()
	var ptile = elements.ToPlacedTile(tile.Rotate(1)) // make road to left

	ptile.Position = position.New(1, 0)
	ptile.GetPlacedFeatureAtSide(side.Left, feature.Road).Meeple = elements.Meeple{Type: elements.NormalMeeple, PlayerID: player.ID()}
	err = game.PlayTurn(ptile)
	if err != nil {
		t.Fatal(err.Error())
	}

	// ------ second turn --------

	tile, err = game.GetCurrentTile()
	if err != nil {
		t.Fatal(err.Error())
	}
	player = game.CurrentPlayer()
	ptile = elements.ToPlacedTile(tile.Rotate(1)) // make road to left

	ptile.Position = position.New(0, -1)
	ptile.GetPlacedFeatureAtSide(side.Left, feature.Road).Meeple = elements.Meeple{Type: elements.NormalMeeple, PlayerID: player.ID()}
	err = game.PlayTurn(ptile)
	if err != nil {
		t.Fatal(err.Error())
	}

	// ------ third turn --------
	tile, err = game.GetCurrentTile()
	if err != nil {
		t.Fatal(err.Error())
	}
	ptile = elements.ToPlacedTile(tile.Rotate(2))

	ptile.Position = position.New(-1, -1)
	err = game.PlayTurn(ptile)
	if err != nil {
		t.Fatal(err.Error())
	}

	// ------ fourth turn --------
	tile, err = game.GetCurrentTile()
	if err != nil {
		t.Fatal(err.Error())
	}
	ptile = elements.ToPlacedTile(tile.Rotate(3))

	ptile.Position = position.New(-1, 0)
	game.PlayTurn(ptile)

	// ------ Check if meeples were removed --------
	ptile, _ = game.board.GetTileAt(position.New(1, 0))

	var meeple = ptile.GetPlacedFeatureAtSide(side.Left, feature.Road).Meeple
	var expectedMeeple = elements.Meeple{elements.NoneMeeple, elements.ID(0)}
	if meeple != expectedMeeple {
		t.Fatalf("Removing first meeple failed! \nFound: \n%#v,\nexpected:\n%#v", meeple, expectedMeeple)
	}

	ptile, _ = game.board.GetTileAt(position.New(0, -1))
	meeple = ptile.GetPlacedFeatureAtSide(side.Left, feature.Road).Meeple
	if meeple != expectedMeeple {
		t.Fatalf("Removing second meeple failed! \nFound: \n%#v,\nexpected:\n%#v", meeple, expectedMeeple)
	}
}
