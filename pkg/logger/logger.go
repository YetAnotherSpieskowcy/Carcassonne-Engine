package logger

import (
	"encoding/json"
	"os"
)

type Logger struct {
	filename string
}

func New(filename string) Logger {
	return Logger{filename}
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

	return nil
}

func (logger *Logger) PlaceTile(player int, rotation int, position []int, meeple int) error { // todo meeple type should be connection.Side
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
	err := logger.logEvent(
		map[string]interface{}{
			"event":  "end",
			"scores": scores,
		})

	if err != nil {
		return err
	}

	return nil
}
