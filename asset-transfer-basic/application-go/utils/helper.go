package utils

import "time"

func GetDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetUnixTime() int64 {
	return time.Now().Unix()
}
