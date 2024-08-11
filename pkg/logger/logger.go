package logger

import (
	"encoding/json"
	"errors"
	"io"
)

var ErrCopyToNotImplemented = errors.New("the type does not implement the CopyTo() method")

type Logger interface {
	AsWriter() io.Writer // only meant to be called by other (i.e. public) methods
	LogEvent(EventType, interface{}) error
	CopyTo(Logger) error
}

type BaseLogger struct {
	writer io.Writer
}

func New(writer io.Writer) BaseLogger {
	return BaseLogger{writer}
}

func (logger *BaseLogger) AsWriter() io.Writer {
	return logger.writer
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

func (logger *BaseLogger) CopyTo(dst Logger) error {
	_ = dst
	return ErrCopyToNotImplemented
}
