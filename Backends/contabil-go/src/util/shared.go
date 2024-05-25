package util

import "time"

func GetTimeNow() string {
	return time.Now().Format(time.RFC3339)
}
