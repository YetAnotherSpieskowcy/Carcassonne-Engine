package logger

import (
	"encoding/json"
	"io"
)

type Logger struct {
	writer io.Writer
}

func New(writer io.Writer) Logger {
	return Logger{writer}
}

func (logger *Logger) logEvent(event map[string]interface{}) error {
	jsonData, err := json.Marshal(event)
	if err != nil {
		return err
	}

	_, err = logger.writer.Write(jsonData)
	if err != nil {
		return err
	}
	_, err = logger.writer.Write([]byte("\n"))
	if err != nil {
		return err
	}

	return nil
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
