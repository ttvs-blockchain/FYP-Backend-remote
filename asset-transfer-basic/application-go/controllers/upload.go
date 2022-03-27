package controllers

import (
	"asset-transfer-basic/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Upload(c *gin.Context) {

	log.Printf("--> Evaluate Transaction: GetAllAssets, function returns all the current assets on the ledger")
	localRecord, err := Contract.EvaluateTransaction("GetAllAssets")
	if err != nil {
		log.Printf("Failed to evaluate transaction: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	merkelTreeRoot, err := utils.GetMerkelTreeRoot(localRecord)
	if err != nil {
		log.Printf("Failed to Create Merkel Tree: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("--> get Merkel Root of the tree finish\n")

	log.Printf("--> Evaluate Transaction: Upload to global chain")

	result, err := GlobalContract.SubmitTransaction("CreateAsset",
		uuid.New().String(),
		"testlocalchainID111",
		merkelTreeRoot,
	)
	if err != nil {
		log.Printf("Failed to Submit transaction: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("--> Submit Transaction: Upload, finish creates new asset with %v, \n", string(result))

	c.JSON(http.StatusOK, gin.H{
		"message": string(result),
	})
}
