package rtcm3

import (
	"encoding/binary"
	"github.com/go-restruct/restruct"
)

// GPS Network RTK Residual Message
type Message1030 struct {
	AbstractMessage
	Epoch              uint32 `struct:"uint32:20"`
	ReferenceStationId uint16 `struct:"uint16:12"`
	NRefs              uint8  `struct:"uint8:7"`
	Satellites         uint8  `struct:"uint8:5,sizeof=SatelliteData"`
	SatelliteData      []struct {
		SatelliteId uint8  `struct:"uint8:6"`
		Soc         uint8  `struct:"uint8"`
		Sod         uint16 `struct:"uint16:9"`
		Soh         uint8  `struct:"uint8:6"`
		SIc         uint16 `struct:"uint16:10"`
		SId         uint16 `struct:"uint16:10"`
	}
}

func DeserializeMessage1030(data []byte) (msg Message1030) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1030) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

// TODO: Implement a Time method for GLONASS Residuals Epoch Time - DF225
// GLONASS Network RTK Residual Message
type Message1031 struct {
	AbstractMessage
	Epoch              uint32 `struct:"uint32:17"`
	ReferenceStationId uint16 `struct:"uint16:12"`
	NRefs              uint8  `struct:"uint8:7"`
	Satellites         uint8  `struct:"uint8:5,sizeof=SatelliteData"`
	SatelliteData      []struct {
		SatelliteId uint8  `struct:"uint8:6"`
		Soc         uint8  `struct:"uint8"`
		Sod         uint16 `struct:"uint16:9"`
		Soh         uint8  `struct:"uint8:6"`
		SIc         uint16 `struct:"uint16:10"`
		SId         uint16 `struct:"uint16:10"`
	}
}

func DeserializeMessage1031(data []byte) (msg Message1031) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1031) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

// Physical Reference Station Position Message
type Message1032 struct {
	AbstractMessage
	NonPhysicalReferenceStationId uint16 `struct:"uint16:12"`
	PhysicalReferenceStationId    uint16 `struct:"uint16:12"`
	EpochYear                     uint8  `struct:"uint8:6"`
	ArpEcefX                      int64  `struct:"int64:38"`
	ArpEcefY                      int64  `struct:"int64:38"`
	ArpEcefZ                      int64  `struct:"int64:38"`
}

func DeserializeMessage1032(data []byte) (msg Message1032) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1032) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}
