package rtcm3

import (
	"encoding/binary"
	"github.com/go-restruct/restruct"
	"time"
)

// L1-Only GPS RTK Observables
type Message1001 struct {
	AbstractMessage
	ReferenceStationID uint16 `struct:"uint16:12"`
	Epoch              uint32 `struct:"uint32:30"`
	SynchronousGNSS    bool   `struct:"uint8:1,variantbool"`
	SignalsProcessed   uint8  `struct:"uint8:5,sizeof=SatelliteData"`
	SmoothingIndicator bool   `struct:"uint8:1,variantbool"`
	SmoothingInterval  uint8  `struct:"uint8:3"`
	SatelliteData      []struct {
		SatelliteID         uint8  `struct:"uint8:6"`
		L1CodeIndicator     bool   `struct:"uint8:1,variantbool"`
		L1Pseudorange       uint32 `struct:"uint32:24"`
		L1PhaseRange        int32  `struct:"int32:20"`
		L1LockTimeIndicator uint8  `struct:"uint8:7"`
	}
}

func DeserializeMessage1001(data []byte) (msg Message1001) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1001) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

func (msg Message1001) Time() time.Time {
	return DF004(msg.Epoch)
}

// Extended L1-Only GPS RTK Observables
type Message1002 struct {
	AbstractMessage
	ReferenceStationID uint16 `struct:"uint16:12"`
	Epoch              uint32 `struct:"uint32:30"`
	SynchronousGNSS    bool   `struct:"uint8:1,variantbool"`
	SignalsProcessed   uint8  `struct:"uint8:5,sizeof=SatelliteData"`
	SmoothingIndicator bool   `struct:"uint8:1,variantbool"`
	SmoothingInterval  uint8  `struct:"uint8:3"`
	SatelliteData      []struct {
		SatelliteID            uint8  `struct:"uint8:6"`
		L1CodeIndicator        bool   `struct:"uint8:1,variantbool"`
		L1Pseudorange          uint32 `struct:"uint32:24"`
		L1PhaseRange           int32  `struct:"int32:20"`
		L1LockTimeIndicator    uint8  `struct:"uint8:7"`
		L1PseudorangeAmbiguity uint8  `struct:"uint8"`
		L1CNR                  uint8  `struct:"uint8"`
	}
}

func DeserializeMessage1002(data []byte) (msg Message1002) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1002) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

func (msg Message1002) Time() time.Time {
	return DF004(msg.Epoch)
}

// L1&L2 GPS RTK Observables
type Message1003 struct {
	AbstractMessage
	ReferenceStationID uint16 `struct:"uint16:12"`
	Epoch              uint32 `struct:"uint32:30"`
	SynchronousGNSS    bool   `struct:"uint8:1,variantbool"`
	SignalsProcessed   uint8  `struct:"uint8:5,sizeof=SatelliteData"`
	SmoothingIndicator bool   `struct:"uint8:1,variantbool"`
	SmoothingInterval  uint8  `struct:"uint8:3"`
	SatelliteData      []struct {
		SatelliteID             uint8  `struct:"uint8:6"`
		L1CodeIndicator         bool   `struct:"uint8:1,variantbool"`
		L1Pseudorange           uint32 `struct:"uint32:24"`
		L1PhaseRange            int32  `struct:"int32:20"`
		L1LockTimeIndicator     uint8  `struct:"uint8:7"`
		L2CodeIndicator         uint8  `struct:"uint8:2"`
		L2PseudorangeDifference int16  `struct:"int16:14"`
		L2PhaseRange            int32  `struct:"int32:20"`
		L2LockTimeIndicator     uint8  `struct:"uint8:7"`
	}
}

func DeserializeMessage1003(data []byte) (msg Message1003) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1003) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

func (msg Message1003) Time() time.Time {
	return DF004(msg.Epoch)
}

// Extended L1&L2 GPS RTK Observables
type Message1004 struct {
	AbstractMessage
	ReferenceStationID uint16 `struct:"uint16:12"`
	Epoch              uint32 `struct:"uint32:30"`
	SynchronousGNSS    bool   `struct:"uint8:1,variantbool"`
	SignalsProcessed   uint8  `struct:"uint8:5,sizeof=SatelliteData"`
	SmoothingIndicator bool   `struct:"uint8:1,variantbool"`
	SmoothingInterval  uint8  `struct:"uint8:3"`
	SatelliteData      []struct {
		SatelliteID             uint8  `struct:"uint8:6"`
		L1CodeIndicator         bool   `struct:"uint8:1,variantbool"`
		L1Pseudorange           uint32 `struct:"uint32:24"`
		L1PhaseRange            int32  `struct:"int32:20"`
		L1LockTimeIndicator     uint8  `struct:"uint8:7"`
		L1PseudorangeAmbiguity  uint8  `struct:"uint8"`
		L1CNR                   uint8  `struct:"uint8"`
		L2CodeIndicator         uint8  `struct:"uint8:2"`
		L2PseudorangeDifference int16  `struct:"int16:14"`
		L2PhaseRange            int32  `struct:"int32:20"`
		L2LockTimeIndicator     uint8  `struct:"uint8:7"`
		L2CNR                   uint8  `struct:"uint8"`
	}
}

