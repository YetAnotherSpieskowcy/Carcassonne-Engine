package elements

type ScoreReport struct {
	// ReceivedPoints[playerID (uint8)] = player's received points
	ReceivedPoints map[uint8]uint32
	// ReturnedMeeples[playerID (uint8)][meeple type (MeepleType)] = number of returned meeples
	// for reference, see also: player.meepleCounts
	ReturnedMeeples map[uint8][]uint8
}

// Adds the contents of otherReport to the contents of this score report
func (report *ScoreReport) Update(otherReport ScoreReport) {
	for playerID, score := range otherReport.ReceivedPoints {
		report.ReceivedPoints[playerID] += score
	}

	for playerID, meeples := range otherReport.ReturnedMeeples {
		if _, ok := report.ReturnedMeeples[playerID]; ok {
			for meepleType, count := range meeples {
				report.ReturnedMeeples[playerID][meepleType] += count
			}
		} else {
			report.ReturnedMeeples[playerID] = append(report.ReturnedMeeples[playerID], meeples...)
		}
	}
}
