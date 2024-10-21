package logger

import (
	"os"
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/position"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/test"
	pb "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/proto"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
)

//nolint:gocyclo// Cyclomatic complexity is not a problem in case of these tests
func TestLoggerWriteRead(t *testing.T) {
	filename := "test_file.pb"

	log, err := NewFromFile(filename)
	if err != nil {
		t.Fatal(err.Error())
	}

	defer os.Remove(filename)

	expectedGameID := 0
	expectedSeed := 0
	expectedPlayerCount := 2
	expectedStartingTile := tilesets.StandardTileSet().StartingTile

	err = log.LogEvent(NewStartEntry(uint32(expectedGameID), uint32(expectedSeed), uint8(expectedPlayerCount), expectedStartingTile))
	if err != nil {
		t.Fatal(err.Error())
	}

	playerID := elements.ID(1)
	expectedTile := test.GetTestPlacedTile()
	err = log.LogEvent(NewPlaceTileEntry(playerID, expectedTile))
	if err != nil {
		t.Fatal(err.Error())
	}

	expectedScores := elements.NewScoreReport()
	expectedScores.ReceivedPoints[playerID] = 1
	expectedScores.ReceivedPoints[elements.ID(2)] = 0
	expectedMeeple := elements.MeepleWithPosition{Meeple: elements.Meeple{Type: elements.NormalMeeple, PlayerID: 1}, Position: position.New(1, 1)}
	expectedScores.ReturnedMeeples[playerID] = make([]elements.MeepleWithPosition, 0)
	expectedScores.ReturnedMeeples[playerID] = append(expectedScores.ReturnedMeeples[playerID], expectedMeeple)
	expectedScores.ReturnedMeeples[elements.ID(2)] = make([]elements.MeepleWithPosition, 0)
	err = log.LogEvent(NewScoreEntry(ScoreEvent, expectedScores))
	if err != nil {
		t.Fatal(err.Error())
	}

	expectedFinalScores := elements.NewScoreReport()
	expectedFinalScores.ReceivedPoints[playerID] = 15
	expectedFinalScores.ReceivedPoints[elements.ID(2)] = 12
	err = log.LogEvent(NewScoreEntry(FinalScoreEvent, expectedFinalScores))
	if err != nil {
		t.Fatal(err.Error())
	}

	reader, err := NewReaderFromFile(filename)
	if err != nil {
		t.Fatal(err.Error())
	}

	channel := reader.ReadLogs()
	for entry := range channel {
		switch content := entry.GetContent().(type) {
		case *pb.Entry_StartEntryContent:
			if entry.Event != pb.EventType_EVENT_TYPE_START_EVENT {
				t.Fatalf("expected %#v, got %#v instead", pb.EventType_EVENT_TYPE_START_EVENT, entry.Event)
			}
			if content.StartEntryContent.GameId != uint32(expectedGameID) {
				t.Fatalf("expected %#v, got %#v instead", expectedGameID, content.StartEntryContent.GameId)
			}
			if content.StartEntryContent.GameSeed != uint32(expectedSeed) {
				t.Fatalf("expected %#v, got %#v instead", expectedSeed, content.StartEntryContent.GameSeed)
			}
			if content.StartEntryContent.PlayerCount != uint32(expectedPlayerCount) {
				t.Fatalf("expected %#v, got %#v instead", expectedPlayerCount, content.StartEntryContent.PlayerCount)
			}
		case *pb.Entry_PlaceTileEntryContent:
			if entry.Event != pb.EventType_EVENT_TYPE_PLACE_TILE_EVENT {
				t.Fatalf("expected %#v, got %#v instead", pb.EventType_EVENT_TYPE_PLACE_TILE_EVENT, entry.Event)
			}
			if content.PlaceTileEntryContent.PlayerId != uint32(playerID) {
				t.Fatalf("expected %#v, got %#v instead", playerID, content.PlaceTileEntryContent.PlayerId)
			}
			if content.PlaceTileEntryContent.Move.Position.X != int32(expectedTile.Position.X()) {
				t.Fatalf("expected %#v, got %#v instead", expectedTile.Position.X(), content.PlaceTileEntryContent.Move.Position.X)
			}
			if content.PlaceTileEntryContent.Move.Position.Y != int32(expectedTile.Position.Y()) {
				t.Fatalf("expected %#v, got %#v instead", expectedTile.Position.Y(), content.PlaceTileEntryContent.Move.Position.Y)
			}
		case *pb.Entry_ScoreEntryContent:
			if entry.Event == pb.EventType_EVENT_TYPE_SCORE_EVENT {
				for _, points := range content.ScoreEntryContent.ScoreReport.ReceivedPoints {
					if points.Score != expectedScores.ReceivedPoints[elements.ID(points.PlayerId)] {
						t.Fatalf("SCORE_EVENT: Player %#v expected %#v, got %#v instead",
							points.PlayerId,
							expectedScores.ReceivedPoints[elements.ID(points.PlayerId)],
							points.Score)
					}
				}
				for _, meeple := range content.ScoreEntryContent.ScoreReport.ReturnedMeeples {
					for _, m := range expectedScores.ReturnedMeeples[elements.ID(meeple.PlayerId)] {
						if meeple.Meeple.MeepleType != pb.MeepleType(m.Meeple.Type) {
							t.Fatalf("SCORE_EVENT: Player %#v expected %#v, got %#v instead",
								meeple.PlayerId,
								m.Meeple.Type,
								meeple.Meeple.MeepleType)
						}
					}
				}
			} else if entry.Event == pb.EventType_EVENT_TYPE_FINAL_SCORE_EVENT {
				for _, points := range content.ScoreEntryContent.ScoreReport.ReceivedPoints {
					if points.Score != expectedFinalScores.ReceivedPoints[elements.ID(points.PlayerId)] {
						t.Fatalf("FINAL_SCORE_EVENT: Player %#v expected %#v, got %#v instead",
							points.PlayerId,
							expectedFinalScores.ReceivedPoints[elements.ID(points.PlayerId)],
							entry.Event)
					}
				}
			} else {
				t.Fatalf("expected SCORE_EVENT or FINAL_SCORE_ENTRY, got %#v instead", entry.Event)
			}
		default:
			t.Fatalf("Unknown entry content")
		}
	}
	log.Close()
	reader.Close()
}

