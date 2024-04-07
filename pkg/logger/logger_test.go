package logger

import (
	"bufio"
	"bytes"
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

func TestFileLogger(t *testing.T) {
	filename := "test_file.jsonl"

	log, err := NewFromFile(filename)
	if err != nil {
		t.Fatal("FAILED")
	}
	defer os.Remove(filename)

	if err != nil {
		t.Fatal("FAILED")
	}

	err = log.LogEvent(NewStartEntry([]int{1, 2, 3}, []string{"Player1", "Player2"}))
	if err != nil {
		t.Fatal("FAILED")
	}

	err = log.LogEvent(NewPlaceTileEntry(0, 1, []int{1, 2}, 0))
	if err != nil {
		t.Fatal("FAILED")
	}

	err = log.LogEvent(NewEndEntry([]int{1, 2}))
	if err != nil {
		t.Fatal("FAILED")
	}

	var startLine StartEntry
	var placeTileLine PlaceTileEntry
	var endLine EndEntry

	log.Close()

	file, err := os.Open(filename)
	if err != nil {
		t.Fatal("FAILED")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	err = json.Unmarshal([]byte(scanner.Text()), &startLine)
	if err != nil {
		t.Fatal("FAILED")
	}
	if startLine.Event != "start" {
		t.Fatal("FAILED")
	}
	if !reflect.DeepEqual(startLine.Deck, []int{1, 2, 3}) {
		t.Fatal("FAILED")
	}
	if !reflect.DeepEqual(startLine.Players, []string{"Player1", "Player2"}) {
		t.Fatal("FAILED")
	}

	scanner.Scan()
	err = json.Unmarshal([]byte(scanner.Text()), &placeTileLine)
	if err != nil {
		t.Fatal("FAILED")
	}
	if placeTileLine.Event != "place" {
		t.Fatal("FAILED")
	}
	if placeTileLine.Rotation != 1 {
		t.Fatal("FAILED")
	}
	if !reflect.DeepEqual(placeTileLine.Position, []int{1, 2}) {
		t.Fatal("FAILED")
	}
	if placeTileLine.Meeple != 0 {
		t.Fatal("FAILED")
	}

	scanner.Scan()
	err = json.Unmarshal([]byte(scanner.Text()), &endLine)
	if err != nil {
		t.Fatal("FAILED")
	}
	if endLine.Event != "end" {
		t.Fatal("FAILED")
	}
	if !reflect.DeepEqual(endLine.Scores, []int{1, 2}) {
		t.Fatal("FAILED")
	}
}

func TestFileLoggerInvalidFiles(t *testing.T) {
	filename := "test_file.jsonl"

	log, err := NewFromFile(filename)
	if err != nil {
		t.Fatal("FAILED")
	}
	defer os.Remove(filename)

	err = log.Close()
	if err != nil {
		t.Fatal("FAILED")
	}

	err = log.Close()
	if err == nil {
		t.Fatal("FAILED")
	}

	err = log.LogEvent(NewStartEntry([]int{1, 2, 3}, []string{"Player1", "Player2"}))
	if err == nil {
		t.Fatal("FAILED")
	}
}

func TestLogger(t *testing.T) {
	buffer := bytes.NewBuffer(nil)

	log := New(buffer)

	err := log.LogEvent(NewStartEntry([]int{1, 2, 3}, []string{"Player1", "Player2"}))
	if err != nil {
		t.Fatal("FAILED")
	}

	err = log.LogEvent(NewPlaceTileEntry(0, 1, []int{1, 2}, 0))
	if err != nil {
		t.Fatal("FAILED")
	}

	err = log.LogEvent(NewEndEntry([]int{1, 2}))
	if err != nil {
		t.Fatal("FAILED")
	}

	line, err := buffer.ReadString(byte('\n'))
	if err != nil {
		t.Fatal("FAILED")
	}

	var startLine StartEntry
	var placeTileLine PlaceTileEntry
	var endLine EndEntry

	err = json.Unmarshal([]byte(line), &startLine)
	if err != nil {
		t.Fatal("FAILED")
	}
	if startLine.Event != "start" {
		t.Fatal("FAILED")
	}
	if !reflect.DeepEqual(startLine.Deck, []int{1, 2, 3}) {
		t.Fatal("FAILED")
	}
	if !reflect.DeepEqual(startLine.Players, []string{"Player1", "Player2"}) {
		t.Fatal("FAILED")
	}

	line, err = buffer.ReadString(byte('\n'))
	if err != nil {
		t.Fatal("FAILED")
	}
	err = json.Unmarshal([]byte(line), &placeTileLine)
	if err != nil {
		t.Fatal("FAILED")
	}
	if placeTileLine.Event != "place" {
		t.Fatal("FAILED")
	}
	if placeTileLine.Rotation != 1 {
		t.Fatal("FAILED")
	}
	if !reflect.DeepEqual(placeTileLine.Position, []int{1, 2}) {
		t.Fatal("FAILED")
	}
	if placeTileLine.Meeple != 0 {
		t.Fatal("FAILED")
	}

	line, err = buffer.ReadString(byte('\n'))
	if err != nil {
		t.Fatal("FAILED")
	}
	err = json.Unmarshal([]byte(line), &endLine)
	if err != nil {
		t.Fatal("FAILED")
	}
	if endLine.Event != "end" {
		t.Fatal("FAILED")
	}
	if !reflect.DeepEqual(endLine.Scores, []int{1, 2}) {
		t.Fatal("FAILED")
	}
}
