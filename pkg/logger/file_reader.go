package logger

import (
	"encoding/binary"
	"io"
	"os"
	"time"

	pb "github.com/YetAnotherSpieskowcy/Carcassonne-Engine/pkg/proto"
	"google.golang.org/protobuf/proto"
)

type FileReader struct {
	file *os.File
}

func NewReaderFromFile(filename string) (FileReader, error) {
	fileReader := FileReader{}
	err := fileReader.Open(filename)
	return fileReader, err
}

func (fr *FileReader) Open(filename string) error {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	fr.file = file
	return err
}

func (fr *FileReader) Close() error {
	err := fr.file.Close()
	return err
}

func (fr FileReader) ReadLogs() <-chan *pb.Entry {
	channel := make(chan *pb.Entry)

	go func() {
		var offset int64 = 0
		for {
			entry := &pb.Entry{}
			buf := make([]byte, 4)

			_, err := fr.file.ReadAt(buf, offset)
			if err == io.EOF {
				break
			} else {

				itemSize := binary.LittleEndian.Uint32(buf)
				offset += 4

				// reading the actual encoded item
				item := make([]byte, itemSize)
				_, err = fr.file.ReadAt(item, offset)

				if err == io.EOF {
					time.Sleep(time.Millisecond)
				} else if err != nil {
					panic(err)
				} else {
					err = proto.Unmarshal(item, entry)
					if err != nil {
						panic(err)
					}
					channel <- entry
					offset += int64(itemSize)
				}

				if entry.Event == pb.EventType_EVENT_TYPE_FINAL_SCORE_EVENT {
					break
				}
			}
		}
		close(channel)
	}()

	return channel
}
