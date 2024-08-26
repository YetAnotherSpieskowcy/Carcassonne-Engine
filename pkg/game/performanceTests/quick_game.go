package performancetests

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/deck"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/stack"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
)

/*
Quick function for playing a simple game in a straight line.
Arugment playGame is used for cases when empty game needs to be measured.
*/
func PlayNTileGame(tileCount int, tile tiles.Tile, playGame bool) error {

	tileSet := tilesets.TileSet{}
	tileSet.StartingTile = tile
	for range tileCount {
		tileSet.Tiles = append(tileSet.Tiles, tile)
	}

	deckStack := stack.NewOrdered(tileSet.Tiles)
	deck := deck.Deck{Stack: &deckStack, StartingTile: tileSet.StartingTile}
	Game, err := game.NewFromDeck(deck, nil)
	if err != nil {
		return err
	}
	ptile := elements.ToPlacedTile(tile)

	if playGame {
		// play game
		for i := range tileCount {
			ptile.Position = position.New(int16(i+1), 0)
			err = Game.PlayTurn(ptile)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
