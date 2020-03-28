package rtcm3

import (
	"encoding/binary"
	"github.com/go-restruct/restruct"
)

// GPS Ephemerides
type Message1019 struct {
	AbstractMessage
	SatelliteID  uint8  `struct:"uint8:6"`
	WeekNumber   uint16 `struct:"uint16:10"`
	SVAccuracy   uint8  `struct:"uint8:4"`
	CodeOnL2     uint8  `struct:"uint8:2"`
	IDOT         int16  `struct:"int16:14"`
	IODE         uint8  `struct:"uint8"`
	Toc          uint16 `struct:"uint16"`
	Af2          int8   `struct:"int8"`
	Af1          int16  `struct:"int16"`
	Af0          int32  `struct:"int32:22"`
	IODC         uint16 `struct:"uint16:10"`
	Crs          int16  `struct:"int16"`
	DeltaN       int16  `struct:"int16"`
	M0           int32  `struct:"int32"`
	Cuc          int16  `struct:"int16"`
	Eccentricity uint32 `struct:"uint32"`
	Cus          int16  `struct:"int16"`
	SrA          uint32 `struct:"uint32"`
	Toe          uint16 `struct:"uint16"`
	Cic          int16  `struct:"int16"`
	Omega0       int32  `struct:"int32"`
	Cis          int16  `struct:"int16"`
	I0           int32  `struct:"int32"`
	Crc          int16  `struct:"int16"`
	Perigee      int32  `struct:"int32"`
	OmegaDot     int32  `struct:"int32:24"`
	Tgd          int8   `struct:"int8"`
	SVHealth     uint8  `struct:"uint8:6"`
	L2PDataFlag  bool   `struct:"uint8:1,variantbool"`
	FitInterval  bool   `struct:"uint8:1,variantbool"`
}

func DeserializeMessage1019(data []byte) (msg Message1019) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1019) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

type Sint uint

// GLONASS Ephemerides
type Message1020 struct {
	AbstractMessage
	SatelliteId               uint8  `struct:"uint8:6"`
	FrequencyChannel          uint8  `struct:"uint8:5"`
	AlmanacHealth             bool   `struct:"uint8:1,variantbool"`
	AlmanacHealthAvailability bool   `struct:"uint8:1,variantbool"`
	P1                        uint8  `struct:"uint8:2"`
	Tk                        uint16 `struct:"uint16:12"`
	Msb                       bool   `struct:"uint8:1,variantbool"`
	P2                        bool   `struct:"uint8:1,variantbool"`
	Tb                        uint8  `struct:"uint8:7"`
	XnTb1                     Sint   `struct:"uint32:24"`
	XnTb                      Sint   `struct:"uint32:27"`
	XnTb2                     Sint   `struct:"uint8:5"`
	YnTb1                     Sint   `struct:"uint32:24"`
	YnTb                      Sint   `struct:"uint32:27"`
	YnTb2                     Sint   `struct:"uint8:5"`
	ZnTb1                     Sint   `struct:"uint32:24"`
	ZnTb                      Sint   `struct:"uint32:27"`
	ZnTb2                     Sint   `struct:"uint8:5"`
	P3                        bool   `struct:"uint8:1,variantbool"`
	GammaN                    Sint   `struct:"uint16:11"`
	Mp                        uint8  `struct:"uint8:2"`
	M1n3                      bool   `struct:"uint8:1,variantbool"`
	TauN                      Sint   `struct:"uint32:22"`
	MDeltaTauN                Sint   `struct:"uint8:5"`
	En                        uint8  `struct:"uint8:5"`
	MP4                       bool   `struct:"uint8:1,variantbool"`
	MFt                       uint8  `struct:"uint8:4"`
	MNt                       uint16 `struct:"uint16:11"`
	MM                        uint8  `struct:"uint8:2"`
	AdditionalData            bool   `struct:"uint8:1,variantbool"`
	Na                        uint16 `struct:"uint16:11"`
	TauC                      Sint   `struct:"uint32"`
	MN4                       uint8  `struct:"uint8:5"`
	MTauGps                   Sint   `struct:"uint32:22"`
	M1n5                      bool   `struct:"uint8:1,variantbool"`
	Reserved                  uint8  `struct:"uint8:7"`
}

func DeserializeMessage1020(data []byte) (msg Message1020) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1020) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

// BDS Satellite Ephemeris Data
type Message1042 struct {
	AbstractMessage
	SatelliteId uint8  `struct:"uint8:6"`
	WeekNumber  uint16 `struct:"uint16:13"`
	SVURAI      uint8  `struct:"uint8:4"`
	IDOT        int16  `struct:"int16:14"`
	AODE        uint8  `struct:"uint8:5"`
	Toc         uint32 `struct:"uint32:17"`
	A2          int16  `struct:"int16:11"`
	A1          int32  `struct:"int32:22"`
	A0          int32  `struct:"int32:24"`
	AODC        uint8  `struct:"uint8:5"`
	Crs         int32  `struct:"int32:18"`
	DeltaN      int16  `struct:"int16"`
	M0          int32  `struct:"int32"`
	Cuc         int32  `struct:"int32:18"`
	E           uint32 `struct:"uint32"`
	Cus         int32  `struct:"int32:18"`
	ASquared    uint32 `struct:"uint32"`
	Toe         uint32 `struct:"uint32:17"`
	Cic         int32  `struct:"int32:18"`
	Omega0      int32  `struct:"int32"`
	Cis         int32  `struct:"int32:18"`
	I0          int32  `struct:"int32"`
	Crc         int32  `struct:"int32:18"`
	Omega       int32  `struct:"int32"`
	OmegaDot    int32  `struct:"int32:24"`
	TGD1        int16  `struct:"int16:10"`
	TGD2        int16  `struct:"int16:10"`
	SVHealth    bool   `struct:"uint8:1,variantbool"`
}

