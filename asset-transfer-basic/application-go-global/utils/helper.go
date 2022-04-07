package utils

import (
	"asset-transfer-basic/models"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func GetDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetUnixTime() int64 {
	return time.Now().Unix()
}

func GetCurrentHash(x models.InputInfo) ([]byte, error) {

	keyHash, err := GetHashString(x.Key)
	if err != nil {
		return nil, err
	}
	x.Key = keyHash
	inputInfoJson, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}

	log.Println("--> Before Getting Hash: ", string(inputInfoJson))

	h := sha256.New()
	fmt.Printf("-->Debug get string before hash %s\n", string(inputInfoJson))
	if _, err := h.Write(inputInfoJson); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

func GetHashString(input string) (string, error) {

	h := sha256.New()
	if _, err := h.Write([]byte(input)); err != nil {
		log.Println("Failed to evaluate json: newInputInfo  %s\n", err)
		return "", err
	}

	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil

}
