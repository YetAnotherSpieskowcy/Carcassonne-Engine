package elements

type Player interface {
	ID() uint8
	MeepleCount(meepleType MeepleType) uint8
	SetMeepleCount(meepleType MeepleType, value uint8)
	Score() uint32
	SetScore(value uint32)
	PlaceTile(board Board, tile PlacedTile) (ScoreReport, error)
}
