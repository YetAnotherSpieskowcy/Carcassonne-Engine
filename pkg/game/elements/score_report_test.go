package elements

import (
	"reflect"
	"testing"
)

func TestJoinReport(t *testing.T) {
	report1 := NewScoreReport()
	report1.ReceivedPoints = map[uint8]uint32{
		1: 1,
		2: 3,
	}
	report1.ReturnedMeeples = map[uint8][]uint8{
		1: []uint8{1},
	}

	report2 := NewScoreReport()
	report2.ReceivedPoints = map[uint8]uint32{
		1: 2,
		2: 5,
	}
	report2.ReturnedMeeples = map[uint8][]uint8{
		1: []uint8{0, 1},
		2: []uint8{0, 2},
	}

	expectedReport := NewScoreReport()
	expectedReport.ReceivedPoints = map[uint8]uint32{
		1: 3,
		2: 8,
	}
	expectedReport.ReturnedMeeples = map[uint8][]uint8{
		1: []uint8{1, 1},
		2: []uint8{0, 2},
	}

	report1.JoinReport(report2)

	if !reflect.DeepEqual(report1, expectedReport) {
		t.Fatalf("expected %#v, got %#v instead", expectedReport, report1)
	}

}
