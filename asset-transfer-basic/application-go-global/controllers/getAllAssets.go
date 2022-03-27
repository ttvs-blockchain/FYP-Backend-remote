package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllAssets(c *gin.Context) {
	log.Println("--> Evaluate Transaction: GetAllAssets, function returns all the current assets on the ledger")
	result, err := GlobalContract.EvaluateTransaction("GetAllAssets")
	if err != nil {
		log.Println("Failed to evaluate transaction: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(string(result))

	c.JSON(http.StatusOK, gin.H{
		"message": string(result),
	})
}
