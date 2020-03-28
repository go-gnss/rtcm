package rtcm3_test

import (
	"bufio"
	"fmt"
	"github.com/go-gnss/rtcm/rtcm3"
	"github.com/google/go-cmp/cmp"
	"os"
	"testing"
)

var (
	// TODO: 1014-7, 1101-3, 1104-6
	messages = []int{
		1001, 1002, 1003, 1004, 1005, 1006, 1007, 1008, 1009, 1010, 1011, 1012,
		1013, 1019, 1020, 1029, 1030, 1031, 1032, 1033, 1042, 1044, 1045, 1057,
		1058, 1059, 1060, 1063, 1064, 1065, 1066, 1071, 1072, 1073, 1074, 1075,
		1076, 1077, 1081, 1082, 1083, 1084, 1085, 1086, 1087, 1091, 1092, 1093,
		1094, 1095, 1096, 1097, 1104, 1107, 1111, 1112, 1113, 1114, 1115, 1116,
		1117, 1121, 1122, 1123, 1124, 1125, 1126, 1127, 1230,
	}
)

func readPayload(msgNumber uint) (payload []byte) {
	r, _ := os.Open("data/" + fmt.Sprint(msgNumber) + "_frame.bin")
	br := bufio.NewReader(r)
	frame, _ := rtcm3.DeserializeFrame(br)
	return frame.Payload
}

func TestSerializeDeserialize(t *testing.T) {
	for _, number := range messages {
		binary := readPayload(uint(number))
		if !cmp.Equal(rtcm3.DeserializeMessage(binary).Serialize(), binary) {
			t.Errorf("%v Deserialization not equal to binary", number)
		}
		if _, unknown := rtcm3.DeserializeMessage(binary).(rtcm3.MessageUnknown); unknown {
			t.Errorf("%v No Deserializer for Message", number)
		}
	}
}

func TestFrame(t *testing.T) {
	r, _ := os.Open("data/1117_frame.bin")
	br := bufio.NewReader(r)

	binary, _ := br.Peek(227)
	deserializedBinary, _ := rtcm3.DeserializeFrame(br)

	frame := rtcm3.Frame{
		Preamble: 211,
		Reserved: 0,
		Length:   121,
		Payload: []byte{
			0x45, 0xd0, 0x0, 0x6a, 0x9c, 0x8a, 0xa0, 0x0, 0x0, 0x71, 0x0, 0x0,
			0x0, 0x0, 0x0, 0x0, 0x0, 0x20, 0x0, 0x80, 0x0, 0x7f, 0xc1, 0x47,
			0x37, 0xbe, 0x7f, 0xfe, 0x48, 0x14, 0xb9, 0xc1, 0xf0, 0x7c, 0xc4,
			0xa, 0xf8, 0x14, 0xa0, 0x0, 0x12, 0xbc, 0xa1, 0x2c, 0x23, 0x89,
			0x33, 0x28, 0x9c, 0x6c, 0x7e, 0x2, 0xb7, 0xe1, 0x55, 0x77, 0xf5,
			0x8f, 0x81, 0x99, 0x4, 0xdc, 0xda, 0x5, 0x32, 0x48, 0x82, 0x73,
			0x86, 0x2, 0x1, 0x80, 0xff, 0xa8, 0x62, 0xff, 0x9d, 0x49, 0xfd,
			0xfe, 0x21, 0x7d, 0xe0, 0xa4, 0x4c, 0x96, 0x5, 0x81, 0x60, 0x58,
			0x16, 0x5, 0x81, 0x60, 0x0, 0x51, 0x93, 0xa4, 0x61, 0x19, 0xe4,
			0x19, 0xbc, 0xcf, 0x66, 0x86, 0xc2, 0xc, 0x45, 0xfc, 0xf3, 0xf6,
			0x20, 0xc1, 0xf1, 0x80, 0x80, 0x12, 0x40, 0x14, 0x0,
		},
		Crc: 0xfaf141,
	}

	if !cmp.Equal(frame.Serialize(), binary) {
		t.Errorf("Frame serialization and binary not equal")
	}

	if !cmp.Equal(frame, deserializedBinary) {
		t.Errorf("Frame and deserialized not equal")
	}
}
