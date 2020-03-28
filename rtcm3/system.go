package rtcm3

import (
	"encoding/binary"
	"github.com/go-restruct/restruct"
)

// System Parameters
type Message1013 struct {
	AbstractMessage
	ReferenceStationId uint16 `struct:"uint16:12"`
	Mjd                uint16 `struct:"uint16"`
	SecondsOfDay       uint32 `struct:"uint32:17"`
	MessageCount       uint8  `struct:"uint8:5,sizeof=Messages"`
	LeapSeconds        uint8  `struct:"uint8"`
	Messages           []struct {
		Id                   uint16 `struct:"uint16:12"`
		SyncFlag             bool   `struct:"uint8:1,variantbool"`
		TransmissionInterval uint16 `struct:"uint16"`
	}
}

func DeserializeMessage1013(data []byte) (msg Message1013) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1013) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

// Unicode Text String
type Message1029 struct {
	AbstractMessage
	ReferenceStationId uint16 `struct:"uint16:12"`
	Mjd                uint16 `struct:"uint16"`
	SecondsOfDay       uint32 `struct:"uint32:17"`
	Characters         uint8  `struct:"uint8:7"`
	CodeUnitsLength    uint8  `struct:"uint8"`
	CodeUnits          string `struct:"[]byte,sizefrom=CodeUnitsLength"`
}

func DeserializeMessage1029(data []byte) (msg Message1029) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1029) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}
