package elements

import (
	"reflect"
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
)

func SimpleMeepleWithPosition(meeple Meeple, pos position.Position) MeepleWithPosition {
	return NewMeepleWithPosition(meeple, pos)
}

func TestUpdateScoreReport(t *testing.T) {
	report := NewScoreReport()
	report.ReceivedPoints = map[ID]uint32{
		1: 10,
		2: 5,
	}
	report.ReturnedMeeples = map[ID][]MeepleWithPosition{
		1: {
			SimpleMeepleWithPosition(Meeple{MeepleType(1), ID(1)}, position.New(0, 0)),
			SimpleMeepleWithPosition(Meeple{MeepleType(3), ID(1)}, position.New(0, 1))},
		2: {
			SimpleMeepleWithPosition(Meeple{MeepleType(2), ID(2)}, position.New(0, 2)),
			SimpleMeepleWithPosition(Meeple{MeepleType(3), ID(2)}, position.New(0, 3))},
		3: {
			SimpleMeepleWithPosition(Meeple{MeepleType(2), ID(3)}, position.New(0, 4))},
	}

	otherReport := NewScoreReport()
	otherReport.ReceivedPoints = map[ID]uint32{
		1: 10,
		3: 7,
	}
	otherReport.ReturnedMeeples = map[ID][]MeepleWithPosition{
		1: {SimpleMeepleWithPosition(Meeple{MeepleType(1), ID(1)}, position.New(0, 5))},
		2: {SimpleMeepleWithPosition(Meeple{MeepleType(3), ID(2)}, position.New(0, 6))},
		3: {SimpleMeepleWithPosition(Meeple{MeepleType(1), ID(3)}, position.New(0, 7)),
			SimpleMeepleWithPosition(Meeple{MeepleType(2), ID(3)}, position.New(0, 8)),
			SimpleMeepleWithPosition(Meeple{MeepleType(3), ID(3)}, position.New(0, 9))},
	}

	expectedReport := NewScoreReport()
	expectedReport.ReceivedPoints = map[ID]uint32{
		1: 20,
		2: 5,
		3: 7,
	}
	expectedReport.ReturnedMeeples = map[ID][]MeepleWithPosition{
		1: {SimpleMeepleWithPosition(Meeple{MeepleType(1), ID(1)}, position.New(0, 0)),
			SimpleMeepleWithPosition(Meeple{MeepleType(3), ID(1)}, position.New(0, 1)),
			SimpleMeepleWithPosition(Meeple{MeepleType(1), ID(1)}, position.New(0, 5))},
		2: {SimpleMeepleWithPosition(Meeple{MeepleType(2), ID(2)}, position.New(0, 2)),
			SimpleMeepleWithPosition(Meeple{MeepleType(3), ID(2)}, position.New(0, 3)),
			SimpleMeepleWithPosition(Meeple{MeepleType(3), ID(2)}, position.New(0, 6))},
		3: {SimpleMeepleWithPosition(Meeple{MeepleType(2), ID(3)}, position.New(0, 4)),
			SimpleMeepleWithPosition(Meeple{MeepleType(1), ID(3)}, position.New(0, 7)),
			SimpleMeepleWithPosition(Meeple{MeepleType(2), ID(3)}, position.New(0, 8)),
			SimpleMeepleWithPosition(Meeple{MeepleType(3), ID(3)}, position.New(0, 9))},
	}

	report.Join(otherReport)

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
	otherReport.ReturnedMeeples = map[ID][]MeepleWithPosition{
		1: {SimpleMeepleWithPosition(Meeple{MeepleType(1), ID(1)}, position.New(0, 0))},
		2: {SimpleMeepleWithPosition(Meeple{MeepleType(3), ID(2)}, position.New(0, 1))},
		3: {SimpleMeepleWithPosition(Meeple{MeepleType(1), ID(3)}, position.New(0, 2)),
			SimpleMeepleWithPosition(Meeple{MeepleType(2), ID(3)}, position.New(0, 3)),
			SimpleMeepleWithPosition(Meeple{MeepleType(3), ID(3)}, position.New(0, 4))},
	}

	report.Join(otherReport)

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
	report.ReturnedMeeples = map[ID][]MeepleWithPosition{
		1: {SimpleMeepleWithPosition(Meeple{MeepleType(1), ID(1)}, position.New(0, 0))},
		2: {SimpleMeepleWithPosition(Meeple{MeepleType(3), ID(2)}, position.New(0, 1))},
		3: {SimpleMeepleWithPosition(Meeple{MeepleType(1), ID(3)}, position.New(0, 2)),
			SimpleMeepleWithPosition(Meeple{MeepleType(2), ID(3)}, position.New(0, 3)),
			SimpleMeepleWithPosition(Meeple{MeepleType(3), ID(3)}, position.New(0, 4))},
	}

	emptyReport := NewScoreReport()

	expectedReport := NewScoreReport()
	expectedReport.ReceivedPoints = map[ID]uint32{
		1: 10,
		3: 7,
	}
	expectedReport.ReturnedMeeples = map[ID][]MeepleWithPosition{
		1: {SimpleMeepleWithPosition(Meeple{MeepleType(1), ID(1)}, position.New(0, 0))},
		2: {SimpleMeepleWithPosition(Meeple{MeepleType(3), ID(2)}, position.New(0, 1))},
		3: {SimpleMeepleWithPosition(Meeple{MeepleType(1), ID(3)}, position.New(0, 2)),
			SimpleMeepleWithPosition(Meeple{MeepleType(2), ID(3)}, position.New(0, 3)),
			SimpleMeepleWithPosition(Meeple{MeepleType(3), ID(3)}, position.New(0, 4))},
	}

	report.Join(emptyReport)

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
	report1.ReturnedMeeples = map[ID][]MeepleWithPosition{
		1: {SimpleMeepleWithPosition(Meeple{MeepleType(1), ID(1)}, position.New(0, 0))},
	}

	report2 := NewScoreReport()
	report2.ReceivedPoints = map[ID]uint32{
		1: 2,
		2: 5,
	}
	report2.ReturnedMeeples = map[ID][]MeepleWithPosition{
		1: {SimpleMeepleWithPosition(Meeple{MeepleType(1), ID(1)}, position.New(0, 1)),
			SimpleMeepleWithPosition(Meeple{MeepleType(2), ID(1)}, position.New(0, 2))},
		2: {SimpleMeepleWithPosition(Meeple{MeepleType(2), ID(2)}, position.New(0, 3)),
			SimpleMeepleWithPosition(Meeple{MeepleType(2), ID(2)}, position.New(0, 4))},
	}

	expectedReport := NewScoreReport()
	expectedReport.ReceivedPoints = map[ID]uint32{
		1: 3,
		2: 8,
	}
	expectedReport.ReturnedMeeples = map[ID][]MeepleWithPosition{
		1: {SimpleMeepleWithPosition(Meeple{MeepleType(1), ID(1)}, position.New(0, 0)),
			SimpleMeepleWithPosition(Meeple{MeepleType(1), ID(1)}, position.New(0, 1)),
			SimpleMeepleWithPosition(Meeple{MeepleType(2), ID(1)}, position.New(0, 2))},
		2: {SimpleMeepleWithPosition(Meeple{MeepleType(2), ID(2)}, position.New(0, 3)),
			SimpleMeepleWithPosition(Meeple{MeepleType(2), ID(2)}, position.New(0, 4))},
	}

	report1.Join(report2)

	if !reflect.DeepEqual(report1, expectedReport) {
		t.Fatalf("expected %#v, got %#v instead", expectedReport, report1)
	}
}

