package util

import (
	"fmt"
	"time"
)

func LogHandler(message string, err error, function string) {
	fmt.Println(message, err, function, time.Now())
}
