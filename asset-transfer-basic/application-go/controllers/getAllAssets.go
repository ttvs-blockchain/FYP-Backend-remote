package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllAssets(c *gin.Context) {
	log.Printf("--> Evaluate Transaction: GetAllAssets, function returns all the current assets on the ledger")
	result, err := Contract.EvaluateTransaction("GetAllAssets")
	if err != nil {
		log.Printf("Failed to evaluate transaction: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("--> Success Transaction: GetAllAssets, result: %v\n", string(result))

	c.JSON(http.StatusOK, gin.H{
		"message": string(result),
	})
}