func DeserializeMessage1042(data []byte) (msg Message1042) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1042) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

// QZSS Ephemerides
type Message1044 struct {
	AbstractMessage
	SatelliteId uint8  `struct:"uint8:4"`
	Toc         uint16 `struct:"uint16"`
	Af2         int8   `struct:"int8"`
	Af1         int16  `struct:"int16"`
	Af0         int32  `struct:"int32:22"`
	IODE        uint8  `struct:"uint8"`
	Crs         int16  `struct:"int16"`
	DeltaN0     int16  `struct:"int16"`
	M0          int32  `struct:"int32"`
	Cuc         int16  `struct:"int16"`
	E           uint32 `struct:"uint32"`
	Cus         int16  `struct:"int16"`
	ASquared    uint32 `struct:"uint32"`
	Toe         uint16 `struct:"uint16"`
	Cic         int16  `struct:"int16"`
	Omega0      int32  `struct:"int32"`
	Cis         int16  `struct:"int16"`
	I0          int32  `struct:"int32"`
	Crc         int16  `struct:"int16"`
	OmegaN      int32  `struct:"int32"`
	OmegaDot    int32  `struct:"int32:24"`
	I0Dot       int16  `struct:"int16:14"`
	CodesL2     uint8  `struct:"uint8:2"`
	WeekNumber  uint16 `struct:"uint16:10"`
	URA         uint8  `struct:"uint8:4"`
	SVHealth    uint8  `struct:"uint8:6"`
	TGD         int8   `struct:"int8"`
	IODC        uint16 `struct:"uint16:10"`
	FitInterval bool   `struct:"uint8:1,variantbool"`
}

func DeserializeMessage1044(data []byte) (msg Message1044) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1044) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

// Galileo F/NAV Satellite Ephemeris Data
type Message1045 struct {
	AbstractMessage
	SatelliteId uint8  `struct:"uint8:6"`
	WeekNumber  uint16 `struct:"uint16:12"`
	IODNav      uint16 `struct:"uint16:10"`
	SVSISA      uint8  `struct:"uint8"`
	IDot        int16  `struct:"int16:14"`
	Toc         uint16 `struct:"uint16:14"`
	Af2         int8   `struct:"int8:6"`
	Af1         int32  `struct:"int32:21"`
	Af0         int32  `struct:"int32:31"`
	Crs         int16  `struct:"int16"`
	DeltaN0     int16  `struct:"int16"`
	M0          int32  `struct:"int32"`
	Cuc         int16  `struct:"int16"`
	E           uint32 `struct:"uint32"`
	Cus         int16  `struct:"int16"`
	ASquared    uint32 `struct:"uint32"`
	Toe         uint16 `struct:"uint16:14"`
	Cic         int16  `struct:"int16"`
	Omega0      int32  `struct:"int32"`
	Cis         int16  `struct:"int16"`
	I0          int32  `struct:"int32"`
	Crc         int16  `struct:"int16"`
	Omega       int32  `struct:"int32"`
	OmegaDot    int32  `struct:"int32:24"`
	BGDE5aE1    int16  `struct:"int16:10"`
	OSHS        uint8  `struct:"uint8:2"`
	OSDVS       bool   `struct:"uint8:1,variantbool"`
	Reserved    uint8  `struct:"uint8:7"`
}

func DeserializeMessage1045(data []byte) (msg Message1045) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1045) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

// Galileo I/NAV Satellite Ephemeris Data
type Message1046 struct {
	AbstractMessage
	SatelliteId           uint8  `struct:"uint8:6"`
	WeekNumber            uint16 `struct:"uint16:12"`
	IODNav                uint16 `struct:"uint16:10"`
	SISAIndex             uint8  `struct:"uint8"`
	IDot                  int16  `struct:"int16:14"`
	Toc                   uint16 `struct:"uint16:14"`
	Af2                   int8   `struct:"int8:6"`
	Af1                   int32  `struct:"int32:21"`
	Af0                   int32  `struct:"int32:31"`
	Crs                   int16  `struct:"int16"`
	DeltaN0               int16  `struct:"int16"`
	M0                    int32  `struct:"int32"`
	Cuc                   int16  `struct:"int16"`
	E                     uint32 `struct:"uint32"`
	Cus                   int16  `struct:"int16"`
	ASquared              uint32 `struct:"uint32"`
	Toe                   uint16 `struct:"uint16:14"`
	Cic                   int16  `struct:"int16"`
	Omega0                int32  `struct:"int32"`
	Cis                   int16  `struct:"int16"`
	I0                    int32  `struct:"int32"`
	Crc                   int16  `struct:"int16"`
	Omega                 int32  `struct:"int32"`
	OmegaDot              int32  `struct:"int32:24"`
	BGDE5aE1              int16  `struct:"int16:10"`
	BGDE5bE1              int16  `struct:"int16:10"`
	E5bSignalHealthStatus uint8  `struct:"uint8:2"`
	E5bDataValidityStatus bool   `struct:"uint8:1,variantbool"`
	E1BSignalHealthStatus uint8  `struct:"uint8:2"`
	E2BDataValidityStatus bool   `struct:"uint8:1,variantbool"`
	Reserved              uint8  `struct:"uint8:2"`
}

func DeserializeMessage1046(data []byte) (msg Message1046) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1046) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}
