package logger

import (
	"encoding/json"
	"io"
	"os"
)

type FileLogger struct {
	Logger
	file *os.File
}

func NewFromFile(filename string) (FileLogger, error) {
	fileLogger := FileLogger{}
	err := fileLogger.Open(filename)
	return fileLogger, err
}

func (fl *FileLogger) Open(filename string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	fl.file = file
	fl.writer = file
	if err != nil {
		return err
	}
	return err
}

func (fl *FileLogger) Close() error {
	err := fl.file.Close()
	return err
}

func (fl FileLogger) ReadLogs() <-chan Entry {
	channel := make(chan Entry)

	go func() {
		var entry Entry
		decoder := json.NewDecoder(fl.file)
		decoder.DisallowUnknownFields()
		for {
			err := decoder.Decode(&entry)
			if err == io.EOF {
				break
			} else if err != nil {
				panic(err)
			}
			channel <- entry
		}
		close(channel)
	}()

	return channel
}
