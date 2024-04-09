package elements

// mutable type
type Board interface {
	TileCount() int
	Tiles() []PlacedTile
	GetTileAt(pos Position) (PlacedTile, bool)
	GetLegalMovesFor(tile Tile) []LegalMove
	HasValidPlacement(tile Tile) bool
	CanBePlaced(tile PlacedTile) bool
	PlaceTile(tile PlacedTile) (ScoreReport, error)
}
