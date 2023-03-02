package rtcm3_test

import (
	"testing"
	"time"

	"github.com/go-gnss/rtcm"
	"github.com/go-gnss/rtcm/rtcm3"
)

func TestDF034(t *testing.T) {
	beforeUtcZero := time.Date(2019, 2, 7, 23, 41, 40, 0, time.UTC)
	if time := rtcm3.DF034(9700000, beforeUtcZero); time != beforeUtcZero {
		t.Errorf("DF034 time incorrect before UTC 0: %v", time)
	}

	afterUtcZero := time.Date(2019, 2, 8, 00, 44, 44, 0, time.UTC)
	if time := rtcm3.DF034(13484000, afterUtcZero); time != afterUtcZero {
		t.Errorf("DF034 time incorrect after UTC 0: %v", time)
	}
}

func TestDF427(t *testing.T) {
	refTimeG0 := time.Date(2023, 2, 25, 23, 59, 42, 0, time.UTC)
	if time := rtcm3.DF427(604786000, refTimeG0); time != refTimeG0 {
		t.Errorf("DF427 time incorrect at GPS 0: %v, expected %v", time, refTimeG0)
	}

	refTimeB0 := time.Date(2023, 2, 25, 23, 59, 56, 0, time.UTC)
	if time := rtcm3.DF427(0, refTimeB0); time != refTimeB0 {
		t.Errorf("DF427 time incorrect at BeiDou 0: %v, expected %v", time, refTimeB0)
	}

	refTimeX := time.Date(2019, 2, 8, 00, 44, 44, 0, time.UTC)
	if time := rtcm3.DF034(13484000, refTimeX); time != refTimeX {
		t.Errorf("DF427 time incorrect: %v, expected %v", time, refTimeX)
	}
}

func TestGlonassTimeMSM(t *testing.T) {
	midweek := time.Date(2019, 3, 28, 4, 2, 9, 0, time.UTC)
	if time := rtcm3.GlonassTimeMSM(562199912, midweek); time != midweek {
		t.Errorf("GlonassMSM time incorrect midweek: %v, should be %v", time, midweek)
	}

	// the following three tests ensure that the 3 hour difference between Glonass time and UTC
	// is checked before calculating the start of week - the bitwise OR adds the day of week to
	// the start of the epoch (first 3 bits of epoch)
	preSundayBug := time.Date(2019, 2, 2, 20, 59, 59, 999000000, time.UTC)
	//                        2019, 2, 2, 23, 59, 59, 999000000, time.GLO
	if time := rtcm3.GlonassTimeMSM(86399999|0x30000000, preSundayBug); time != preSundayBug {
		t.Errorf("PreBug time incorrect: %v, should be %v", time, preSundayBug)
	}

	startSundayBug := time.Date(2019, 2, 2, 21, 0, 0, 0, time.UTC) // preSundayBug + 1 second
	//                          2019, 2, 3,  0, 0, 0, 0, time.GLO
	if time := rtcm3.GlonassTimeMSM(0, startSundayBug); time != startSundayBug {
		t.Errorf("StartBug time incorrect: %v, should be %v", time, startSundayBug)
	}

	endSundayBug := time.Date(2019, 2, 2, 23, 59, 59, 0, time.UTC)
	//                        2019, 2, 3, 02, 59, 59, 0, time.GLO
	if time := rtcm3.GlonassTimeMSM(10799000, endSundayBug); time != endSundayBug {
		t.Errorf("SundayBug time incorrect: %v, should be %v", time, endSundayBug)
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
