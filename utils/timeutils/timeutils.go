package timeutils

import "time"

// 毫秒
func Now() int64 {
	return time.Now().UnixNano() / 1e6
}
