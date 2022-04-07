package controllers

import (
	"asset-transfer-basic/models"
	"asset-transfer-basic/utils"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func VerifyPath(c *gin.Context) {
	log.Println("--> Evaluate Transaction: Verify, function returns an asset with a given assetID")

	var verifyInfo models.VerifyInfo
	if err := c.ShouldBindJSON(&verifyInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := verifyInfo.VerifyPath.GlobalRootID

	log.Println("--> Evaluate Transaction: ID is ", id)

	result, err := GlobalContract.EvaluateTransaction("ReadAsset", id)
	if err != nil {
		log.Println("Failed to evaluate transaction: %s\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println(string(result))
	var globalChainInfo models.GlobalChainInfo
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

	merkleTreeRoot, err := base64.StdEncoding.DecodeString(globalChainInfo.MerkleTreeRoot)
	if err != nil {
		log.Printf("Failed to evaluate transaction: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var path [][]byte
	indexes := verifyInfo.VerifyPath.Indexes
	newCurrentHash, err := utils.GetCurrentHash(verifyInfo.VerifyInputInfo)

	if err != nil {
		log.Printf("Failed to obtain current hash: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println("--> Evaluate Transaction: new current hash is ", base64.StdEncoding.EncodeToString(newCurrentHash))

	for _, s := range verifyInfo.VerifyPath.Path {
		cur, err := base64.StdEncoding.DecodeString(s)
		if err != nil {
			log.Printf("Failed to evaluate transaction: %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		path = append(path, cur)
	}

	a := newCurrentHash

	for i, s := range path {

		if indexes[i] == 1 {
			a = utils.GetHash(a, s)
		} else {
			a = utils.GetHash(s, a)
		}

		if err != nil {
			log.Printf("Failed to evaluate transaction: %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Printf("the test is %s\n", base64.StdEncoding.EncodeToString(a))

	}

	resultCheck := bytes.Equal(a, merkleTreeRoot)
	fmt.Printf("\n\n\nCheck result root is %s, result is  %v\n\n\n", globalChainInfo.MerkleTreeRoot, resultCheck)

	c.JSON(http.StatusOK, gin.H{
		"message": resultCheck,
	})
}
