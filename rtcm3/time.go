package rtcm3

import (
	"time"

	"github.com/go-gnss/rtcm"
)

// GPS Epoch Time (TOW)
func DF004(e uint32) time.Time {
	now := time.Now().UTC()
	return DF004Time(e, now)
}

func DF004Time(e uint32, start time.Time) time.Time {
	// TODO: GpsLeapSeconds should take a time to return leap seconds as at that time
	start = start.Add(rtcm.GpsLeapSeconds()) // UTC to GPS time
	sow := start.Truncate(time.Hour*24).AddDate(0, 0, -int(start.Weekday()))
	tow := time.Duration(e) * time.Millisecond
	return sow.Add(-(rtcm.GpsLeapSeconds())).Add(tow)
}

// BeiDou Epoch Time (TOW)
func DF427(e uint32, start time.Time) time.Time {
	start = start.Add(rtcm.GpsLeapSeconds() - 14*time.Second) // UTC to GPS time
	sow := start.Truncate(time.Hour*24).AddDate(0, 0, -int(start.Weekday()))
	tow := time.Duration(e) * time.Millisecond
	return sow.Add(-(rtcm.GpsLeapSeconds())).Add(tow + 14*time.Second)
}

// GPS Epoch Time 1s
func DF385(e uint32) time.Time {
	now := time.Now().UTC()
	sow := now.Truncate(time.Hour*24).AddDate(0, 0, -int(now.Weekday()))
	tow := time.Duration(e) * time.Second
	return sow.Add(-(rtcm.GpsLeapSeconds())).Add(tow)
}

// GLONASS Epoch Time (tk)
func DF034(e uint32, now time.Time) time.Time {
	hours := e / 3600000
	moduloGlonassHours := ((int(hours) - 3%24) + 24) % 24
	rest := int(e) - (int(hours) * 3600000)
	tod := time.Duration(rest+(moduloGlonassHours*3600000)) * time.Millisecond
	dow := now.Truncate(time.Hour * 24)
	return dow.Add(tod)
}

func GlonassTimeMSMNow(e uint32) time.Time {
	return GlonassTimeMSM(e, time.Now())
}

// DF416 + DF034
func GlonassTimeMSM(e uint32, now time.Time) time.Time {
	now = now.UTC().Add(3 * time.Hour)
	sow := now.Truncate(time.Hour*24).AddDate(0, 0, -int(now.Weekday()))
	dow := int((e >> 27) & 0x7)
	tod := time.Duration(e&0x7FFFFFF) * time.Millisecond
	// TODO: Do we truncate the 3 hours earlier? Write test
	return sow.AddDate(0, 0, dow).Add(tod).Add(-(3 * time.Hour))
}

// GLONASS Epoch Time 1s
func DF386(e uint32, now time.Time) time.Time {
	hours := e / 3600
	moduloGlonassHours := ((int(hours) - 3%24) + 24) % 24
	rest := int(e) - (int(hours) * 3600)
	tod := time.Duration(rest+(moduloGlonassHours*3600)) * time.Second
	dow := now.Truncate(time.Hour * 24)
	return dow.Add(tod)
}
