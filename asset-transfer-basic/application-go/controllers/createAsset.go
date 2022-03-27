package controllers

import (
	"asset-transfer-basic/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAsset(c *gin.Context) {

	var asset models.Asset
	if err := c.ShouldBindJSON(&asset); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("--> Submit Transaction: CreateAsset, creates new asset with %v, \n", asset)

	result, err := Contract.SubmitTransaction("CreateAsset",
		asset.CertNo, asset.ID, asset.Name, asset.Brand, asset.NumOfDose, asset.Time, asset.Issuer, asset.Remark)

	if err != nil {
		log.Printf("Failed to Submit transaction: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("--> Submit Transaction: CreateAsset, finish creates new asset with %v, \n", string(result))

	// TODO let local chain have hashed done
	var info = models.LocalChainInfo{"testid", "testhash", "testblknum", "testtimestamp"}
	models.InsertRow(asset, info)

	c.JSON(http.StatusOK, gin.H{
		"message": string(result),
	})
}
