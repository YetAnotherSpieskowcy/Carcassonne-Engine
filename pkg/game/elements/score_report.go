package elements

type ScoreReport struct {
	// ReceivedPoints[playerID (uint8)] = player's received points
	ReceivedPoints map[uint8]uint32
	// ReturnedMeeples[playerID (uint8)][meeple type (MeepleType)] = number of returned meeples
	// for reference, see also: player.meepleCounts
	ReturnedMeeples map[uint8][]uint8
}

func (report *ScoreReport) JoinReport(other ScoreReport) {
	var existKey bool
	// join received points
	for playerID, score := range other.ReceivedPoints {
		_, existKey = report.ReceivedPoints[playerID]
		if !existKey {
			// create field
			report.ReceivedPoints[playerID] = 0
		}

		// add points
		report.ReceivedPoints[playerID] += score
	}

	// join returned meeples
	for playerID, meepleArray := range other.ReturnedMeeples {
		_, existKey = report.ReturnedMeeples[playerID]
		if !existKey {
			// create field
			report.ReturnedMeeples[playerID] = []uint8{}
		}

		// compare length
		if len(report.ReturnedMeeples[playerID]) < len(meepleArray) {
			// lengthen the array with zeros
			report.ReturnedMeeples[playerID] = append(report.ReturnedMeeples[playerID], make([]uint8, len(meepleArray)-len(report.ReturnedMeeples[playerID]))...)
		}

		for meepleType, meepleCount := range meepleArray {
			report.ReturnedMeeples[playerID][meepleType] += meepleCount
		}
	}
}

func NewScoreReport() ScoreReport {
	report := ScoreReport{}

	report.ReceivedPoints = make(map[uint8]uint32)
	report.ReturnedMeeples = make(map[uint8][]uint8)
	return report
}
