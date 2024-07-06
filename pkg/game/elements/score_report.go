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
		scoreReport.ReceivedPoints[playerID] = uint32(score)
	}

	for _, meeple := range meeples {
		_, ok := scoreReport.ReturnedMeeples[uint8(meeple.PlayerID)]
		if !ok {
			scoreReport.ReturnedMeeples[uint8(meeple.PlayerID)] = []uint8{0, 0}
		}
		scoreReport.ReturnedMeeples[uint8(meeple.PlayerID)][meeple.Type]++
	}

	return scoreReport
}
