package logger

import (
	"encoding/json"
	"io"
	"os"
	"time"
)

type FileReader struct {
	file *os.File
}

func NewReaderFromFile(filename string) (FileReader, error) {
	fileReader := FileReader{}
	err := fileReader.Open(filename)
	return fileReader, err
}

func (fl *FileReader) Open(filename string) error {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	fl.file = file
	return err
}

func (fl *FileReader) Close() error {
	err := fl.file.Close()
	return err
}

func (fl FileReader) ReadLogs() <-chan Entry {
	channel := make(chan Entry)

	go func() {
		var entry Entry
		decoder := json.NewDecoder(fl.file)
		decoder.DisallowUnknownFields()
		for {
			err := decoder.Decode(&entry)
			if err == io.EOF {
				time.Sleep(1000000) // = 1ms
			} else if err != nil {
				panic(err)
			} else {
				channel <- entry
			}

			if entry.Event == EndEvent {
				break
			}
		}
		close(channel)
	}()

	return channel
}
