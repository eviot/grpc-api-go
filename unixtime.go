package gapi

import "time"

func NewUnixTime(t time.Time) UnixTime {
	return UnixTime{
		Sec:  t.Unix(),
		Nsec: t.UnixNano(),
	}
}

func (ut UnixTime) Time() time.Time {
	return time.Unix(ut.Sec, ut.Nsec)
}
