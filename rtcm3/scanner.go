package rtcm3

import (
	"bufio"
	"io"
)

type Scanner struct {
	Reader *bufio.Reader
}

// NewScanner returns a Scanner for the given io.Reader
func NewScanner(r io.Reader) Scanner {
	return Scanner{bufio.NewReader(r)}
}

// Next reads from IO until a Message is found
func (scanner Scanner) Next() (message Message, err error) {
	frame, err := scanner.NextFrame()
	if err != nil {
		return nil, err
	}
	return DeserializeMessage(frame.Payload), err // probably have DeserializeMessage return err
}

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
