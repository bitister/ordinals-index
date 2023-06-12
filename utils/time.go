package utils

import (
	"time"
)

func GetStime() int64 {
	//var currentTime = time.Now()
	//var startTime = time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 10, 0, 0, 0, currentTime.Location())
	//return startTime.Unix()
	return time.Now().Unix()
}

func GetDay() int64 {
	var currentTime = time.Now()
	var startTime = time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())
	return startTime.Unix()
}

// func GetStime() int64 {
// 	currentTime := time.Now()
// 	startTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())
// 	sTime := startTime.Unix()
// 	eTime := time.Now().Unix()
// 	duration := eTime - sTime
// 	count := duration / 600
// 	return startTime.Unix() + count*600
// }

func GetTime(timestamp int64) string {
	timeNow := time.Unix(timestamp, 0)
	return timeNow.Format("20060102150405")
}
