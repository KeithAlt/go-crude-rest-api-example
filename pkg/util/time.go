package util

import "time"

// TODO: make it easier to get time
var date = time.Now()

func GetTime() string {
	return date.Format("01-02-2006 15:04:05")
}
