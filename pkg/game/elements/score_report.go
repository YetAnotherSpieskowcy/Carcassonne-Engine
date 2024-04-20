package elements

type ScoreReport struct {
	// ReceivedPoints[playerID (uint8)] = player's received points
	ReceivedPoints map[uint8]uint32
	// ReturnedMeeples[playerID (uint8)][meeple type (MeepleType)] = number of returned meeples
	// for reference, see also: player.meepleCounts
	ReturnedMeeples map[uint8][]uint8
}

func MakeScoreReport() ScoreReport {
	report := ScoreReport{}

	report.ReceivedPoints = make(map[uint8]uint32)
	report.ReturnedMeeples = make(map[uint8][]uint8)
	return report
}
