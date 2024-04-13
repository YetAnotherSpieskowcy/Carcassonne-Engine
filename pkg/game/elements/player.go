package elements

type Player interface {
	ID() uint8
	MeepleCount(meepleType MeepleType) uint8
	SetMeepleCount(meepleType MeepleType, value uint8)
	Score() uint32
	SetScore(value uint32)
	// how am I supposed to name this sensibly...
	GetEligibleMovesFrom(moves []LegalMove) []LegalMove
	// how am I supposed to name this sensibly...
	IsEligibleFor(move LegalMove) bool
	PlaceTile(board Board, move LegalMove) (ScoreReport, error)
}
