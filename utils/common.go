package utils

import "time"

func ParseTime(stringDate string) time.Time {
	t, err := time.Parse("2006:04:02 15:04", stringDate)
	if err != nil {
		Logger.Error("parse time error:", err.Error())
		return time.Now()
	}
	return t
}
