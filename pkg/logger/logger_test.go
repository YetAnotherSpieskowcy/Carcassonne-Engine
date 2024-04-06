package logger

import (
	"bufio"
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

func TestLog(t *testing.T) {
	filename := "test_file.jsonl"
	log := Logger{}

	err := log.Open(filename)
	if err != nil {
		t.Fatal("FAILED")
	}

	defer os.Remove(filename)
	defer log.Close()

	err = log.Start([]int{1, 2, 3}, []string{"Player1", "Player2"})
	if err != nil {
		t.Fatal("FAILED")
	}

	err = log.PlaceTile(0, 1, []int{1, 2}, 0)
	if err != nil {
		t.Fatal("FAILED")
	}

	err = log.End([]int{1, 2})
	if err != nil {
		t.Fatal("FAILED")
	}

	var startLine struct {
		Event   string
		Deck    []int
		Players []string
	}

	var placeTileLine struct {
		Event    string
		Player   int
		Rotation int
		Position []int
		Meeple   int
	}

	var endLine struct {
		Event  string
		Scores []int
	}

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
