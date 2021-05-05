package date_utils

import (
	"time"
)

const (
	apiDataLayout = "2006-01-02T15:04:05Z"
	apiDbLayout = "2006-01-02 15:04:05"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	now := time.Now().UTC()
	return now.Format(apiDataLayout)
}

func GetNowDBFormat() string {
	now := time.Now().UTC()
	return now.Format(apiDbLayout)
}