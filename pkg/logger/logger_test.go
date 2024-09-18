package logger

import (
	"bufio"
	"bytes"
	"encoding/json"
	"os"
	"reflect"
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/deck"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/elements"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/test"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/player"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/stack"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tiles"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/tilesets"
)

func getTestDeck() deck.Deck {
	tileSet := tilesets.StandardTileSet()
	tileSet.Tiles = []tiles.Tile{test.GetTestTile(), test.GetTestTile()}
	deckStack := stack.NewOrdered(tileSet.Tiles)
	return deck.Deck{
		Stack:        &deckStack,
		StartingTile: tileSet.StartingTile,
	}
}

//nolint:gocyclo// Cyclomatic complexity is not a problem in case of these tests
func TestFileLogger(t *testing.T) {
	filename := "test_file.jsonl"
	player1 := player.New(1)
	player2 := player.New(2)

	log, err := NewFromFile(filename)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer os.Remove(filename)

	deck := getTestDeck()
	expectedStartingTile := deck.StartingTile
	expectedStack := deck.GetRemaining()
	expectedPlayerCount := 2

	err = log.LogEvent(StartEvent, NewStartEntryContent(expectedStartingTile, expectedStack, expectedPlayerCount))
	if err != nil {
		t.Fatal(err.Error())
	}

	expectedTile := test.GetTestPlacedTile()
	err = log.LogEvent(PlaceTileEvent, NewPlaceTileEntryContent(player1.ID(), expectedTile))
	if err != nil {
		t.Fatal(err.Error())
	}

	expectedScores := elements.NewScoreReport()
	expectedScores.ReceivedPoints[player1.ID()] = 1
	expectedScores.ReceivedPoints[player2.ID()] = 2
	err = log.LogEvent(ScoreEvent, NewScoreEntryContent(expectedScores))
	if err != nil {
		t.Fatal(err.Error())
	}

	expectedFinalScores := make(map[elements.ID]uint32, 0)
	expectedFinalScores[player1.ID()] = 1
	expectedFinalScores[player2.ID()] = 2
	err = log.LogEvent(EndEvent, NewEndEntryContent(expectedFinalScores))
	if err != nil {
		t.Fatal(err.Error())
	}

	var entryLine Entry
	var startContent StartEntryContent
	var placeTileContent PlaceTileEntryContent
	var scoreContent ScoreEntryContent
	var endContent EndEntryContent

	log.Close()

	file, err := os.Open(filename)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	err = json.Unmarshal([]byte(scanner.Text()), &entryLine)
	if err != nil {
		t.Fatal(err.Error())
	}
	if entryLine.Event != StartEvent {
		t.Fatalf("expected %#v, got %#v instead", StartEvent, entryLine.Event)
	}
	err = json.Unmarshal(entryLine.Content, &startContent)
	if err != nil {
		t.Fatal(err.Error())
	}
	if !reflect.DeepEqual(startContent.StartingTile, expectedStartingTile) {
		t.Fatalf("expected %#v, got %#v instead", expectedStartingTile, startContent.StartingTile)
	}
	if !reflect.DeepEqual(startContent.Stack, expectedStack) {
		t.Fatalf("expected %#v, got %#v instead", expectedStack, startContent.Stack)
	}
	if !reflect.DeepEqual(startContent.PlayerCount, expectedPlayerCount) {
		t.Fatalf("expected %#v, got %#v instead", expectedPlayerCount, startContent.PlayerCount)
	}

	scanner.Scan()
	err = json.Unmarshal([]byte(scanner.Text()), &entryLine)
	if err != nil {
		t.Fatal(err.Error())
	}
	if entryLine.Event != PlaceTileEvent {
		t.Fatalf("expected %#v, got %#v instead", PlaceTileEvent, entryLine.Event)
	}
	err = json.Unmarshal(entryLine.Content, &placeTileContent)
	if err != nil {
		t.Fatal(err.Error())
	}
	if placeTileContent.PlayerID != player1.ID() {
		t.Fatalf("expected %#v, got %#v instead", player1.ID(), placeTileContent.PlayerID)
	}
	if !reflect.DeepEqual(placeTileContent.Move, expectedTile) {
		t.Fatalf("expected %#v, got %#v instead", expectedTile, placeTileContent.Move)
	}

	scanner.Scan()
	err = json.Unmarshal([]byte(scanner.Text()), &entryLine)
	if err != nil {
		t.Fatal(err.Error())
	}
	if entryLine.Event != ScoreEvent {
		t.Fatalf("expected %#v, got %#v instead", ScoreEvent, entryLine.Event)
	}
	err = json.Unmarshal(entryLine.Content, &scoreContent)
	if err != nil {
		t.Fatal(err.Error())
	}
	if !reflect.DeepEqual(scoreContent.Scores, expectedScores) {
		t.Fatalf("expected %#v, got %#v instead", expectedScores, scoreContent.Scores)
	}

	scanner.Scan()
	err = json.Unmarshal([]byte(scanner.Text()), &entryLine)
	if err != nil {
		t.Fatal(err.Error())
	}
	if entryLine.Event != EndEvent {
		t.Fatalf("expected %#v, got %#v instead", EndEvent, entryLine.Event)
	}
	err = json.Unmarshal(entryLine.Content, &endContent)
	if err != nil {
		t.Fatal(err.Error())
	}
	if !reflect.DeepEqual(endContent.FinalScores, expectedFinalScores) {
		t.Fatalf("expected %#v, got %#v instead", expectedScores, endContent.FinalScores)
	}
}

