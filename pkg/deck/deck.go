package deck

import (
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/stack"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
)

type Deck struct {
	*stack.Stack[tiles.Tile]
	StartingTile tiles.Tile
}

func (deck Deck) TileSet() tilesets.TileSet {
	return tilesets.TileSet{
		Tiles:        deck.GetTiles(),
		StartingTile: deck.StartingTile,
	}
}
