package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadAsset(c *gin.Context) {
	log.Println("--> Evaluate Transaction: ReadAsset, function returns an asset with a given assetID")
	CertNo := c.Query("CertNo")
	result, err := Contract.EvaluateTransaction("ReadAsset", CertNo)
	if err != nil {
		log.Println("Failed to evaluate transaction: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(string(result))

	c.JSON(http.StatusOK, gin.H{
		"message": string(result),
	})
}
