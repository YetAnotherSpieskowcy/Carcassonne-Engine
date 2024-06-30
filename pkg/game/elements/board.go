package elements

import "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"

// mutable type
type Board interface {
	TileCount() int
	Tiles() []PlacedTile
	GetTileAt(pos Position) (PlacedTile, bool)
	GetTilePlacementsFor(tile tiles.Tile) []PlacedTile
	TileHasValidPlacement(tile tiles.Tile) bool
	GetLegalMovesFor(tile PlacedTile) []PlacedTile
	CanBePlaced(tile PlacedTile) bool
	PlaceTile(tile PlacedTile) (ScoreReport, error)
}
