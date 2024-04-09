package elements

type Player interface {
	ID() uint8
	MeepleCount() uint8
	SetMeepleCount(value uint8)
	Score() uint32
	SetScore(value uint32)
	PlaceTile(board Board, tile PlacedTile) (ScoreReport, error)
}