func TestLoggerReadWhileStillWriting(t *testing.T) {
	filename := "test_file.pb"

	log, err := NewFromFile(filename)
	if err != nil {
		t.Fatal(err.Error())
	}

	defer os.Remove(filename)

	expectedGameID := 0
	expectedSeed := 0
	expectedPlayerCount := 2
	expectedStartingTile := tilesets.StandardTileSet().StartingTile

	err = log.LogEvent(NewStartEntry(uint32(expectedGameID), uint32(expectedSeed), uint8(expectedPlayerCount), expectedStartingTile))
	if err != nil {
		t.Fatal(err.Error())
	}

	playerID := elements.ID(1)
	expectedTile := test.GetTestPlacedTile()
	err = log.LogEvent(NewPlaceTileEntry(playerID, expectedTile))
	if err != nil {
		t.Fatal(err.Error())
	}

	expectedScores := elements.NewScoreReport()
	expectedScores.ReceivedPoints[playerID] = 1
	expectedScores.ReceivedPoints[elements.ID(2)] = 0
	expectedMeeple := elements.MeepleWithPosition{Meeple: elements.Meeple{Type: elements.NormalMeeple, PlayerID: 1}, Position: position.New(1, 1)}
	expectedScores.ReturnedMeeples[playerID] = make([]elements.MeepleWithPosition, 0)
	expectedScores.ReturnedMeeples[playerID] = append(expectedScores.ReturnedMeeples[playerID], expectedMeeple)
	expectedScores.ReturnedMeeples[elements.ID(2)] = make([]elements.MeepleWithPosition, 0)
	err = log.LogEvent(NewScoreEntry(ScoreEvent, expectedScores))
	if err != nil {
		t.Fatal(err.Error())
	}

	reader, err := NewReaderFromFile(filename)
	if err != nil {
		t.Fatal(err.Error())
	}

	channel := reader.ReadLogs()
	for entry := range channel {
		switch content := entry.GetContent().(type) {
		case *pb.Entry_StartEntryContent:
			if entry.Event != pb.EventType_EVENT_TYPE_START_EVENT {
				t.Fatalf("expected %#v, got %#v instead", pb.EventType_EVENT_TYPE_START_EVENT, entry.Event)
			}
			if content.StartEntryContent.GameId != uint32(expectedGameID) {
				t.Fatalf("expected %#v, got %#v instead", expectedGameID, content.StartEntryContent.GameId)
			}
			if content.StartEntryContent.GameSeed != uint32(expectedSeed) {
				t.Fatalf("expected %#v, got %#v instead", expectedSeed, content.StartEntryContent.GameSeed)
			}
			if content.StartEntryContent.PlayerCount != uint32(expectedPlayerCount) {
				t.Fatalf("expected %#v, got %#v instead", expectedPlayerCount, content.StartEntryContent.PlayerCount)
			}
		case *pb.Entry_PlaceTileEntryContent:
			if entry.Event != pb.EventType_EVENT_TYPE_PLACE_TILE_EVENT {
				t.Fatalf("expected %#v, got %#v instead", pb.EventType_EVENT_TYPE_PLACE_TILE_EVENT, entry.Event)
			}
			if content.PlaceTileEntryContent.PlayerId != uint32(playerID) {
				t.Fatalf("expected %#v, got %#v instead", playerID, content.PlaceTileEntryContent.PlayerId)
			}
			if content.PlaceTileEntryContent.Move.Position.X != int32(expectedTile.Position.X()) {
				t.Fatalf("expected %#v, got %#v instead", expectedTile.Position.X(), content.PlaceTileEntryContent.Move.Position.X)
			}
			if content.PlaceTileEntryContent.Move.Position.Y != int32(expectedTile.Position.Y()) {
				t.Fatalf("expected %#v, got %#v instead", expectedTile.Position.Y(), content.PlaceTileEntryContent.Move.Position.Y)
			}
		case *pb.Entry_ScoreEntryContent:
			if entry.Event == pb.EventType_EVENT_TYPE_SCORE_EVENT {
				for _, points := range content.ScoreEntryContent.ScoreReport.ReceivedPoints {
					if points.Score != expectedScores.ReceivedPoints[elements.ID(points.PlayerId)] {
						t.Fatalf("SCORE_EVENT: Player %#v expected %#v, got %#v instead",
							points.PlayerId,
							expectedScores.ReceivedPoints[elements.ID(points.PlayerId)],
							points.Score)
					}
				}
				for _, meeple := range content.ScoreEntryContent.ScoreReport.ReturnedMeeples {
					for _, m := range expectedScores.ReturnedMeeples[elements.ID(meeple.PlayerId)] {
						if meeple.Meeple.MeepleType != pb.MeepleType(m.Meeple.Type) {
							t.Fatalf("SCORE_EVENT: Player %#v expected %#v, got %#v instead",
								meeple.PlayerId,
								m.Meeple.Type,
								meeple.Meeple.MeepleType)
						}
					}
				}
			} else {
				t.Fatalf("expected SCORE_EVENT, got %#v instead", entry.Event)
			}
		default:
			t.Fatalf("Unknown entry content")
		}
	}

	expectedFinalScores := elements.NewScoreReport()
	expectedFinalScores.ReceivedPoints[playerID] = 15
	expectedFinalScores.ReceivedPoints[elements.ID(2)] = 12
	err = log.LogEvent(NewScoreEntry(FinalScoreEvent, expectedFinalScores))
	if err != nil {
		t.Fatal(err.Error())
	}

	for entry := range channel {
		switch content := entry.GetContent().(type) {
		case *pb.Entry_ScoreEntryContent:
			if entry.Event == pb.EventType_EVENT_TYPE_FINAL_SCORE_EVENT {
				for _, points := range content.ScoreEntryContent.ScoreReport.ReceivedPoints {
					if points.Score != expectedFinalScores.ReceivedPoints[elements.ID(points.PlayerId)] {
						t.Fatalf("FINAL_SCORE_EVENT: Player %#v expected %#v, got %#v instead",
							points.PlayerId,
							expectedFinalScores.ReceivedPoints[elements.ID(points.PlayerId)],
							entry.Event)
					}
				}
			} else {
				t.Fatalf("expected FINAL_SCORE_ENTRY, got %#v instead", entry.Event)
			}
		}
	}

	log.Close()
	reader.Close()
}
