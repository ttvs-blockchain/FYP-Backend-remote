package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"log"
	"time"
)

func GetDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetUnixTime() int64 {
	return time.Now().Unix()
}

func GetHashString(input string) (string, error) {

	h := sha256.New()
	if _, err := h.Write([]byte(input)); err != nil {
		log.Println("Failed to evaluate json: newInputInfo  %s\n", err)
		return "", err
	}

	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil

}