//nolint:gocyclo// Cyclomatic complexity is not a problem in case of these tests
func TestReadLogs(t *testing.T) {
	filename := "test_file.jsonl"
	player1 := player.New(1)
	player2 := player.New(2)

	log, err := NewFromFile(filename)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer os.Remove(filename)

	deck := getTestDeck()
	expectedStartingTile := deck.StartingTile
	expectedStack := deck.GetRemaining()
	expectedPlayerCount := 2

	err = log.LogEvent(StartEvent, NewStartEntryContent(expectedStartingTile, expectedStack, expectedPlayerCount))
	if err != nil {
		t.Fatal(err.Error())
	}

	expectedTile := test.GetTestPlacedTile()
	err = log.LogEvent(PlaceTileEvent, NewPlaceTileEntryContent(player1.ID(), expectedTile))
	if err != nil {
		t.Fatal(err.Error())
	}

	expectedScores := elements.NewScoreReport()
	expectedScores.ReceivedPoints[player1.ID()] = 1
	expectedScores.ReceivedPoints[player2.ID()] = 2
	err = log.LogEvent(ScoreEvent, NewScoreEntryContent(expectedScores))
	if err != nil {
		t.Fatal(err.Error())
	}

	expectedFinalScores := make(map[elements.ID]uint32, 0)
	expectedFinalScores[player1.ID()] = 1
	expectedFinalScores[player2.ID()] = 2
	err = log.LogEvent(EndEvent, NewEndEntryContent(expectedFinalScores))
	if err != nil {
		t.Fatal(err.Error())
	}

	log.Close()

	var startContent StartEntryContent
	var placeTileContent PlaceTileEntryContent
	var scoreContent ScoreEntryContent
	var endContent EndEntryContent

	reader, err := NewReaderFromFile(filename)

	if err != nil {
		t.Fatal(err.Error())
	}

	for e := range reader.ReadLogs() {
		switch e.Event {
		case StartEvent:
			err = json.Unmarshal(e.Content, &startContent)
			if err != nil {
				t.Fatal(err.Error())
			}
			if !reflect.DeepEqual(startContent.StartingTile, expectedStartingTile) {
				t.Fatalf("expected %#v, got %#v instead", expectedStartingTile, startContent.StartingTile)
			}
			if !reflect.DeepEqual(startContent.Stack, expectedStack) {
				t.Fatalf("expected %#v, got %#v instead", expectedStack, startContent.Stack)
			}
			if !reflect.DeepEqual(startContent.PlayerCount, expectedPlayerCount) {
				t.Fatalf("expected %#v, got %#v instead", expectedPlayerCount, startContent.PlayerCount)
			}
		case PlaceTileEvent:
			err = json.Unmarshal(e.Content, &placeTileContent)
			if err != nil {
				t.Fatal(err.Error())
			}
			if placeTileContent.PlayerID != player1.ID() {
				t.Fatalf("expected %#v, got %#v instead", player1.ID(), placeTileContent.PlayerID)
			}
			if !reflect.DeepEqual(placeTileContent.Move, expectedTile) {
				t.Fatalf("expected %#v, got %#v instead", expectedTile, placeTileContent.Move)
			}
		case ScoreEvent:
			err = json.Unmarshal(e.Content, &scoreContent)

			if err != nil {
				t.Fatal(err.Error())
			}

			if !reflect.DeepEqual(scoreContent.Scores, expectedScores) {
				t.Fatalf("expected %#v, got %#v instead", expectedScores, scoreContent.Scores)
			}
		case EndEvent:
			err = json.Unmarshal(e.Content, &endContent)

			if err != nil {
				t.Fatal(err.Error())
			}

			if !reflect.DeepEqual(endContent.FinalScores, expectedFinalScores) {
				t.Fatalf("expected %#v, got %#v instead", expectedFinalScores, endContent.FinalScores)
			}
		default:
			t.Fatalf("unexpected event type")
		}
	}

	reader.Close()
}

