package rtcm

// TODO: Consider where this code belongs
import (
    "time"
)

var LeapSeconds = []time.Time{
    time.Date(1972, 06, 30, 23, 59, 59, 0, time.UTC),
    time.Date(1972, 12, 31, 23, 59, 59, 0, time.UTC),
    time.Date(1973, 12, 31, 23, 59, 59, 0, time.UTC),
    time.Date(1974, 12, 31, 23, 59, 59, 0, time.UTC),
    time.Date(1975, 12, 31, 23, 59, 59, 0, time.UTC),
    time.Date(1976, 12, 31, 23, 59, 59, 0, time.UTC),
    time.Date(1977, 12, 31, 23, 59, 59, 0, time.UTC),
    time.Date(1978, 12, 31, 23, 59, 59, 0, time.UTC),
    time.Date(1979, 12, 31, 23, 59, 59, 0, time.UTC),
    time.Date(1981, 06, 30, 23, 59, 59, 0, time.UTC),
    time.Date(1982, 06, 30, 23, 59, 59, 0, time.UTC),
    time.Date(1983, 06, 30, 23, 59, 59, 0, time.UTC),
    time.Date(1985, 06, 30, 23, 59, 59, 0, time.UTC),
    time.Date(1987, 12, 31, 23, 59, 59, 0, time.UTC),
    time.Date(1989, 12, 31, 23, 59, 59, 0, time.UTC),
    time.Date(1990, 12, 31, 23, 59, 59, 0, time.UTC),
    time.Date(1992, 06, 30, 23, 59, 59, 0, time.UTC),
    time.Date(1993, 06, 30, 23, 59, 59, 0, time.UTC),
    time.Date(1994, 06, 30, 23, 59, 59, 0, time.UTC),
    time.Date(1995, 12, 30, 23, 59, 59, 0, time.UTC),
    time.Date(1997, 06, 30, 23, 59, 59, 0, time.UTC),
    time.Date(1998, 12, 30, 23, 59, 59, 0, time.UTC),
    time.Date(2005, 12, 30, 23, 59, 59, 0, time.UTC),
    time.Date(2008, 12, 30, 23, 59, 59, 0, time.UTC),
    time.Date(2012, 06, 30, 23, 59, 59, 0, time.UTC),
    time.Date(2015, 06, 30, 23, 59, 59, 0, time.UTC),
    time.Date(2016, 12, 30, 23, 59, 59, 0, time.UTC),
}

func GpsLeapSeconds() time.Duration {
    return time.Duration(len(LeapSeconds[9:])) * time.Second
}
