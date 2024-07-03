package logger

import "os"

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
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	fl.file = file
	fl.writer = file
	return err
}

func (fl *FileLogger) Close() error {
	err := fl.file.Close()
	return err
}
