package util

import (
	"time"
)

type TagError struct {
	HtmlStatus int
	Inner      error
}

func GetTagError(htmlStatus int, inner error) *TagError {
	return &TagError{HtmlStatus: htmlStatus, Inner: inner}
}

func GetTimeNow() string {
	return time.Now().Format(time.RFC3339)
}
