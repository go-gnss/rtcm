package rtcm3

import (
	"encoding/binary"
	"github.com/go-restruct/restruct"
)

type Message1021 struct {
	AbstractMessage
	SourceNameCounter                      uint8  `struct:"uint8:5"`
	SourceName                             string `struct:"[]byte,sizefrom=SourceNameCounter"`
	TargetNameCounter                      uint8  `struct:"uint8:5"`
	TargetName                             string `struct:"[]byte,sizefrom=TargetNameCounter"`
	SystemIdentificationNumber             uint8  `struct:"uint8"`
	UtilizedTransformationMessageIndicator uint16 `struct:"uint16:10"`
	PlateNumber                            uint8  `struct:"uint5"`
	ComputationIndicator                   uint8  `struct:"uint8"`
	HeightIndicator                        uint8  `struct:"uint8:2"`
	PhiV                                   int32  `struct:"int32:19"`
	LambdaV                                int32  `struct:"int32:20"`
	DeltaPhiV                              uint16 `struct:"uint16:14"`
	DeltaLambdaV                           uint16 `struct:"uint16:14"`
	dX                                     int32  `struct:"int32:23"`
	dY                                     int32  `struct:"int32:23"`
	dZ                                     int32  `struct:"int32:23"`
	R1                                     int32  `struct:"int32"`
	R2                                     int32  `struct:"int32"`
	R3                                     int32  `struct:"int32"`
	dS                                     int32  `struct:"int32:25"`
	AddAS                                  int32  `struct:"int32:24"`
	AddBS                                  int32  `struct:"int32:25"`
	AddATau                                int32  `struct:"int32:24"`
	AddBTau                                int32  `struct:"int32:25"`
	HorizontalHMQualityIndicator           uint8  `struct:"uint8:3"`
	VerticalHMQualityIndicator             uint8  `struct:"uint8:3"`
}

func DeserializeMessage1021(data []byte) (msg Message1021) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1021) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

type Message1022 struct {
	AbstractMessage
	SourceNameCounter                      uint8  `struct:"uint8:5"`
	SourceName                             string `struct:"[]byte,sizefrom=SourceNameCounter"`
	TargetNameCounter                      uint8  `struct:"uint8:5"`
	TargetName                             string `struct:"[]byte,sizefrom=TargetNameCounter"`
	SystemIdentificationNumber             uint8  `struct:"uint8"`
	UtilizedTransformationMessageIndicator uint16 `struct:"uint16:10"`
	PlateNumber                            uint8  `struct:"uint5"`
	ComputationIndicator                   uint8  `struct:"uint8"`
	HeightIndicator                        uint8  `struct:"uint8:2"`
	PhiV                                   int32  `struct:"int32:19"`
	LambdaV                                int32  `struct:"int32:20"`
	DeltaPhiV                              uint16 `struct:"uint16:14"`
	DeltaLambdaV                           uint16 `struct:"uint16:14"`
	dX                                     int32  `struct:"int32:23"`
	dY                                     int32  `struct:"int32:23"`
	dZ                                     int32  `struct:"int32:23"`
	R1                                     int32  `struct:"int32"`
	R2                                     int32  `struct:"int32"`
	R3                                     int32  `struct:"int32"`
	dS                                     int32  `struct:"int32:25"`
	XP                                     int64  `struct:"int64:35"`
	YP                                     int64  `struct:"int64:35"`
	ZP                                     int64  `struct:"int64:35"`
	AddAS                                  int32  `struct:"int32:24"`
	AddBS                                  int32  `struct:"int32:25"`
	AddATau                                int32  `struct:"int32:24"`
	AddBTau                                int32  `struct:"int32:25"`
	HorizontalHMQualityIndicator           uint8  `struct:"uint8:3"`
	VerticalHMQualityIndicator             uint8  `struct:"uint8:3"`
}

func DeserializeMessage1022(data []byte) (msg Message1022) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1022) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

