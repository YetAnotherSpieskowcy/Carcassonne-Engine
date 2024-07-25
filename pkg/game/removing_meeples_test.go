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
	var tile, _ = game.GetCurrentTile()
	var player = game.CurrentPlayer()
	var ptile = elements.ToPlacedTile(tile.Rotate(3)) // make road to right

	ptile.Position = position.New(-1, 0)
	ptile.GetPlacedFeatureAtSide(side.Right, feature.Road).Meeple = elements.Meeple{elements.NormalMeeple, elements.ID(player.ID())}
	game.PlayTurn(ptile)

	// ------ second turn --------

	tile, _ = game.GetCurrentTile()
	player = game.CurrentPlayer()
	ptile = elements.ToPlacedTile(tile.Rotate(1)) // make road to left

	ptile.Position = position.New(1, 0)
	game.PlayTurn(ptile)

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
	var tile, _ = game.GetCurrentTile()
	var player = game.CurrentPlayer()
	var ptile = elements.ToPlacedTile(tile.Rotate(1)) // make road to left

	ptile.Position = position.New(1, 0)
	ptile.GetPlacedFeatureAtSide(side.Left, feature.Road).Meeple = elements.Meeple{elements.NormalMeeple, elements.ID(player.ID())}
	game.PlayTurn(ptile)

	// ------ second turn --------

	tile, _ = game.GetCurrentTile()
	player = game.CurrentPlayer()
	ptile = elements.ToPlacedTile(tile.Rotate(1)) // make road to left

	ptile.Position = position.New(0, -1)
	ptile.GetPlacedFeatureAtSide(side.Left, feature.Road).Meeple = elements.Meeple{elements.NormalMeeple, elements.ID(player.ID())}
	game.PlayTurn(ptile)

	// ------ third turn --------
	tile, _ = game.GetCurrentTile()
	ptile = elements.ToPlacedTile(tile.Rotate(2))

	ptile.Position = position.New(-1, -1)
	game.PlayTurn(ptile)

	// ------ fourth turn --------
	tile, _ = game.GetCurrentTile()
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
