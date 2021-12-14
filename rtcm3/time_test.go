package rtcm3_test

import (
	"testing"
	"time"

	"github.com/go-gnss/rtcm"
	"github.com/go-gnss/rtcm/rtcm3"
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

func TestDF004(t *testing.T) {
	utcInitialDate := time.Date(2021, 12, 12, 0, 0, 0, 0, time.UTC)
	utcDate := utcInitialDate
	end := 60 * 60 * 24 * 8 // week + 1 day
	for i := 0; i <= end; i++ {
		utcDate = utcInitialDate.Add(time.Duration(i) * time.Second)

		gpsDate := utcDate.Add(rtcm.GpsLeapSeconds())
		sow := gpsDate.Truncate(time.Hour*24).AddDate(0, 0, -int(gpsDate.Weekday()))
		e := uint32(gpsDate.Sub(sow).Milliseconds())

		latency := utcDate.Sub(rtcm3.DF004Time(e, utcDate)).Milliseconds()
		if latency != 0 {
			t.Errorf("DF004 time incorrect. Expected %v, got %v (%vms difference)", utcDate, rtcm3.DF004Time(e, utcDate), latency)
		}
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
