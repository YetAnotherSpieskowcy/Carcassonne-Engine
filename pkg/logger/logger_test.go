package logger

import (
	"bufio"
	"bytes"
	"encoding/json"
	"os"
	"reflect"
	"testing"

	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/deck"
	"github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/game/test"
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
	err = log.LogEvent(NewStartEntry(deck, expectedPlayerCount))
	if err != nil {
		t.Fatal(err.Error())
	}

	expectedTile := test.GetTestPlacedTile()
	err = log.LogEvent(NewPlaceTileEntry(expectedTile.Player, expectedTile.LegalMove))
	if err != nil {
		t.Fatal(err.Error())
	}

	expectedScores := []uint32{1, 2}
	err = log.LogEvent(NewEndEntry(expectedScores))
	if err != nil {
		t.Fatal(err.Error())
	}

	var startLine StartEntry
	var placeTileLine PlaceTileEntry
	var endLine EndEntry

	log.Close()

	file, err := os.Open(filename)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	err = json.Unmarshal([]byte(scanner.Text()), &startLine)
	if err != nil {
		t.Fatal(err.Error())
	}
	if startLine.Event != "start" {
		t.Fatalf("expected %#v, got %#v instead", "start", startLine.Event)
	}
	if !reflect.DeepEqual(startLine.StartingTile, expectedStartingTile) {
		t.Fatalf("expected %#v, got %#v instead", expectedStartingTile, startLine.StartingTile)
	}
	if !reflect.DeepEqual(startLine.Stack, expectedStack) {
		t.Fatalf("expected %#v, got %#v instead", expectedStack, startLine.Stack)
	}
	if !reflect.DeepEqual(startLine.PlayerCount, expectedPlayerCount) {
		t.Fatalf("expected %#v, got %#v instead", expectedPlayerCount, startLine.PlayerCount)
	}

	scanner.Scan()
	err = json.Unmarshal([]byte(scanner.Text()), &placeTileLine)
	if err != nil {
		t.Fatal(err.Error())
	}
	if placeTileLine.Event != "place" {
		t.Fatalf("expected %#v, got %#v instead", "place", placeTileLine.Event)
	}
	if placeTileLine.PlayerID != expectedTile.Player.ID() {
		t.Fatalf("expected %#v, got %#v instead", expectedTile.Player.ID(), placeTileLine.PlayerID)
	}
	if !reflect.DeepEqual(placeTileLine.Move, expectedTile.LegalMove) {
		t.Fatalf("expected %#v, got %#v instead", expectedTile.LegalMove, placeTileLine.Move)
	}

	scanner.Scan()
	err = json.Unmarshal([]byte(scanner.Text()), &endLine)
	if err != nil {
		t.Fatal(err.Error())
	}
	if endLine.Event != "end" {
		t.Fatalf("expected %#v, got %#v instead", "end", endLine.Event)
	}
	if !reflect.DeepEqual(endLine.Scores, expectedScores) {
		t.Fatalf("expected %#v, got %#v instead", expectedScores, endLine.Scores)
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

	err = log.LogEvent(NewStartEntry(getTestDeck(), 2))
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
	err := log.LogEvent(NewStartEntry(deck, expectedPlayerCount))
	if err != nil {
		t.Fatal(err.Error())
	}

	expectedTile := test.GetTestPlacedTile()
	err = log.LogEvent(NewPlaceTileEntry(expectedTile.Player, expectedTile.LegalMove))
	if err != nil {
		t.Fatal(err.Error())
	}

	expectedScores := []uint32{1, 2}
	err = log.LogEvent(NewEndEntry(expectedScores))
	if err != nil {
		t.Fatal(err.Error())
	}

	line, err := buffer.ReadString(byte('\n'))
	if err != nil {
		t.Fatal(err.Error())
	}

	var startLine StartEntry
	var placeTileLine PlaceTileEntry
	var endLine EndEntry

	err = json.Unmarshal([]byte(line), &startLine)
	if err != nil {
		t.Fatal(err.Error())
	}
	if startLine.Event != "start" {
		t.Fatalf("expected %#v, got %#v instead", "start", startLine.Event)
	}
	if !reflect.DeepEqual(startLine.StartingTile, expectedStartingTile) {
		t.Fatalf("expected %#v, got %#v instead", expectedStartingTile, startLine.StartingTile)
	}
	if !reflect.DeepEqual(startLine.Stack, expectedStack) {
		t.Fatalf("expected %#v, got %#v instead", expectedStack, startLine.Stack)
	}
	if !reflect.DeepEqual(startLine.PlayerCount, expectedPlayerCount) {
		t.Fatalf("expected %#v, got %#v instead", expectedPlayerCount, startLine.PlayerCount)
	}

	line, err = buffer.ReadString(byte('\n'))
	if err != nil {
		t.Fatal(err.Error())
	}
	err = json.Unmarshal([]byte(line), &placeTileLine)
	if err != nil {
		t.Fatal(err.Error())
	}
	if placeTileLine.PlayerID != expectedTile.Player.ID() {
		t.Fatalf("expected %#v, got %#v instead", expectedTile.Player.ID(), placeTileLine.PlayerID)
	}
	if !reflect.DeepEqual(placeTileLine.Move, expectedTile.LegalMove) {
		t.Fatalf("expected %#v, got %#v instead", expectedTile.LegalMove, placeTileLine.Move)
	}

	line, err = buffer.ReadString(byte('\n'))
	if err != nil {
		t.Fatal(err.Error())
	}
	err = json.Unmarshal([]byte(line), &endLine)
	if err != nil {
		t.Fatal(err.Error())
	}
	if endLine.Event != "end" {
		t.Fatalf("expected %#v, got %#v instead", "end", endLine.Event)
	}
	if !reflect.DeepEqual(endLine.Scores, expectedScores) {
		t.Fatalf("expected %#v, got %#v instead", expectedScores, endLine.Scores)
	}
}
