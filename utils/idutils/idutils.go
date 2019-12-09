package idutils

import (
	"github.com/zheng-ji/goSnowFlake"
)

var (
	iw *goSnowFlake.IdWorker
)

func GetId() (id int64) {
	return s.NextId()
}