func TestReadLogsWhileStillWriting(t *testing.T) {
	filename := "test_file.jsonl"
	player1 := player.New(1)
	player2 := player.New(2)

	log, err := NewFromFile(filename)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer os.Remove(filename)

	deck := getTestDeck()
	expectedStartingTile := deck.StartingTile
	expectedStack := deck.GetRemaining()
	expectedPlayerCount := 2

	err = log.LogEvent(StartEvent, NewStartEntryContent(expectedStartingTile, expectedStack, expectedPlayerCount))
	if err != nil {
		t.Fatal(err.Error())
	}

	expectedTile := test.GetTestPlacedTile()
	err = log.LogEvent(PlaceTileEvent, NewPlaceTileEntryContent(player1.ID(), expectedTile))
	if err != nil {
		t.Fatal(err.Error())
	}

	var startContent StartEntryContent
	var placeTileContent PlaceTileEntryContent
	var endContent EndEntryContent

	reader, err := NewReaderFromFile(filename)

	if err != nil {
		t.Fatal(err.Error())
	}

	channel := reader.ReadLogs()

	entry := <-channel
	if entry.Event != StartEvent {
		t.Fatalf("expected %#v, got %#v instead", StartEvent, entry.Event)
	}
	err = json.Unmarshal(entry.Content, &startContent)
	if err != nil {
		t.Fatal(err.Error())
	}
	if !reflect.DeepEqual(startContent.StartingTile, expectedStartingTile) {
		t.Fatalf("expected %#v, got %#v instead", expectedStartingTile, startContent.StartingTile)
	}
	if !reflect.DeepEqual(startContent.Stack, expectedStack) {
		t.Fatalf("expected %#v, got %#v instead", expectedStack, startContent.Stack)
	}
	if !reflect.DeepEqual(startContent.PlayerCount, expectedPlayerCount) {
		t.Fatalf("expected %#v, got %#v instead", expectedPlayerCount, startContent.PlayerCount)
	}

	entry = <-channel
	if entry.Event != PlaceTileEvent {
		t.Fatalf("expected %#v, got %#v instead", PlaceTileEvent, entry.Event)
	}
	err = json.Unmarshal(entry.Content, &placeTileContent)
	if err != nil {
		t.Fatal(err.Error())
	}
	if placeTileContent.PlayerID != player1.ID() {
		t.Fatalf("expected %#v, got %#v instead", player1.ID(), placeTileContent.PlayerID)
	}
	if !reflect.DeepEqual(placeTileContent.Move, expectedTile) {
		t.Fatalf("expected %#v, got %#v instead", expectedTile, placeTileContent.Move)
	}

	expectedFinalScores := make(map[elements.ID]uint32, 0)
	expectedFinalScores[player1.ID()] = 1
	expectedFinalScores[player2.ID()] = 2
	err = log.LogEvent(EndEvent, NewEndEntryContent(expectedFinalScores))
	if err != nil {
		t.Fatal(err.Error())
	}

	log.Close()

	entry = <-channel
	if entry.Event != EndEvent {
		t.Fatalf("expected %#v, got %#v instead", EndEvent, entry.Event)
	}
	err = json.Unmarshal(entry.Content, &endContent)

	if err != nil {
		t.Fatal(err.Error())
	}

	if !reflect.DeepEqual(endContent.FinalScores, expectedFinalScores) {
		t.Fatalf("expected %#v, got %#v instead", expectedFinalScores, endContent.FinalScores)
	}

	reader.Close()
}

