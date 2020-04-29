package rtcm3

import (
	"bufio"
	"io"
)

// Scanner is used for parsing RTCM3 Messages from a Reader
type Scanner struct {
	Reader *bufio.Reader
}

func NewScanner(r io.Reader) Scanner {
	return Scanner{bufio.NewReader(r)}
}

// NextMessage reads from Scanner until a Message is found
func (scanner Scanner) NextMessage() (message Message, err error) {
	frame, err := scanner.NextFrame()
	if err != nil {
		return nil, err
	}
	return DeserializeMessage(frame.Payload), err // DeserializeMessage should return err
}

// NextFrame reads from Scanner until a valid Frame is found
func (scanner Scanner) NextFrame() (frame Frame, err error) {
	for {
		frame, err := DeserializeFrame(scanner.Reader)
		if err != nil {
			if err.Error() == "invalid preamble" || err.Error() == "invalid CRC" {
				// Continue reading from next byte if a valid Frame was not found
				// TODO: return byte array of skipped bytes
				continue
			}
		}
		return frame, err
	}
}
