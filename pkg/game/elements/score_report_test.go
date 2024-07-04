package elements

import (
	"reflect"
	"testing"
)

func TestUpdateScoreReport(t *testing.T) {
	report := NewScoreReport()
	report.ReceivedPoints = map[ID]uint32{
		1: 10,
		2: 5,
	}
	report.ReturnedMeeples = map[ID][]uint8{
		1: {1, 0, 1},
		2: {0, 1, 1},
		3: {0, 1, 0},
	}

	otherReport := NewScoreReport()
	otherReport.ReceivedPoints = map[ID]uint32{
		1: 10,
		3: 7,
	}
	otherReport.ReturnedMeeples = map[ID][]uint8{
		1: {1, 0, 0},
		2: {0, 0, 1},
		3: {1, 1, 1},
	}

	expectedReport := NewScoreReport()
	expectedReport.ReceivedPoints = map[ID]uint32{
		1: 20,
		2: 5,
		3: 7,
	}
	expectedReport.ReturnedMeeples = map[ID][]uint8{
		1: {2, 0, 1},
		2: {0, 1, 2},
		3: {1, 2, 1},
	}

	report.Update(otherReport)

	if !reflect.DeepEqual(report, expectedReport) {
		t.Fatalf("expected %#v,\ngot %#v instead", expectedReport, report)
	}
}

func TestUpdateEmptyScoreReport(t *testing.T) {
	report := NewScoreReport()

	otherReport := NewScoreReport()
	otherReport.ReceivedPoints = map[ID]uint32{
		1: 10,
		3: 7,
	}
	otherReport.ReturnedMeeples = map[ID][]uint8{
		1: {1, 0, 0},
		2: {0, 0, 1},
		3: {1, 1, 1},
	}

	report.Update(otherReport)

	if !reflect.DeepEqual(report, otherReport) {
		t.Fatalf("expected %#v,\ngot %#v instead", otherReport, report)
	}
}

func TestUpdateScoreReportWithEmptyReport(t *testing.T) {
	report := NewScoreReport()
	report.ReceivedPoints = map[ID]uint32{
		1: 10,
		3: 7,
	}
	report.ReturnedMeeples = map[ID][]uint8{
		1: {1, 0, 0},
		2: {0, 0, 1},
		3: {1, 1, 1},
	}

	emptyReport := NewScoreReport()

	expectedReport := NewScoreReport()
	expectedReport.ReceivedPoints = map[ID]uint32{
		1: 10,
		3: 7,
	}
	expectedReport.ReturnedMeeples = map[ID][]uint8{
		1: {1, 0, 0},
		2: {0, 0, 1},
		3: {1, 1, 1},
	}

	report.Update(emptyReport)

	if !reflect.DeepEqual(report, expectedReport) {
		t.Fatalf("expected %#v,\ngot %#v instead", expectedReport, report)
	}
}

func TestUpdateScoreReportWithDifferentReturnedMeeplesLength(t *testing.T) {
	report1 := NewScoreReport()
	report1.ReceivedPoints = map[ID]uint32{
		1: 1,
		2: 3,
	}
	report1.ReturnedMeeples = map[ID][]uint8{
		1: {1},
	}

	report2 := NewScoreReport()
	report2.ReceivedPoints = map[ID]uint32{
		1: 2,
		2: 5,
	}
	report2.ReturnedMeeples = map[ID][]uint8{
		1: {0, 1},
		2: {0, 2},
	}

	expectedReport := NewScoreReport()
	expectedReport.ReceivedPoints = map[ID]uint32{
		1: 3,
		2: 8,
	}
	expectedReport.ReturnedMeeples = map[ID][]uint8{
		1: {1, 1},
		2: {0, 2},
	}

	report1.Update(report2)

	if !reflect.DeepEqual(report1, expectedReport) {
		t.Fatalf("expected %#v, got %#v instead", expectedReport, report1)
	}
}