func TestFileLoggerInvalidFiles(t *testing.T) {
	filename := "test_file.jsonl"

	log, err := NewFromFile(filename)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer os.Remove(filename)

	err = log.Close()
	if err != nil {
		t.Fatal(err.Error())
	}

	err = log.Close()
	if err == nil {
		t.Fatal("FAILED")
	}

	deck := getTestDeck()
	err = log.LogEvent(StartEvent, NewStartEntryContent(deck.StartingTile, deck.Stack.GetTiles(), 2))
	if err == nil {
		t.Fatal("FAILED")
	}
}

//nolint:gocyclo// Cyclomatic complexity is not a problem in case of these tests
func TestLogger(t *testing.T) {
	buffer := bytes.NewBuffer(nil)

	log := New(buffer)

	deck := getTestDeck()
	expectedStack := deck.GetRemaining()
	expectedStartingTile := deck.StartingTile
	expectedPlayerCount := 2
	err := log.LogEvent(StartEvent, NewStartEntryContent(expectedStartingTile, expectedStack, expectedPlayerCount))
	if err != nil {
		t.Fatal(err.Error())
	}
	playerID := player.New(1)
	expectedTile := test.GetTestPlacedTile()
	err = log.LogEvent(PlaceTileEvent, NewPlaceTileEntryContent(playerID.ID(), expectedTile))
	if err != nil {
		t.Fatal(err.Error())
	}

	expectedScores := elements.NewScoreReport()
	expectedScores.ReceivedPoints[playerID.ID()] = 1
	expectedScores.ReceivedPoints[elements.ID(2)] = 2
	err = log.LogEvent(ScoreEvent, NewScoreEntryContent(expectedScores))
	if err != nil {
		t.Fatal(err.Error())
	}

	line, err := buffer.ReadString(byte('\n'))
	if err != nil {
		t.Fatal(err.Error())
	}

	var entryLine Entry
	var startContent StartEntryContent
	var placeTileContent PlaceTileEntryContent
	var endContent ScoreEntryContent

	err = json.Unmarshal([]byte(line), &entryLine)
	if err != nil {
		t.Fatal(err.Error())
	}
	if entryLine.Event != StartEvent {
		t.Fatalf("expected %#v, got %#v instead", StartEvent, entryLine.Event)
	}
	err = json.Unmarshal(entryLine.Content, &startContent)
	if err != nil {
		t.Fatal(err.Error())
	}
	if !reflect.DeepEqual(startContent.StartingTile, expectedStartingTile) {
		t.Fatalf("expected %#v, got %#v instead", expectedStartingTile, startContent.StartingTile)
	}
	if !reflect.DeepEqual(startContent.Stack, expectedStack) {
		t.Fatalf("expected %#v, got %#v instead", expectedStack, startContent.Stack)
	}
	if !reflect.DeepEqual(startContent.PlayerCount, expectedPlayerCount) {
		t.Fatalf("expected %#v, got %#v instead", expectedPlayerCount, startContent.PlayerCount)
	}

	line, err = buffer.ReadString(byte('\n'))
	if err != nil {
		t.Fatal(err.Error())
	}
	err = json.Unmarshal([]byte(line), &entryLine)
	if err != nil {
		t.Fatal(err.Error())
	}
	if entryLine.Event != PlaceTileEvent {
		t.Fatalf("expected %#v, got %#v instead", PlaceTileEvent, entryLine.Event)
	}
	err = json.Unmarshal(entryLine.Content, &placeTileContent)
	if err != nil {
		t.Fatal(err.Error())
	}
	if placeTileContent.PlayerID != playerID.ID() {
		t.Fatalf("expected %#v, got %#v instead", playerID.ID(), placeTileContent.PlayerID)
	}
	if !reflect.DeepEqual(placeTileContent.Move, expectedTile) {
		t.Fatalf("expected %#v, got %#v instead", expectedTile, placeTileContent.Move)
	}

	line, err = buffer.ReadString(byte('\n'))
	if err != nil {
		t.Fatal(err.Error())
	}
	err = json.Unmarshal([]byte(line), &entryLine)
	if err != nil {
		t.Fatal(err.Error())
	}
	if entryLine.Event != ScoreEvent {
		t.Fatalf("expected %#v, got %#v instead", ScoreEvent, entryLine.Event)
	}
	err = json.Unmarshal(entryLine.Content, &endContent)
	if err != nil {
		t.Fatal(err.Error())
	}
	if !reflect.DeepEqual(endContent.Scores, expectedScores) {
		t.Fatalf("expected %#v, got %#v instead", expectedScores, endContent.Scores)
	}
}

