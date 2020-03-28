package rtcm3

import (
	"encoding/binary"
	"github.com/go-restruct/restruct"
	"time"
)

// SSR GPS Orbit Correction
type Message1057 struct {
	AbstractMessage
	Epoch                    uint32 `struct:"uint32:20"`
	SSRUpdateInterval        uint8  `struct:"uint8:4"`
	MultipleMessageIndicator bool   `struct:"uint8:1,variantbool"`
	SatelliteReferenceDatum  bool   `struct:"uint8:1,variantbool"`
	IODSSR                   uint8  `struct:"uint8:4"`
	SSRProviderID            uint16 `struct:"uint16"`
	SSRSolutionID            uint8  `struct:"uint8:4"`
	Satellites               uint8  `struct:"uint8:6,sizeof=SatelliteData"`
	SatelliteData            []struct {
		SatelliteID        uint8 `struct:"uint8:6"`
		IODE               uint8 `struct:"uint8"`
		DeltaRadial        int32 `struct:"int32:22"`
		DeltaAlongTrack    int32 `struct:"int32:20"`
		DeltaCrossTrack    int32 `struct:"int32:20"`
		DotDeltaRadial     int32 `struct:"int32:21"`
		DotDeltaAlongTrack int32 `struct:"int32:19"`
		DotDeltaCrossTrack int32 `struct:"int32:19"`
	}
}

func DeserializeMessage1057(data []byte) (msg Message1057) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1057) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

func (msg Message1057) Time() time.Time {
	return DF385(msg.Epoch)
}

// SSR GPS Clock Correction
type Message1058 struct {
	AbstractMessage
	Epoch                    uint32 `struct:"uint32:20"`
	SSRUpdateInterval        uint8  `struct:"uint8:4"`
	MultipleMessageIndicator bool   `struct:"uint8:1,variantbool"`
	IODSSR                   uint8  `struct:"uint8:4"`
	SSRProviderID            uint16 `struct:"uint16"`
	SSRSolutionID            uint8  `struct:"uint8:4"`
	Satellites               uint8  `struct:"uint8:6,sizeof=SatelliteData"`
	SatelliteData            []struct {
		SatelliteID  uint8 `struct:"uint8:6"`
		DeltaClockC0 int32 `struct:"int32:22"`
		DeltaClockC1 int32 `struct:"int32:21"`
		DeltaClockC2 int32 `struct:"int32:27"`
	}
}

func DeserializeMessage1058(data []byte) (msg Message1058) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1058) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

func (msg Message1058) Time() time.Time {
	return DF385(msg.Epoch)
}

// SSR GPS Code Bias
type Message1059 struct {
	AbstractMessage
	Epoch                    uint32 `struct:"uint32:20"`
	SSRUpdateInterval        uint8  `struct:"uint8:4"`
	MultipleMessageIndicator bool   `struct:"uint8:1,variantbool"`
	IODSSR                   uint8  `struct:"uint8:4"`
	SSRProviderID            uint16 `struct:"uint16"`
	SSRSolutionID            uint8  `struct:"uint8:4"`
	Satellites               uint8  `struct:"uint8:6,sizeof=SatelliteData"`
	SatelliteData            []struct {
		SatelliteID         uint8 `struct:"uint8:6"`
		CodeBiasesProcessed uint8 `struct:"uint8:5,sizeof=CodeBiases"`
		CodeBiases          []struct {
			SignalTrackingModeIndicator uint8 `struct:"uint8:5"`
			CodeBias                    int16 `struct:"int16:14"`
		}
	}
}

func DeserializeMessage1059(data []byte) (msg Message1059) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1059) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

func (msg Message1059) Time() time.Time {
	return DF385(msg.Epoch)
}

// SSR GPS Combined Orbit and Clock Corrections
type Message1060 struct {
	AbstractMessage
	Epoch                    uint32 `struct:"uint32:20"`
	SSRUpdateInterval        uint8  `struct:"uint8:4"`
	MultipleMessageIndicator bool   `struct:"uint8:1,variantbool"`
	SatelliteReferenceDatum  bool   `struct:"uint8:1,variantbool"`
	IODSSR                   uint8  `struct:"uint8:4"`
	SSRProviderID            uint16 `struct:"uint16"`
	SSRSolutionID            uint8  `struct:"uint8:4"`
	Satellites               uint8  `struct:"uint8:6,sizeof=SatelliteData"`
	SatelliteData            []struct {
		SatelliteID        uint8 `struct:"uint8:6"`
		IOD                uint8 `struct:"uint8"`
		DeltaRadial        int32 `struct:"int32:22"`
		DeltaAlongTrack    int32 `struct:"int32:20"`
		DeltaCrossTrack    int32 `struct:"int32:20"`
		DotDeltaRadial     int32 `struct:"int32:21"`
		DotDeltaAlongTrack int32 `struct:"int32:19"`
		DotDeltaCrossTrack int32 `struct:"int32:19"`
		DeltaClockC0       int32 `struct:"int32:22"`
		DeltaClockC1       int32 `struct:"int32:21"`
		DeltaClockC2       int32 `struct:"int32:27"`
	}
}

