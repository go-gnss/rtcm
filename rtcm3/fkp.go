package rtcm3

import (
	"encoding/binary"
	"github.com/go-restruct/restruct"
)

// GPS Network FKP Gradient
type Message1034 struct {
	AbstractMessage
	ReferenceStationID        uint16 `struct:"uint16:12"`
	FKPEpoch                  uint32 `struct:"uint32:20"`
	SatelliteSignalsProcessed uint8  `struct:"uint8:5,sizeof=SatelliteData"`
	SatelliteData             []struct {
		SatelliteID uint8 `struct:"uint8:6"`
		IODE        uint8 `struct:"uint8"`
		N0          int16 `struct:"int16:12"`
		E0          int16 `struct:"int16:12"`
		NI          int16 `struct:"int16:14"`
		EI          int16 `struct:"int16:14"`
	}
}

func DeserializeMessage1034(data []byte) (msg Message1034) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1034) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

// GLONASS Network FKP Gradient
type Message1035 struct {
	AbstractMessage
	ReferenceStationID        uint16 `struct:"uint16:12"`
	FKPEpoch                  uint32 `struct:"uint32:17"`
	SatelliteSignalsProcessed uint8  `struct:"uint8:5,sizeof=SatelliteData"`
	SatelliteData             []struct {
		SatelliteID uint8 `struct:"uint8:6"`
		IOD         uint8 `struct:"uint8"`
		N0          int16 `struct:"int16:12"`
		E0          int16 `struct:"int16:12"`
		NI          int16 `struct:"int16:14"`
		EI          int16 `struct:"int16:14"`
	}
}

func DeserializeMessage1035(data []byte) (msg Message1035) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1035) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}
