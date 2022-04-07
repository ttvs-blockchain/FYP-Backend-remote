package controllers

// import (
// 	"asset-transfer-basic/models"
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func GetPath(c *gin.Context) {
// 	log.Println("--> Evaluate Transaction: ReadAsset, function returns an asset with a given assetID")
// 	certID := c.Query("CertID")
// 	asset, personInfoHash, path, err := models.ReadPath(certID)

// 	fmt.Printf("debug 1111, asset is %s, personInfoHash is %s, path is %s\n", asset, personInfoHash, path)
// 	if err != nil {
// 		// log.Println("Failed to evaluate transaction: %v\n", err)
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	fmt.Printf("debug 2222, asset is %s, personInfoHash is %s, path is %s\n", asset, personInfoHash, path)

// 	c.JSON(http.StatusOK, gin.H{
// 		"message":        "okay",
// 		"asset":          asset,
// 		"personInfoHash": personInfoHash,
// 		"path":           path,
// 	})
// }
