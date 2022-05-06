package date_utils

import "time"

const apiDateTimeFormat = "2006-01-01T10:10:10Z"

func NowTime() time.Time {
	return time.Now().UTC()
}

func NowTimeString() string {
	return NowTime().Format(apiDateTimeFormat)
}
