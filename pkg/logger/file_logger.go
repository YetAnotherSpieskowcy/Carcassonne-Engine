package logger

import "os"

type FileLogger struct {
	logger Logger
	file   *os.File
}

func NewFromFile(filename string) (FileLogger, error) {
	fileLogger := FileLogger{}
	err := fileLogger.Open(filename)
	return fileLogger, err
}

func (fl *FileLogger) Open(filename string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	fl.file = file
	fl.logger = Logger{file}
	return err
}

func (fl *FileLogger) Close() error {
	err := fl.file.Close()
	return err
}

func (fl *FileLogger) Start(deck []int, players []string) error {
	return fl.logger.Start(deck, players)
}

func (fl *FileLogger) PlaceTile(player int, rotation int, position []int, meeple int) error {
	return fl.logger.PlaceTile(player, rotation, position, meeple)
}

func (fl *FileLogger) End(scores []int) error {
	return fl.logger.End(scores)
}
