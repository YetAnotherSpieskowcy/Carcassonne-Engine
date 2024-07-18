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

	log, err := NewFromFile(filename)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer os.Remove(filename)

	if err != nil {
		t.Fatal(err.Error())
	}

	deck := getTestDeck()
	expectedStartingTile := deck.StartingTile
	expectedStack := deck.GetRemaining()
	expectedPlayerCount := 2

	err = log.LogEvent(StartEvent, NewStartEntryContent(expectedStartingTile, expectedStack, expectedPlayerCount))
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

	var entryLine Entry
	var startContent StartEntryContent
	var placeTileContent PlaceTileEntryContent
	var endContent ScoreEntryContent

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
	if placeTileContent.PlayerID != playerID.ID() {
		t.Fatalf("expected %#v, got %#v instead", playerID.ID(), placeTileContent.PlayerID)
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
	err = json.Unmarshal(entryLine.Content, &endContent)
	if err != nil {
		t.Fatal(err.Error())
	}
	if !reflect.DeepEqual(endContent.Scores, expectedScores) {
		t.Fatalf("expected %#v, got %#v instead", expectedScores, endContent.Scores)
	}
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
