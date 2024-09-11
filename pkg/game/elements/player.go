package elements

type ID uint8

const (
	NonePlayer ID = iota
)

/*
Same parameters as player, but everything is public.
*/
type SerializedPlayer struct {
	ID           ID
	MeepleCounts []uint8
	Score        uint32
}

type Player interface {
	DeepClone() Player
	ID() ID
	MeepleCount(meepleType MeepleType) uint8
	SetMeepleCount(meepleType MeepleType, value uint8)
	Score() uint32
	SetScore(value uint32)
	// how am I supposed to name this sensibly...
	GetEligibleMovesFrom(moves []PlacedTile) []PlacedTile
	// how am I supposed to name this sensibly...
	IsEligibleFor(move PlacedTile) bool
	PlaceTile(board Board, move PlacedTile) (ScoreReport, error)
	Serialized() SerializedPlayer
}
