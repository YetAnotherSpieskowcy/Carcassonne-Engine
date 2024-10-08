package logger

import (
	"encoding/json"
	"io"
	"os"
)

type FileLogger struct {
	BaseLogger
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

func (fl *FileLogger) CopyTo(dst Logger) error {
	currentOffset, err := fl.file.Seek(0, io.SeekCurrent)
	if err != nil {
		return err
	}

	if _, err = fl.file.Seek(0, io.SeekStart); err != nil {
		return err
	}

	writer := dst.AsWriter()
	if _, err = io.CopyN(writer, fl.file, currentOffset); err != nil {
		return err
	}

	_, err = fl.file.Seek(currentOffset, io.SeekStart)

	return err
}