func DeserializeMessage1060(data []byte) (msg Message1060) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1060) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

func (msg Message1060) Time() time.Time {
	return DF385(msg.Epoch)
}

// SSR GPS URA
type Message1061 struct {
	AbstractMessage
	Epoch                    uint32 `struct:"uint32:20"`
	SSRUpdateInterval        uint8  `struct:"uint8:4"`
	MultipleMessageIndicator bool   `struct:"uint8:1,variantbool"`
	IODSSR                   uint8  `struct:"uint8:4"`
	SSRProviderID            uint16 `struct:"uint16"`
	SSRSolutionID            uint8  `struct:"uint8:4"`
	Satellites               uint8  `struct:"uint8:6,sizeof=SatelliteData"`
	SatelliteData            []struct {
		SatelliteID uint8 `struct:"uint8:5"`
		SSRURA      uint8 `struct:"uint8:6"`
	}
}

func DeserializeMessage1061(data []byte) (msg Message1061) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1061) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

func (msg Message1061) Time() time.Time {
	return DF385(msg.Epoch)
}

// SSR GPS High Rate Clock Correction
type Message1062 struct {
	AbstractMessage
	Epoch                    uint32 `struct:"uint32:20"`
	SSRUpdateInterval        uint8  `struct:"uint8:4"`
	MultipleMessageIndicator bool   `struct:"uint8:1,variantbool"`
	IODSSR                   uint8  `struct:"uint8:4"`
	SSRProviderID            uint16 `struct:"uint16"`
	SSRSolutionID            uint8  `struct:"uint8:4"`
	Satellites               uint8  `struct:"uint8:6,sizeof=SatelliteData"`
	SatelliteData            []struct {
		SatelliteID     uint8 `struct:"uint8:5"`
		ClockCorrection int32 `struct:"uint32:22"`
	}
}

func DeserializeMessage1062(data []byte) (msg Message1062) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1062) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

func (msg Message1062) Time() time.Time {
	return DF385(msg.Epoch)
}

// SSR GLONASS Orbit Correction
type Message1063 struct {
	AbstractMessage
	Epoch                    uint32 `struct:"uint32:17"`
	SSRUpdateInterval        uint8  `struct:"uint8:4"`
	MultipleMessageIndicator bool   `struct:"uint8:1,variantbool"`
	SatelliteReferenceDatum  bool   `struct:"uint8:1,variantbool"`
	IODSSR                   uint8  `struct:"uint8:4"`
	SSRProviderID            uint16 `struct:"uint16"`
	SSRSolutionID            uint8  `struct:"uint8:4"`
	Satellites               uint8  `struct:"uint8:6,sizeof=SatelliteData"`
	SatelliteData            []struct {
		SatelliteID        uint8 `struct:"uint8:5"`
		IOD                uint8 `struct:"uint8"`
		DeltaRadial        int32 `struct:"int32:22"`
		DeltaAlongTrack    int32 `struct:"int32:20"`
		DeltaCrossTrack    int32 `struct:"int32:20"`
		DotDeltaRadial     int32 `struct:"int32:21"`
		DotDeltaAlongTrack int32 `struct:"int32:19"`
		DotDeltaCrossTrack int32 `struct:"int32:19"`
	}
}

func DeserializeMessage1063(data []byte) (msg Message1063) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1063) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

func (msg Message1063) Time() time.Time {
	return DF386(msg.Epoch, time.Now())
}

// SSR GLONASS Clock Correction
type Message1064 struct {
	AbstractMessage
	Epoch                    uint32 `struct:"uint32:17"`
	SSRUpdateInterval        uint8  `struct:"uint8:4"`
	MultipleMessageIndicator bool   `struct:"uint8:1,variantbool"`
	IODSSR                   uint8  `struct:"uint8:4"`
	SSRProviderID            uint16 `struct:"uint16"`
	SSRSolutionID            uint8  `struct:"uint8:4"`
	Satellites               uint8  `struct:"uint8:6,sizeof=SatelliteData"`
	SatelliteData            []struct {
		SatelliteID  uint8 `struct:"uint8:5"`
		DeltaClockC0 int32 `struct:"int32:22"`
		DeltaClockC1 int32 `struct:"int32:21"`
		DeltaClockC2 int32 `struct:"int32:27"`
	}
}

func DeserializeMessage1064(data []byte) (msg Message1064) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1064) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

