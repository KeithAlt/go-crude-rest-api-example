package util

import "time"

// GetTime returns the current time in string format
func GetTime() string {
	return time.Now().Format("01-02-2006 15:04:05")
}
