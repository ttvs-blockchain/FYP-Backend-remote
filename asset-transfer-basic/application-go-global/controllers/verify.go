package controllers

import (
	"asset-transfer-basic/models"
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Verify(c *gin.Context) {
	log.Println("--> Evaluate Transaction: Verify, function returns an asset with a given assetID")

	var mktreePath models.MerkelTreePath
	if err := c.ShouldBindJSON(&mktreePath); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := mktreePath.GlobalID

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
	fmt.Printf("\n\n\nCheck result root is %s, result is  %v\n\n\n", merkelTreeRoot, resultCheck)

	c.JSON(http.StatusOK, gin.H{
		"message": resultCheck,
	})
}

func GetHash(a []byte, b []byte) []byte {

	h := sha256.New()
	fmt.Printf("the input is %s,    %s\n",
		base64.StdEncoding.EncodeToString(a),
		base64.StdEncoding.EncodeToString(b))

	if _, err := h.Write(append(a, b...)); err != nil {
		// return nil, err
		fmt.Printf("GG")
	}

	fmt.Printf("the out is %s\n",
		base64.StdEncoding.EncodeToString(h.Sum(nil)))
	return h.Sum(nil)
}
