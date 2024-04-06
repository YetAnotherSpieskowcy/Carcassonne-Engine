package logger

import (
	"encoding/json"
	"os"
)

type Logger struct {
	file *os.File
}

func New(file os.File) Logger {
	return Logger{&file}
}

func (logger *Logger) logEvent(event map[string]interface{}) error {
	jsonData, err := json.Marshal(event)
	if err != nil {
		return err
	}

	jsonData = append(jsonData, byte('\n'))

	_, err = logger.file.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}

func (logger *Logger) Open(filename string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	logger.file = file
	return err
}

func (logger *Logger) Close() error {
	err := logger.file.Close()
	return err
}

func (logger *Logger) Start(deck []int, players []string) error { // todo deck type should be: []tiles.Tile
	err := logger.logEvent(
		map[string]interface{}{
			"event":   "start",
			"deck":    deck,
			"players": players,
		})
	return err
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
	return err
}

func (logger *Logger) End(scores []int) error {
	err := logger.logEvent(
		map[string]interface{}{
			"event":  "end",
			"scores": scores,
		})
	return err
}
