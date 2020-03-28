package rtcm3

import (
	"encoding/binary"
	"github.com/bamiaux/iobit"
	"github.com/go-restruct/restruct"
	"time"
)

// L1-Only GLONASS RTK Observables
type Message1009 struct {
	AbstractMessage
	ReferenceStationId uint16 `struct:"uint16:12"`
	Epoch              uint32 `struct:"uint32:27"`
	SynchronousGnss    bool   `struct:"uint8:1,variantbool"`
	SignalCount        uint8  `struct:"uint8:5,sizeof=SignalData"`
	SmoothingIndicator bool   `struct:"uint8:1,variantbool"`
	SmoothingInterval  uint8  `struct:"uint8:3"`
	SignalData         []struct {
		SatelliteId         uint8  `struct:"uint8:6"`
		L1CodeIndicator     bool   `struct:"uint8:1,variantbool"`
		FrequencyChannel    uint8  `struct:"uint8:5"`
		L1Pseudorange       uint32 `struct:"uint32:25"`
		L1PhaseRange        int32  `struct:"int32:20"`
		L1LockTimeIndicator uint8  `struct:"uint8:7"`
	}
}

func DeserializeMessage1009(data []byte) (msg Message1009) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1009) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

func (msg Message1009) Time() time.Time {
	return DF034(msg.Epoch, time.Now().UTC())
}

// Extended L1-Only GLONASS RTK Observables
type Message1010 struct {
	AbstractMessage
	ReferenceStationId uint16 `struct:"uint16:12"`
	Epoch              uint32 `struct:"uint32:27"`
	SynchronousGnss    bool   `struct:"uint8:1,variantbool"`
	SignalCount        uint8  `struct:"uint8:5,sizeof=SignalData"`
	SmoothingIndicator bool   `struct:"uint8:1,variantbool"`
	SmoothingInterval  uint8  `struct:"uint8:3"`
	SignalData         []struct {
		SatelliteId            uint8  `struct:"uint8:6"`
		L1CodeIndicator        bool   `struct:"uint8:1,variantbool"`
		FrequencyChannel       uint8  `struct:"uint8:5"`
		L1Pseudorange          uint32 `struct:"uint32:25"`
		L1PhaseRange           int32  `struct:"int32:20"`
		L1LockTimeIndicator    uint8  `struct:"uint8:7"`
		L1PseudorangeAmbiguity uint8  `struct:"uint8:7"`
		L1Cnr                  uint8  `struct:"uint8"`
	}
}

func DeserializeMessage1010(data []byte) (msg Message1010) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1010) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

func (msg Message1010) Time() time.Time {
	return DF034(msg.Epoch, time.Now().UTC())
}

// L1&L2 GLONASS RTK Observables
type Message1011 struct {
	AbstractMessage
	ReferenceStationId uint16 `struct:"uint16:12"`
	Epoch              uint32 `struct:"uint32:27"`
	SynchronousGnss    bool   `struct:"uint8:1,variantbool"`
	SignalCount        uint8  `struct:"uint8:5,sizeof=SignalData"`
	SmoothingIndicator bool   `struct:"uint8:1,variantbool"`
	SmoothingInterval  uint8  `struct:"uint8:3"`
	SignalData         []struct {
		SatelliteId         uint8  `struct:"uint8:6"`
		L1CodeIndicator     bool   `struct:"uint8:1,variantbool"`
		FrequencyChannel    uint8  `struct:"uint8:5"`
		L1Pseudorange       uint32 `struct:"uint32:25"`
		L1PhaseRange        int32  `struct:"int32:20"`
		L1LockTimeIndicator uint8  `struct:"uint8:7"`
		L2CodeIndicator     uint8  `struct:"uint8:2"`
		L2Pseudorange       uint16 `struct:"uint16:14"`
		L2PhaseRange        int32  `struct:"int32:20"`
		L2LockTimeIndicator uint8  `struct:"uint8:7"`
	}
}

