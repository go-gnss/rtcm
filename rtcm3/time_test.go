package rtcm3_test

import (
	"github.com/go-gnss/rtcm/rtcm3"
	"testing"
	"time"
)

func TestDF034(t *testing.T) {
	beforeUtcZero := time.Date(2019, 2, 7, 23, 41, 40, 0, time.UTC)
	if beforeUtcZero != rtcm3.DF034(9700000, beforeUtcZero) {
		t.Errorf("DF034 time incorrect before UTC 0: %v", rtcm3.DF034(9700000, beforeUtcZero))
	}

	afterUtcZero := time.Date(2019, 2, 8, 00, 44, 44, 0, time.UTC)
	if afterUtcZero != rtcm3.DF034(13484000, afterUtcZero) {
		t.Errorf("DF034 time incorrect after UTC 0: %v", rtcm3.DF034(13484000, afterUtcZero))
	}
}

func TestDF386(t *testing.T) {
	beforeUtcZero := time.Date(2019, 2, 7, 23, 41, 40, 0, time.UTC)
	if beforeUtcZero != rtcm3.DF386(9700, beforeUtcZero) {
		t.Errorf("DF386 time incorrect before UTC 0: %v", rtcm3.DF386(9700, beforeUtcZero))
	}

	afterUtcZero := time.Date(2019, 2, 8, 00, 44, 44, 0, time.UTC)
	if afterUtcZero != rtcm3.DF386(13484, afterUtcZero) {
		t.Errorf("DF386 time incorrect after UTC 0: %v", rtcm3.DF386(13484, afterUtcZero))
	}
}
