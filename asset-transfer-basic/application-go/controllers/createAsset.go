package controllers

import (
	"asset-transfer-basic/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAsset(c *gin.Context) {

	var input models.InputInfo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.CertDetail.Time = input.CertDetail.Time[0:16]
	fmt.Printf("--> Input check: %s\n", input)
	asset := input.CertDetail
	personHash := input.PersonHash

	log.Printf("--> Submit Transaction: CreateAsset, creates new asset with %v, \n", asset)
	result, err := Contract.SubmitTransaction("CreateAsset",
		asset.CertNo, asset.ID, asset.Name, asset.Brand, asset.NumOfDose, asset.Time, asset.Issuer, asset.Remark)

	if err != nil {
		log.Printf("Failed to Submit transaction: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("--> Submit Transaction: CreateAsset, start store in DB\n")

	// TODO let local chain have hashed done

	err = models.InsertLocalDBCert(asset, personHash)

	if err != nil {
		log.Printf("Failed to Insert Row in DB for transaction: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": string(result),
	})
}
