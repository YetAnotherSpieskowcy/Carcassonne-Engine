package logger

import (
	"encoding/json"
	"io"
	"os"
	"reflect"
)

type LogReader struct {
	file                 *os.File
	decoder              *json.Decoder
	shouldRestartDecoder bool
	fileEndOffset        int64
}

func NewLogReader(filename string) (LogReader, error) {
	logReader := LogReader{}
	err := logReader.Open(filename)
	return logReader, err
}

func (reader *LogReader) Open(filename string) error {
	file, err := os.Open(filename)
	reader.file = file
	reader.decoder = json.NewDecoder(reader.file)
	reader.decoder.DisallowUnknownFields()
	return err
}

func (reader *LogReader) Close() error {
	err := reader.file.Close()
	return err
}

func (reader LogReader) ReadLogs() <-chan Entry {
	channel := make(chan Entry)

	go func() {
		var entry Entry
		decoder := json.NewDecoder(reader.file)
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

// Reads a single log entry. Returns the entry and a bool indicating whether or not it was correctly read.
func (reader *LogReader) ReadEntry() (Entry, bool) {
	if reader.shouldRestartDecoder {
		// reopen the decoder if file end was reached previously
		_, err := reader.file.Seek(reader.fileEndOffset, io.SeekStart)
		if err != nil {
			panic(err)
		}
		reader.decoder = json.NewDecoder(reader.file)
		reader.shouldRestartDecoder = false
	}

	value := reflect.ValueOf(reader.decoder).Elem()
	startOffset := value.FieldByName("scanned").Int()

	var entry Entry
	err := reader.decoder.Decode(&entry)

	if err == nil {
		return entry, true
	}

	reader.shouldRestartDecoder = true
	reader.fileEndOffset = startOffset
	/*
		silently skipping errors (instead of panicking) because they can be	returned
		when reading a partially written data.

		example situation where this may occur:
		- writer: {"entry": "entry con
		- reader: <error>
		- writer: tent"}
		- reader: successfully read {"entry": "entry content"}

		This isn't the cleanest solution, but I don't think there is any simple way
		to guard against such cases, since file write operations aren't guaranteed to be atomic
	*/
	return Entry{}, false
}