func (msg Message1064) Time() time.Time {
	return DF386(msg.Epoch, time.Now())
}

// SSR GLONASS Code Bias
type Message1065 struct {
	AbstractMessage
	Epoch                    uint32 `struct:"uint32:17"`
	SSRUpdateInterval        uint8  `struct:"uint8:4"`
	MultipleMessageIndicator bool   `struct:"uint8:1,variantbool"`
	IODSSR                   uint8  `struct:"uint8:4"`
	SSRProviderID            uint16 `struct:"uint16"`
	SSRSolutionID            uint8  `struct:"uint8:4"`
	Satellites               uint8  `struct:"uint8:6,sizeof=SatelliteData"`
	SatelliteData            []struct {
		SatelliteID         uint8 `struct:"uint8:5"`
		CodeBiasesProcessed uint8 `struct:"uint8:5,sizeof=CodeBiases"`
		CodeBiases          []struct {
			SignalTrackingModeIndicator uint8 `struct:"uint8:5"`
			CodeBias                    int16 `struct:"int16:14"`
		}
	}
}

func DeserializeMessage1065(data []byte) (msg Message1065) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1065) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

func (msg Message1065) Time() time.Time {
	return DF386(msg.Epoch, time.Now())
}

// SSR GLONASS Combined Orbit and Clock Corrections
type Message1066 struct {
	AbstractMessage
	Epoch                    uint32 `struct:"uint32:17"`
	SSRUpdateInterval        uint8  `struct:"uint8:4"`
	MultipleMessageIndicator bool   `struct:"uint8:1,variantbool"`
	SatelliteReferenceDatum  bool   `struct:"uint8:1,variantbool"`
	IODSSR                   uint8  `struct:"uint8:4"`
	SSRProviderID            uint16 `struct:"uint16"`
	SSRSolutionID            uint8  `struct:"uint8:4"`
	Satellites               uint8  `struct:"uint8:6,sizeof=SatelliteData"`
	SatelliteData            []struct {
		SatelliteID        uint8 `struct:"uint8:5"`
		IOD                uint8 `struct:"uint8"`
		DeltaRadial        int32 `struct:"int32:22"`
		DeltaAlongTrack    int32 `struct:"int32:20"`
		DeltaCrossTrack    int32 `struct:"int32:20"`
		DotDeltaRadial     int32 `struct:"int32:21"`
		DotDeltaAlongTrack int32 `struct:"int32:19"`
		DotDeltaCrossTrack int32 `struct:"int32:19"`
		DeltaClockC0       int32 `struct:"int32:22"`
		DeltaClockC1       int32 `struct:"int32:21"`
		DeltaClockC2       int32 `struct:"int32:27"`
	}
}

func DeserializeMessage1066(data []byte) (msg Message1066) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1066) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

func (msg Message1066) Time() time.Time {
	return DF386(msg.Epoch, time.Now())
}

// SSR GLONASS URA
type Message1067 struct {
	AbstractMessage
	Epoch                    uint32 `struct:"uint32:17"`
	SSRUpdateInterval        uint8  `struct:"uint8:4"`
	MultipleMessageIndicator bool   `struct:"uint8:1,variantbool"`
	IODSSR                   uint8  `struct:"uint8:4"`
	SSRProviderID            uint16 `struct:"uint16"`
	SSRSolutionID            uint8  `struct:"uint8:4"`
	Satellites               uint8  `struct:"uint8:6,sizeof=SatelliteData"`
	SatelliteData            []struct {
		SatelliteID uint8 `struct:"uint8:5"`
		SSRURA      uint8 `struct:"uint8:6"`
	}
}

func DeserializeMessage1067(data []byte) (msg Message1067) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1067) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

func (msg Message1067) Time() time.Time {
	return DF386(msg.Epoch, time.Now())
}

// SSR GLONASS High Rate Clock Correction
type Message1068 struct {
	AbstractMessage
	Epoch                    uint32 `struct:"uint32:17"`
	SSRUpdateInterval        uint8  `struct:"uint8:4"`
	MultipleMessageIndicator bool   `struct:"uint8:1,variantbool"`
	IODSSR                   uint8  `struct:"uint8:4"`
	SSRProviderID            uint16 `struct:"uint16"`
	SSRSolutionID            uint8  `struct:"uint8:4"`
	Satellites               uint8  `struct:"uint8:6,sizeof=SatelliteData"`
	SatelliteData            []struct {
		SatelliteID     uint8 `struct:"uint8:5"`
		ClockCorrection int32 `struct:"uint32:22"`
	}
}

func DeserializeMessage1068(data []byte) (msg Message1068) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1068) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

func (msg Message1068) Time() time.Time {
	return DF386(msg.Epoch, time.Now())
}
