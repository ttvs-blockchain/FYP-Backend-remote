package controllers

import (
	"asset-transfer-basic/models"
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func VerifyCert(c *gin.Context) {
	log.Println("--> Evaluate Transaction: VerifyCert, function returns all the current assets on the ledger")

	var verifyInfo models.VerifyInfo
	if err := c.ShouldBindJSON(&verifyInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("--> Get Path and Cert on local server, %s", verifyInfo.CertNo)

	resp, err := http.Get("http://localhost:8080/GetPath?CertNo=" + verifyInfo.CertNo)
	if err != nil || resp.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	fmt.Println(resp.StatusCode)

	var response models.GetPathInfo
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println("Failed to evaluate json: GetPathInfo 1111 %s\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// fmt.Printf("-->MerkelTreePath is %s,\n", response.MKTreePathDetail)

	var mktreePath models.MerkelTreePath
	err = json.Unmarshal([]byte(response.MKTreePathDetail), &mktreePath)
	if err != nil {
		log.Println("Failed to evaluate json: GetPathInfo 222 %s, %s\n", response.MKTreePathDetail, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := mktreePath.GlobalID

	oldCurrentHash, err := base64.StdEncoding.DecodeString(mktreePath.CurrentHash)
	if err != nil {
		log.Println("Failed to evaluate json: GetPathInfo 222 %s, %s\n", response.MKTreePathDetail, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	personHash := verifyInfo.PersonHash

	asset := response.AssetDetail

	newInputInfo := models.InputInfo{asset, personHash}

	input := newInputInfo.CertDetail.Time
	fmt.Printf("-->old time is %s\n", input)

	newInputInfo.CertDetail.Time = strings.ReplaceAll(input[0:10], "-", "/")

	newInputInfoJson, err := json.Marshal(newInputInfo)

	if err != nil {
		log.Println("Failed to evaluate json: newInputInfo  %s\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("-->New Current Hash is %s\n", newInputInfoJson)

	h := sha256.New()
	if _, err := h.Write([]byte(newInputInfoJson)); err != nil {
		log.Println("Failed to evaluate json: newInputInfo  %s\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newCurrentHash := h.Sum(nil)

	if bytes.Equal(oldCurrentHash, newCurrentHash) != true {
		log.Printf("Current Hash Not Match, old is %s, new is %s\n",
			mktreePath.CurrentHash,
			base64.StdEncoding.EncodeToString(h.Sum(nil)))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Current Hash Not Match"})
		return
	}

	log.Println("--> Evaluate Transaction: ID is ", id)

	result, err := GlobalContract.EvaluateTransaction("ReadAsset", id)
	if err != nil {
		log.Println("Failed to evaluate transaction: %s\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println(string(result))
	var globalChainInfo models.GlocalChainInfo
	err = json.Unmarshal(result, &globalChainInfo)
	if err != nil {
		log.Printf("Failed to evaluate transaction: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err != nil {
		log.Printf("Failed to evaluate transaction: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	merkelTreeRoot, err := base64.StdEncoding.DecodeString(globalChainInfo.MerkelTreeRoot)

	fmt.Printf("--> Merkel Tree Root is %v \n", merkelTreeRoot)
	if err != nil {
		log.Printf("Failed to evaluate transaction: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var path [][]byte
	indexes := mktreePath.Indexes
	currentHash, err := base64.StdEncoding.DecodeString(mktreePath.CurrentHash)
	if err != nil {
		log.Printf("Failed to evaluate transaction: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, s := range mktreePath.Path {
		cur, err := base64.StdEncoding.DecodeString(s)
		if err != nil {
			log.Printf("Failed to evaluate transaction: %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		path = append(path, cur)
	}

	a := currentHash

	for i, s := range path {

		if indexes[i] == 1 {
			a = GetHash(a, s)
		} else {
			a = GetHash(s, a)
		}

		if err != nil {
			log.Printf("Failed to evaluate transaction: %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Printf("the test is %s\n", base64.StdEncoding.EncodeToString(a))

	}

	resultCheck := bytes.Equal(a, merkelTreeRoot)
	fmt.Printf("\n\n\nCheck result by path %v\n\n\n", resultCheck)

	c.JSON(http.StatusOK, gin.H{
		"message": resultCheck,
	})
}