func DeserializeMessage1011(data []byte) (msg Message1011) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1011) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

func (msg Message1011) Time() time.Time {
	return DF034(msg.Epoch, time.Now().UTC())
}

// Extended L1&L2 GLONASS RTK Observables
type Message1012 struct {
	AbstractMessage
	ReferenceStationId uint16 `struct:"uint16:12"`
	Epoch              uint32 `struct:"uint32:27"`
	SynchronousGnss    bool   `struct:"uint8:1,variantbool"`
	SignalCount        uint8  `struct:"uint8:5,sizeof=SignalData"`
	SmoothingIndicator bool   `struct:"uint8:1,variantbool"`
	SmoothingInterval  uint8  `struct:"uint8:3"`
	SignalData         []struct {
		SatelliteId            uint8  `struct:"uint8:6"`
		L1CodeIndicator        bool   `struct:"uint8:1,variantbool"`
		FrequencyChannel       uint8  `struct:"uint8:5"`
		L1Pseudorange          uint32 `struct:"uint32:25"`
		L1PhaseRange           int32  `struct:"int32:20"`
		L1LockTimeIndicator    uint8  `struct:"uint8:7"`
		L1PseudorangeAmbiguity uint8  `struct:"uint8:7"`
		L1Cnr                  uint8  `struct:"uint8"`
		L2CodeIndicator        uint8  `struct:"uint8:2"`
		L2Pseudorange          uint16 `struct:"uint16:14"`
		L2PhaseRange           int32  `struct:"int32:20"`
		L2LockTimeIndicator    uint8  `struct:"uint8:7"`
		L2Cnr                  uint8  `struct:"uint8"`
	}
}

func DeserializeMessage1012(data []byte) (msg Message1012) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1012) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

func (msg Message1012) Time() time.Time {
	return DF034(msg.Epoch, time.Now().UTC())
}

// GLONASS Ionospheric Correction Differences
type Message1037 struct {
	AbstractMessage
	NetworkID                       uint8  `struct:"uint8"`
	SubnetworkID                    uint8  `struct:"uint8:4"`
	Epoch                           uint32 `struct:"uint32:20"`
	MultipleMessageIndicator        bool   `struct:"uint8:1,variantbool"`
	MasterReferenceStationID        uint16 `struct:"uint16:12"`
	AuxiliaryReferenceStationID     uint16 `struct:"uint16:12"`
	DataEntriesCount                uint8  `struct:"uint8:4,sizeof=IonosphericCorrectionDifference"`
	IonosphericCorrectionDifference []struct {
		SatelliteID                                 uint8 `struct:"uint8:6"`
		AmbiguityStatusFlag                         uint8 `struct:"uint8:2"`
		NonSyncCount                                uint8 `struct:"uint8:3"`
		IonosphericCarrierPhaseCorrectionDifference int32 `struct:"int32:17"`
	}
}

func DeserializeMessage1037(data []byte) (msg Message1037) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1037) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

// GLONASS Geometric Correction Differences
type Message1038 struct {
	AbstractMessage
	NetworkID                       uint8  `struct:"uint8"`
	SubnetworkID                    uint8  `struct:"uint8:4"`
	Epoch                           uint32 `struct:"uint32:20"`
	MultipleMessageIndicator        bool   `struct:"uint8:1,variantbool"`
	MasterReferenceStationID        uint16 `struct:"uint16:12"`
	AuxiliaryReferenceStationID     uint16 `struct:"uint16:12"`
	DataEntriesCount                uint8  `struct:"uint8:4,sizeof=IonosphericCorrectionDifference"`
	IonosphericCorrectionDifference []struct {
		SatelliteID                               uint8 `struct:"uint8:6"`
		AmbiguityStatusFlag                       uint8 `struct:"uint8:2"`
		NonSyncCount                              uint8 `struct:"uint8:3"`
		GeometricCarrierPhaseCorrectionDifference int32 `struct:"int32:17"`
		IOD                                       uint8 `struct:"uint8"`
	}
}

