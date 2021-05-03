package common

import "time"

func CreateDate(year, month, day int) time.Time {
	d := time.Date(year, time.Month(month), day, 0, 0, 0, 0, JP())
	return d
}