func DeserializeMessage1004(data []byte) (msg Message1004) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1004) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

func (msg Message1004) Time() time.Time {
	return DF004(msg.Epoch)
}

// Network Auxiliary Station Data Message
type Message1014 struct {
	AbstractMessage
	NetworkID                    uint8  `struct:"uint8:8"`
	SubnetworkID                 uint8  `struct:"uint8:4"`
	AuxiliaryStationsTransmitted uint8  `struct:"uint8:5"`
	MasterReferenceStationID     uint16 `struct:"uint16:12"`
	AuxiliaryReferenceStationID  uint16 `struct:"uint16:12"`
	AuxMasterDeltaLatitude       int32  `struct:"int32:20"`
	AuxMasterDeltaLongitude      int32  `struct:"int32:21"`
	AuxMasterDeltaHeight         int32  `struct:"int32:23"`
}

func DeserializeMessage1014(data []byte) (msg Message1014) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1014) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

// GPS Ionospheric Correction Differences
type Message1015 struct {
	AbstractMessage
	NetworkID                   uint8  `struct:"uint8"`
	SubnetworkID                uint8  `struct:"uint8:4"`
	Epoch                       uint32 `struct:"uint32:23"`
	MultipleMessageIndicator    bool   `struct:"uint8:1,variantbool"`
	MasterReferenceStationID    uint16 `struct:"uint16:12"`
	AuxiliaryReferenceStationID uint16 `struct:"uint16:12"`
	SatelliteCount              uint8  `struct:"uint8:4,sizeof=SatelliteData"`
	SatelliteData               []struct {
		SatelliteID                                 uint8 `struct:"uint8:6"`
		AmbiguityStatusFlag                         uint8 `struct:"uint8:2"`
		NonSyncCount                                uint8 `struct:"uint8:3"`
		IonosphericCarrierPhaseCorrectionDifference int32 `struct:"int32:17"`
	}
}

func DeserializeMessage1015(data []byte) (msg Message1015) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1015) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

// GPS Geometric Correction Differences
type Message1016 struct {
	AbstractMessage
	NetworkID                   uint8  `struct:"uint8"`
	SubnetworkID                uint8  `struct:"uint8:4"`
	Epoch                       uint32 `struct:"uint32:23"`
	MultipleMessageIndicator    bool   `struct:"uint8:1,variantbool"`
	MasterReferenceStationID    uint16 `struct:"uint16:12"`
	AuxiliaryReferenceStationID uint16 `struct:"uint16:12"`
	SatelliteCount              uint8  `struct:"uint8:4,sizeof=SatelliteData"`
	SatelliteData               []struct {
		SatelliteID                               uint8 `struct:"uint8:6"`
		AmbiguityStatusFlag                       uint8 `struct:"uint8:2"`
		NonSyncCount                              uint8 `struct:"uint8:3"`
		GeometricCarrierPhaseCorrectionDifference int32 `struct:"int32:17"`
		IODE                                      uint8 `struct:"uint8"`
	}
}

func DeserializeMessage1016(data []byte) (msg Message1016) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1016) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

// GPS Combined Geometric and Ionospheric Correction Differences
type Message1017 struct {
	AbstractMessage
	NetworkID                   uint8  `struct:"uint8"`
	SubnetworkID                uint8  `struct:"uint8:4"`
	Epoch                       uint32 `struct:"uint32:23"`
	MultipleMessageIndicator    bool   `struct:"uint8:1,variantbool"`
	MasterReferenceStationID    uint16 `struct:"uint16:12"`
	AuxiliaryReferenceStationID uint16 `struct:"uint16:12"`
	SatelliteCount              uint8  `struct:"uint8:4,sizeof=SatelliteData"`
	SatelliteData               []struct {
		SatelliteID                                 uint8 `struct:"uint8:6"`
		AmbiguityStatusFlag                         uint8 `struct:"uint8:2"`
		NonSyncCount                                uint8 `struct:"uint8:3"`
		GeometricCarrierPhaseCorrectionDifference   int32 `struct:"int32:17"`
		IODE                                        uint8 `struct:"uint8"`
		IonosphericCarrierPhaseCorrectionDifference int32 `struct:"uint32:17"`
	}
}

func DeserializeMessage1017(data []byte) (msg Message1017) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1017) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}
