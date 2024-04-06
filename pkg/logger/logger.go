package logger

import (
	"encoding/json"
	"fmt"
	"os"
)

type loggerState byte

const (
	new loggerState = iota
	started
	ended
)

type Logger struct {
	filename string
	state    loggerState
}

func New(filename string) Logger {
	return Logger{filename, new}
}

func (logger *Logger) logEvent(event map[string]interface{}) error {
	jsonData, err := json.Marshal(event)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(logger.filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(string(jsonData) + "\n")
	if err != nil {
		return err
	}

	return nil
}

func (logger *Logger) createFile() error {
	file, err := os.Create(logger.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}

func (logger *Logger) Start(deck []int, players []string) error { // todo deck type should be: []tiles.Tile
	if logger.state != new {
		return fmt.Errorf("logger already started")
	}

	err := logger.createFile()
	if err != nil {
		return err
	}

	err = logger.logEvent(
		map[string]interface{}{
			"event":   "start",
			"deck":    deck,
			"players": players,
		})
	if err != nil {
		return err
	}

	logger.state = started

	return nil
}

func (logger *Logger) PlaceTile(player int, rotation int, position []int, meeple int) error { // todo meeple type should be connection.Side
	if logger.state != started {
		return fmt.Errorf("logger already ended or not yet started")
	}

	err := logger.logEvent(
		map[string]interface{}{
			"event":    "place",
			"player":   player,
			"rotation": rotation,
			"position": position,
			"meeple":   meeple,
		})
	if err != nil {
		return err
	}

	return nil
}

func (logger *Logger) End(scores []int) error {
	if logger.state != started {
		return fmt.Errorf("logger already ended or not yet started")
	}

	err := logger.logEvent(
		map[string]interface{}{
			"event":  "end",
			"scores": scores,
		})

	if err != nil {
		return err
	}

	logger.state = ended

	return nil
}
