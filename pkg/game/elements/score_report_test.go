package elements

import (
	"reflect"
	"testing"
)

func TestUpdateScoreReport(t *testing.T) {
	report := ScoreReport{
		ReceivedPoints: map[uint8]uint32{
			1: 10,
			2: 5,
		},
		ReturnedMeeples: map[uint8][]uint8{
			1: {1, 0, 1},
			2: {0, 1, 1},
			3: {0, 1, 0},
		},
	}
	otherReport := ScoreReport{
		ReceivedPoints: map[uint8]uint32{
			1: 10,
			3: 7,
		},
		ReturnedMeeples: map[uint8][]uint8{
			1: {1, 0, 0},
			2: {0, 0, 1},
			3: {1, 1, 1},
		},
	}
	expected := ScoreReport{
		ReceivedPoints: map[uint8]uint32{
			1: 20,
			2: 5,
			3: 7,
		},
		ReturnedMeeples: map[uint8][]uint8{
			1: {2, 0, 1},
			2: {0, 1, 2},
			3: {1, 2, 1},
		},
	}

	report.Update(otherReport)

	if !reflect.DeepEqual(report, expected) {
		t.Fatalf("expected %#v,\ngot %#v instead", expected, report)
	}
}

func TestUpdateEmptyScoreReport(t *testing.T) {
	report := NewScoreReport()
	otherReport := ScoreReport{
		ReceivedPoints: map[uint8]uint32{
			1: 10,
			3: 7,
		},
		ReturnedMeeples: map[uint8][]uint8{
			1: {1, 0, 0},
			2: {0, 0, 1},
			3: {1, 1, 1},
		},
	}
	report.Update(otherReport)

	if !reflect.DeepEqual(report, otherReport) {
		t.Fatalf("expected %#v,\ngot %#v instead", otherReport, report)
	}
}
