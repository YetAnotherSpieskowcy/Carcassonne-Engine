package logger

import (
	"encoding/binary"
	"errors"
	"io"

	pb "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/proto" //nolint:all
	"google.golang.org/protobuf/proto"
)

var ErrCopyToNotImplemented = errors.New("the type does not implement the CopyTo() method")

type Logger interface {
	AsWriter() io.Writer // only meant to be called by other (i.e. public) methods
	LogEvent(pb.Entry) error
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

func (logger *BaseLogger) LogEvent(entry pb.Entry) error {
	out, err := proto.Marshal(&entry)
	if err != nil {
		return err
	}

	// write dize of entry to the file
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, uint32(len(out)))
	if _, err := logger.writer.Write(buf); err != nil {
		return err
	}

	// write message
	if _, err := logger.writer.Write(out); err != nil {
		return err
	}

	return nil
}

func (logger *BaseLogger) CopyTo(dst Logger) error {
	_ = dst
	return ErrCopyToNotImplemented
}

type EmptyLogger struct{}

func NewEmpty() EmptyLogger {
	return EmptyLogger{}
}

func (*EmptyLogger) AsWriter() io.Writer {
	return nil
}

func (*EmptyLogger) LogEvent(entry pb.Entry) error { //nolint:revive // causes gopy to fail
	return nil
}

func (*EmptyLogger) CopyTo(dst Logger) error { //nolint:revive // causes gopy to fail
	return nil
}
