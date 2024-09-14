package logger

import (
	"encoding/json"
	"io"
	"os"
)

type LogReader struct {
	file           *os.File
	decoder        *json.Decoder
	fileEndReached bool
	fileEndOffset  int64
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

// Reads a single log entry. Returns the entry and a bool indicating whether or not it was correctly read.
func (reader *LogReader) ReadEntry() (Entry, bool) {
	if reader.fileEndReached {
		// reopen the decoder if file end was reached previously
		_, err := reader.file.Seek(reader.fileEndOffset, io.SeekStart)
		if err != nil {
			panic(err)
		}
		reader.decoder = json.NewDecoder(reader.file)
		reader.fileEndReached = false
	}

	var entry Entry

	err := reader.decoder.Decode(&entry)

	if err == nil {
		return entry, true
	} else {
		if err == io.EOF {
			reader.fileEndReached = true
			endOffset, err := reader.file.Seek(0, io.SeekCurrent)
			if err != nil {
				panic(err)
			}
			reader.fileEndOffset = endOffset

			return Entry{}, false

		} else if _, ok := err.(*json.SyntaxError); ok {
			return Entry{}, false
			/*
				silently skipping json syntax errors (instead of panicking) because these can be returned when reading a partially written data.
				example situation where this may occur:
				- writer: {"entry": "entry con
				- reader: <syntax error>
				- writer: tent"}
				- reader: successfully read {"entry": "entry content"}

				This isn't the cleanest solution, but I don't think there is any simple way to guard against such cases
			*/

		}
		// else:
		panic(err)
	}
}
