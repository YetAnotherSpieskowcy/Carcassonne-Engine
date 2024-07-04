package elements

type ScoreReport struct {
	// ReceivedPoints[playerID (uint8)] = player's received points
	ReceivedPoints map[ID]uint32
	// ReturnedMeeples[playerID (uint8)][meeple type (MeepleType)] = number of returned meeples
	// for reference, see also: player.meepleCounts
	ReturnedMeeples map[ID][]uint8
}

func NewScoreReport() ScoreReport {
	return ScoreReport{
		ReceivedPoints:  map[ID]uint32{},
		ReturnedMeeples: map[ID][]uint8{},
	}
}

// Adds the contents of otherReport to the contents of this score report
func (report *ScoreReport) Update(otherReport ScoreReport) {
	for playerID, score := range otherReport.ReceivedPoints {
		report.ReceivedPoints[playerID] += score
	}

	for playerID, meeples := range otherReport.ReturnedMeeples {
		_, keyExists := report.ReturnedMeeples[playerID]
		if !keyExists {
			report.ReturnedMeeples[playerID] = []uint8{}
		}

		if len(report.ReturnedMeeples[playerID]) < len(meeples) {
			// lengthen the meeples array with zeros
			// this should not be necessary if we assumed that for all reports, report.ReturnedMeeples should be of length = elements.MeepleTypeCount. Todo?
			zerosNumber := len(meeples) - len(report.ReturnedMeeples[playerID])
			report.ReturnedMeeples[playerID] = append(report.ReturnedMeeples[playerID], make([]uint8, zerosNumber)...)
		}

		for meepleType, meepleCount := range meeples {
			report.ReturnedMeeples[playerID][meepleType] += meepleCount
		}
	}
}

/*
Create score report by checking meeples control on the same Fully Connected Feature (like a whole city/road etc), ignoring not scoring meeples.
Returns a score report
*/
func CalculateScoreReportOnMeeples(score int, meeples []Meeple) ScoreReport {
	var mostMeeples = uint8(0)
	var scoredPlayers = []uint8{}
	playerMeeples := make(map[uint8]uint8)
	// count meeples, and find max
	for _, meeple := range meeples {
		_, existKey := playerMeeples[uint8(meeple.PlayerID)]
		if !existKey {
			playerMeeples[uint8(meeple.PlayerID)] = 0
		}
		playerMeeples[uint8(meeple.PlayerID)]++
		if playerMeeples[uint8(meeple.PlayerID)] > mostMeeples {
			mostMeeples = playerMeeples[uint8(meeple.PlayerID)]
		}
	}

	// find players with max
	for playerID, count := range playerMeeples {
		if count == mostMeeples {
			scoredPlayers = append(scoredPlayers, playerID)
		}
	}

	// -------- create report -------------
	scoreReport := NewScoreReport()

	for _, playerID := range scoredPlayers {
		scoreReport.ReceivedPoints[ID(playerID)] = uint32(score)
	}

	for _, meeple := range meeples {
		_, ok := scoreReport.ReturnedMeeples[meeple.PlayerID]
		if !ok {
			scoreReport.ReturnedMeeples[meeple.PlayerID] = []uint8{0, 0}
		}
		scoreReport.ReturnedMeeples[meeple.PlayerID][meeple.MeepleType]++
	}

	return scoreReport
}
