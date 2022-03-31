package controllers

// only for debug
// import (
// 	"asset-transfer-basic/models"
// 	"asset-transfer-basic/utils"
// 	"log"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/google/uuid"
// )

// func CreateAsset(c *gin.Context) {

// 	var asset models.GlocalChainInfo
// 	if err := c.ShouldBindJSON(&asset); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	log.Printf("--> Submit Transaction: CreateAsset, creates new asset with %v, \n", asset)

// 	result, err := GlobalContract.SubmitTransaction("CreateAsset",
// 		uuid.New().String(),
// 		"localchain1",
// 		asset.GlobalChainTxHash)

// 	if err != nil {
// 		log.Printf("Failed to Submit transaction: %v\n", err)
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	log.Printf("--> Submit Transaction: CreateAsset, finish creates new asset with %v, \n", string(result))

// 	// TODO let local chain have hashed done
// 	var info = models.GlocalChainInfo{
// 		asset.CertIDList,
// 		asset.GlobalChainTxHash,
// 		1,
// 		utils.GetUnixTime()}

// 	err = models.InsertGlobal(info)

// 	if err != nil {
// 		log.Printf("Failed to Insert Row in DB for transaction: %v\n", err)
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": string(result),
// 	})
// }
