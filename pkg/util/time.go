package util

import "time"

var date = time.Now()

// GetTime returns the current time in string format
func GetTime() string {
	return date.Format("01-02-2006 15:04:05")
}
