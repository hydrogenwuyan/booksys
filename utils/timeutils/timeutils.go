package timeutils

import "time"

// 毫秒
func Now() int64 {
	return time.Now().UnixNano() / 1e6
}

func MillisecondToTime(ms int64) time.Time {
	return time.Unix(ms/1000, (ms%1000)*int64(time.Millisecond))
}

func TimeToMillisecond(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}
