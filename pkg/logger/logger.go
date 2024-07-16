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

func (logger *Logger) LogEvent(eventName string, event interface{}) error {
	jsonData, err := json.Marshal(event)
	if err != nil {
		return err
	}

	entry := NewEntry(eventName, jsonData)
	jsonEntry, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	_, err = logger.writer.Write(jsonEntry)
	if err != nil {
		return err
	}
	_, err = logger.writer.Write([]byte("\n"))
	if err != nil {
		return err
	}

	return nil
}
