package performancetests

import (
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/deck"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/stack"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
)

/*
Quick function for playing a simple game
The tiles are placed in a straight line from x=0 to x=125. After that a new row is started (x=0, y=1) and so on
*/
func PlayNTileGame(tileCount int, tile tiles.Tile, b *testing.B) error {

	tileSet := tilesets.TileSet{}
	tileSet.StartingTile = tile
	for range tileCount {
		tileSet.Tiles = append(tileSet.Tiles, tile)
	}

	deckStack := stack.NewOrdered(tileSet.Tiles)
	deck := deck.Deck{Stack: &deckStack, StartingTile: tileSet.StartingTile}
	Game, err := game.NewFromDeck(deck, nil, 2)
	if err != nil {
		return err
	}
	ptile := elements.ToPlacedTile(tile)

	// play game
	b.StartTimer()
	x := int8(1)
	y := int8(0)
	for range tileCount {
		ptile.Position = position.New(x, y)
		x++
		if x == 126 {
			x = 0
			y++
		}

		err = Game.PlayTurn(ptile)
		if err != nil {
			return err
		}
	}
	b.StopTimer()

	return nil
}