//nolint:gocyclo// Cyclomatic complexity is not a problem in case of these tests
func TestParseEntries(t *testing.T) {
	buffer := bytes.NewBuffer(nil)

	log := New(buffer)

	deck := getTestDeck()
	expectedStack := deck.GetRemaining()
	expectedStartingTile := deck.StartingTile
	expectedPlayerCount := 2
	err := log.LogEvent(StartEvent, NewStartEntryContent(expectedStartingTile, expectedStack, expectedPlayerCount))
	if err != nil {
		t.Fatal(err.Error())
	}
	playerID := player.New(1)
	expectedTile := test.GetTestPlacedTile()
	err = log.LogEvent(PlaceTileEvent, NewPlaceTileEntryContent(playerID.ID(), expectedTile))
	if err != nil {
		t.Fatal(err.Error())
	}

	expectedScores := elements.NewScoreReport()
	expectedScores.ReceivedPoints[playerID.ID()] = 1
	expectedScores.ReceivedPoints[elements.ID(2)] = 2
	err = log.LogEvent(ScoreEvent, NewScoreEntryContent(expectedScores))
	if err != nil {
		t.Fatal(err.Error())
	}

	line, err := buffer.ReadString(byte('\n'))
	if err != nil {
		t.Fatal(err.Error())
	}

	var entryLine Entry

	err = json.Unmarshal([]byte(line), &entryLine)
	if err != nil {
		t.Fatal(err.Error())
	}
	if entryLine.Event != StartEvent {
		t.Fatalf("expected %#v, got %#v instead", StartEvent, entryLine.Event)
	}
	startContent := ParseStartEntryContent(entryLine.Content)
	if err != nil {
		t.Fatal(err.Error())
	}
	if !reflect.DeepEqual(startContent.StartingTile, expectedStartingTile) {
		t.Fatalf("expected %#v, got %#v instead", expectedStartingTile, startContent.StartingTile)
	}
	if !reflect.DeepEqual(startContent.Stack, expectedStack) {
		t.Fatalf("expected %#v, got %#v instead", expectedStack, startContent.Stack)
	}
	if !reflect.DeepEqual(startContent.PlayerCount, expectedPlayerCount) {
		t.Fatalf("expected %#v, got %#v instead", expectedPlayerCount, startContent.PlayerCount)
	}

	line, err = buffer.ReadString(byte('\n'))
	if err != nil {
		t.Fatal(err.Error())
	}
	err = json.Unmarshal([]byte(line), &entryLine)
	if err != nil {
		t.Fatal(err.Error())
	}
	if entryLine.Event != PlaceTileEvent {
		t.Fatalf("expected %#v, got %#v instead", PlaceTileEvent, entryLine.Event)
	}
	placeTileContent := ParsePlaceTileEntryContent(entryLine.Content)
	if err != nil {
		t.Fatal(err.Error())
	}
	if placeTileContent.PlayerID != playerID.ID() {
		t.Fatalf("expected %#v, got %#v instead", playerID.ID(), placeTileContent.PlayerID)
	}
	if !reflect.DeepEqual(placeTileContent.Move, expectedTile) {
		t.Fatalf("expected %#v, got %#v instead", expectedTile, placeTileContent.Move)
	}

	line, err = buffer.ReadString(byte('\n'))
	if err != nil {
		t.Fatal(err.Error())
	}
	err = json.Unmarshal([]byte(line), &entryLine)
	if err != nil {
		t.Fatal(err.Error())
	}
	if entryLine.Event != ScoreEvent {
		t.Fatalf("expected %#v, got %#v instead", ScoreEvent, entryLine.Event)
	}
	scoreContent := ParseScoreEntryContent(entryLine.Content)
	if err != nil {
		t.Fatal(err.Error())
	}
	if !reflect.DeepEqual(scoreContent.Scores, expectedScores) {
		t.Fatalf("expected %#v, got %#v instead", expectedScores, scoreContent.Scores)
	}
}