func DeserializeMessage1038(data []byte) (msg Message1038) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1038) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

// GLONASS Combined Geometric and Ionospheric Correction Differences
type Message1039 struct {
	AbstractMessage
	NetworkID                       uint8  `struct:"uint8"`
	SubnetworkID                    uint8  `struct:"uint8:4"`
	Epoch                           uint32 `struct:"uint32:20"`
	MultipleMessageIndicator        bool   `struct:"uint8:1,variantbool"`
	MasterReferenceStationID        uint16 `struct:"uint16:12"`
	AuxiliaryReferenceStationID     uint16 `struct:"uint16:12"`
	DataEntriesCount                uint8  `struct:"uint8:4,sizeof=IonosphericCorrectionDifference"`
	IonosphericCorrectionDifference []struct {
		SatelliteID                                 uint8 `struct:"uint8:6"`
		AmbiguityStatusFlag                         uint8 `struct:"uint8:2"`
		NonSyncCount                                uint8 `struct:"uint8:3"`
		GeometricCarrierPhaseCorrectionDifference   int32 `struct:"int32:17"`
		IOD                                         uint8 `struct:"uint8"`
		IonosphericCarrierPhaseCorrectionDifference int32 `struct:"int32:17"`
	}
}

func DeserializeMessage1039(data []byte) (msg Message1039) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1039) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

// GLONASS L1 and L2 Code-Phase Biases
type Message1230 struct {
	AbstractMessage
	ReferenceStationId uint16
	CodePhaseBias      bool
	Reserved           uint8
	SignalsMask        uint8
	L1CACodePhaseBias  int16
	L1PCodePhaseBias   int16
	L2CACodePhaseBias  int16
	L2PCodePhaseBias   int16
}

func DeserializeMessage1230(data []byte) (msg Message1230) {
	r := iobit.NewReader(data)
	msg = Message1230{
		AbstractMessage: AbstractMessage{
			r.Uint16(12),
		},
		ReferenceStationId: r.Uint16(12),
		CodePhaseBias:      r.Bit(),
		Reserved:           r.Uint8(3),
		SignalsMask:        r.Uint8(4),
	}
	if (msg.SignalsMask & 8) == 8 {
		msg.L1CACodePhaseBias = r.Int16(16)
	}
	if (msg.SignalsMask & 4) == 4 {
		msg.L1PCodePhaseBias = r.Int16(16)
	}
	if (msg.SignalsMask & 2) == 2 {
		msg.L2CACodePhaseBias = r.Int16(16)
	}
	if (msg.SignalsMask & 1) == 1 {
		msg.L2PCodePhaseBias = r.Int16(16)
	}
	return msg
}

func (msg Message1230) Serialize() []byte {
	data := make([]byte, 4)
	w := iobit.NewWriter(data)
	w.PutUint16(12, msg.AbstractMessage.MessageNumber)
	w.PutUint16(12, msg.ReferenceStationId)
	w.PutBit(msg.CodePhaseBias)
	w.PutUint8(3, msg.Reserved)
	w.PutUint8(4, msg.SignalsMask)
	w.Flush()
	if (msg.SignalsMask & 8) == 8 {
		data = append(data, uint8(msg.L1CACodePhaseBias>>8), uint8(msg.L1CACodePhaseBias&0xff))
	}
	if (msg.SignalsMask & 4) == 4 {
		data = append(data, uint8(msg.L1PCodePhaseBias>>8), uint8(msg.L1PCodePhaseBias&0xff))
	}
	if (msg.SignalsMask & 2) == 2 {
		data = append(data, uint8(msg.L2CACodePhaseBias>>8), uint8(msg.L2CACodePhaseBias&0xff))
	}
	if (msg.SignalsMask & 1) == 1 {
		data = append(data, uint8(msg.L2PCodePhaseBias>>8), uint8(msg.L2PCodePhaseBias&0xff))
	}
	return data
}
