package elements

import "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"

// mutable type
type Board interface {
	TileCount() int
	Tiles() []PlacedTile
	GetTileAt(pos Position) (PlacedTile, bool)
	GetLegalMovesFor(tile tiles.Tile) []LegalMove
	HasValidPlacement(tile tiles.Tile) bool
	CanBePlaced(tile PlacedTile) bool
	PlaceTile(tile PlacedTile) (ScoreReport, error)
}
