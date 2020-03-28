package rtcm3

import (
	"encoding/binary"
	"github.com/go-restruct/restruct"
)

type AntennaReferencePoint struct {
	ReferenceStationId        uint16 `struct:"uint16:12"`
	ItrfRealizationYear       uint8  `struct:"uint8:6"`
	GpsIndicator              bool   `struct:"uint8:1,variantbool"`
	GlonassIndicator          bool   `struct:"uint8:1,variantbool"`
	GalileoIndicator          bool   `struct:"uint8:1,variantbool"`
	ReferenceStationIndicator bool   `struct:"uint8:1,variantbool"`
	ReferencePointX           int64  `struct:"int64:38"`
	SingleReceiverOscilator   bool   `struct:"uint8:1,variantbool"`
	Reserved                  bool   `struct:"uint8:1,variantbool"`
	ReferencePointY           int64  `struct:"int64:38"`
	QuarterCycleIndicator     uint8  `struct:"uint8:2"`
	ReferencePointZ           int64  `struct:"int64:38"`
}

// Stationary RTK Reference Station ARP
type Message1005 struct {
	AbstractMessage
	AntennaReferencePoint
}

func DeserializeMessage1005(data []byte) (msg Message1005) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1005) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

// Stationary RTK Reference Station ARP with Antenna Height
type Message1006 struct {
	AbstractMessage
	AntennaReferencePoint
	AntennaHeight uint16 `struct:"uint16"`
}

func DeserializeMessage1006(data []byte) (msg Message1006) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1006) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

type MessageAntennaDescriptor struct {
	ReferenceStationId      uint16 `struct:"uint16:12"`
	AntennaDescriptorLength uint8  `struct:"uint8"`
	AntennaDescriptor       string `struct:"[]byte,sizefrom=AntennaDescriptorLength"`
	AntennaSetupId          uint8  `struct:"uint8"`
}

// Antenna Descriptor
type Message1007 struct {
	AbstractMessage
	MessageAntennaDescriptor
}

func DeserializeMessage1007(data []byte) (msg Message1007) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1007) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

// Antenna Descriptor & Serial Number
type Message1008 struct {
	AbstractMessage
	MessageAntennaDescriptor
	SerialNumberLength uint8  `struct:"uint8"`
	SerialNumber       string `struct:"[]byte,sizefrom=SerialNumberLength"`
}

func DeserializeMessage1008(data []byte) (msg Message1008) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1008) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

// Receiver and Antenna Descriptors
type Message1033 struct {
	AbstractMessage
	MessageAntennaDescriptor
	AntennaSerialNumberLength     uint8  `struct:"uint8"`
	AntennaSerialNumber           string `struct:"[]byte,sizefrom=AntennaSerialNumberLength"`
	ReceiverTypeDescriptorLength  uint8  `struct:"uint8"`
	ReceiverTypeDescriptor        string `struct:"[]byte,sizefrom=ReceiverTypeDescriptorLength"`
	ReceiverFirmwareVersionLength uint8  `struct:"uint8"`
	ReceiverFirmwareVersion       string `struct:"[]byte,sizefrom=ReceiverFirmwareVersionLength"`
	ReceiverSerialNumberLength    uint8  `struct:"uint8"`
	ReceiverSerialNumber          string `struct:"[]byte,sizefrom=ReceiverSerialNumberLength"`
}

func DeserializeMessage1033(data []byte) (msg Message1033) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1033) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}