type Message1023 struct {
	AbstractMessage
	SystemIdentificationNumber             uint8  `struct:"uint8"`
	HorizontalShiftIndicator               bool   `struct:"uint8:1,variantbool"`
	VerticalShiftIndicator                 bool   `struct:"uint8:1,variantbool"`
	Phi0                                   int32  `struct:"int32:21"`
	Lambda0                                int32  `struct:"int32:22"`
	DeltaPhi                               uint16 `struct:"uint16:12"`
	DeltaLambda                            uint16 `struct:"uint16:12"`
	MeanDeltaPhi                           int8   `struct:"int8"`
	MeanDeltaLambda                        int8   `struct:"int8"`
	MeanDeltaH                             int16  `struct:"int16:15"`
	DeltaPhiI                              int16  `struct:"int16:9"`
	DeltaLambdaI                           int16  `struct:"int16:9"`
	DeltaHI                                int16  `struct:"int16:9"`
	HorizontalInterpolationMethodIndicator uint8  `struct:"uint8:2"`
	VerticalInterpolationMethodIndicator   uint8  `struct:"uint8:2"`
	HorizontalGridQualityIndicator         uint8  `struct:"uint8:3"`
	VerticalGridQualityIndicator           uint8  `struct:"uint8:3"`
	ModifiedJulianDayNumber                uint16 `struct:"uint16"`
}

func DeserializeMessage1023(data []byte) (msg Message1023) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1023) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

type Message1024 struct {
	AbstractMessage
	SystemIdentificationNumber             uint8  `struct:"uint8"`
	HorizontalShiftIndicator               bool   `struct:"uint8:1,variantbool"`
	VerticalShiftIndicator                 bool   `struct:"uint8:1,variantbool"`
	N0                                     int32  `struct:"int32:25"`
	E0                                     int32  `struct:"int32:26"`
	MeanDeltaN                             int16  `struct:"int16:10"`
	MeanDeltaE                             int16  `struct:"int16:10"`
	MeanDeltaH                             int16  `struct:"int16:15"`
	DeltaNI                                int16  `struct:"int16:9"`
	DeltaEI                                int16  `struct:"int16:9"`
	DeltaHI                                int16  `struct:"int16:9"`
	HorizontalInterpolationMethodIndicator uint8  `struct:"uint8:2"`
	VerticalInterpolationMethodIndicator   uint8  `struct:"uint8:2"`
	HorizontalGridQualityIndicator         uint8  `struct:"uint8:3"`
	VerticalGridQualityIndicator           uint8  `struct:"uint8:3"`
	ModifiedJulianDayNumber                uint16 `struct:"uint16"`
}

func DeserializeMessage1024(data []byte) (msg Message1024) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1024) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

type Message1025 struct {
	AbstractMessage
	SystemIdentificationNumber uint8 `struct:"uint8"`
	LaNO                       int64 `struct:"int64:34"`
	LoNO                       int64 `struct:"int64:35"`
	AddSN0                     int32 `struct:"int32:30"`
	FE                         int64 `struct:"int64:36"`
	FN                         int64 `struct:"int64:35"`
}

func DeserializeMessage1025(data []byte) (msg Message1025) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1025) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

type Message1026 struct {
	AbstractMessage
	SystemIdentificationNumber uint8 `struct:"uint8"`
	ProjectionType             uint8 `struct:"uint8:6"`
	LaNO                       int64 `struct:"int64:34"`
	LoNO                       int64 `struct:"int64:35"`
	LaSP1                      int64 `struct:"int64:34"`
	LaSP2                      int64 `struct:"int64:34"`
	EFO                        int64 `struct:"int64:36"`
	NFO                        int64 `struct:"int64:35"`
}

func DeserializeMessage1026(data []byte) (msg Message1026) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1026) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}

type Message1027 struct {
	AbstractMessage
	SystemIdentificationNumber uint8  `struct:"uint8"`
	ProjectionType             uint8  `struct:"uint8:6"`
	RectificationFlag          bool   `struct:"uint8:1,variantbool"`
	LaPC                       int64  `struct:"int64:34"`
	LoPC                       int64  `struct:"int64:35"`
	AzIL                       uint64 `struct:"uint64:35"`
	DiffARSG                   int32  `struct:"int32:26"`
	AddSIL                     uint32 `struct:"uint32:30"`
	EPC                        uint64 `struct:"uint64:36"`
	NPC                        int64  `struct:"int64:35"`
}

func DeserializeMessage1027(data []byte) (msg Message1027) {
	restruct.Unpack(data, binary.BigEndian, &msg)
	return msg
}

func (msg Message1027) Serialize() []byte {
	data, _ := restruct.Pack(binary.BigEndian, &msg)
	return data
}
