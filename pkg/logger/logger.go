package logger

import (
	"encoding/json"
	"io"
)

type Logger interface {
	LogEvent(EventType, interface{}) error
}

type BaseLogger struct {
	writer io.Writer
}

func New(writer io.Writer) BaseLogger {
	return BaseLogger{writer}
}

func (logger *BaseLogger) LogEvent(eventName EventType, event interface{}) error {
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