func TestGetPlayersWithMostMeeplesOnePlayer(t *testing.T) {
	meeples := map[ID][]uint8{
		1: {0, 1},
		2: {0, 3},
		3: {0, 2},
	}

	expectedPlayers := []ID{2}
	actualplayers := GetPlayersWithMostMeeples(meeples)

	if !reflect.DeepEqual(expectedPlayers, actualplayers) {
		t.Fatalf("expected %#v, got %#v instead", expectedPlayers, actualplayers)
	}
}

func TestGetPlayersWithMostMeeplesTwoPlayers(t *testing.T) {
	meeples := map[ID][]uint8{
		1: {0, 1},
		2: {0, 2},
		3: {0, 2},
	}

	expectedPlayers := []ID{2, 3}
	expectedPlayers2 := []ID{3, 2} // the order does not matter
	actualplayers := GetPlayersWithMostMeeples(meeples)

	if !(reflect.DeepEqual(expectedPlayers, actualplayers) || reflect.DeepEqual(expectedPlayers2, actualplayers)) {
		t.Fatalf("expected %#v, got %#v instead", expectedPlayers, actualplayers)
	}
}

func TestMeepleInReportExists(t *testing.T) {
	report := NewScoreReport()
	meeple1 := MeepleWithPosition{
		Meeple: Meeple{
			Type:     NormalMeeple,
			PlayerID: 1,
		},
		Position: position.New(0, 0),
	}

	meeple2 := MeepleWithPosition{
		Meeple: Meeple{
			Type:     NormalMeeple,
			PlayerID: 2,
		},
		Position: position.New(1, 0),
	}

	report.ReturnedMeeples = map[ID][]MeepleWithPosition{
		1: {meeple1},
		2: {meeple2},
	}

	if !report.MeepleInReport(meeple1) {
		t.Fatalf("Meeple1 not found in report!")
	}

	if !report.MeepleInReport(meeple2) {
		t.Fatalf("Meeple2 not found in report!")
	}
}

func TestMeepleInReportNotExists(t *testing.T) {
	report := NewScoreReport()
	dummyMeeple := MeepleWithPosition{
		Meeple: Meeple{
			Type:     NormalMeeple,
			PlayerID: 1,
		},
		Position: position.New(0, 0),
	}

	if report.MeepleInReport(dummyMeeple) {
		t.Fatalf("Meeple should not be in report!")
	}
}
