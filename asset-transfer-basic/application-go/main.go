/*
Copyright 2020 IBM All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

type Asset struct {
	CertNo    string `form:"certNo" json:"certNo" xml:"certNo"  binding:"required"`
	ID        string `form:"id" json:"id" xml:"color"  binding:"required"`
	Name      string `form:"name" json:"name" xml:"name"  binding:"required"`
	Brand     string `form:"brand" json:"brand" xml:"brand"  binding:"required"`
	NumOfDose string `form:"numOfDose" json:"numOfDose" xml:"numOfDose"  binding:"required"`
	Time      string `form:"time" json:"time" xml:"time"  binding:"required"`
	Issuer    string `form:"issuer" json:"issuer" xml:"issuer"  binding:"required"`
	Remark    string `form:"remark" json:"remark" xml:"remark"  binding:""`
}

func main() {

	log.Println("============ application-golang starts ============")

	err := os.Setenv("DISCOVERY_AS_LOCALHOST", "true")
	if err != nil {
		log.Fatalf("Error setting DISCOVERY_AS_LOCALHOST environemnt variable: %v", err)
	}

	wallet, err := gateway.NewFileSystemWallet("wallet")
	if err != nil {
		log.Fatalf("Failed to create wallet: %v", err)
	}

	if !wallet.Exists("appUser") {
		err = populateWallet(wallet)
		if err != nil {
			log.Fatalf("Failed to populate wallet contents: %v", err)
		}
	}

	ccpPath := filepath.Join(
		"..",
		"..",
		"test-network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"connection-org1.yaml",
	)

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),
		gateway.WithIdentity(wallet, "appUser"),
	)
	if err != nil {
		log.Fatalf("Failed to connect to gateway: %v", err)
	}
	defer gw.Close()

	network, err := gw.GetNetwork("mychannel")
	if err != nil {
		log.Fatalf("Failed to get network: %v", err)
	}

	contract := network.GetContract("basic")

	// log.Println("--> Submit Transaction: InitLedger, function creates the initial set of assets on the ledger")
	// result, err := contract.SubmitTransaction("InitLedger")
	// if err != nil {
	// 	log.Fatalf("Failed to Submit transaction: %v", err)
	// }
	// log.Println(string(result))

	// log.Println("--> Evaluate Transaction: GetAllAssets, function returns all the current assets on the ledger")
	// result, err = contract.EvaluateTransaction("GetAllAssets")
	// if err != nil {
	// 	log.Fatalf("Failed to evaluate transaction: %v", err)
	// }
	// log.Println(string(result))

	// log.Println("--> Submit Transaction: CreateAsset, creates new asset with ID, color, owner, size, and appraisedValue arguments")
	// result, err = contract.SubmitTransaction("CreateAsset", "asset13", "yellow", "5", "Tom", "1300")
	// if err != nil {
	// 	log.Fatalf("Failed to Submit transaction: %v", err)
	// }
	// log.Println(string(result))

	// log.Println("--> Evaluate Transaction: ReadAsset, function returns an asset with a given assetID")
	// result, err = contract.EvaluateTransaction("ReadAsset", "asset13")
	// if err != nil {
	// 	log.Fatalf("Failed to evaluate transaction: %v\n", err)
	// }
	// log.Println(string(result))

	// log.Println("--> Evaluate Transaction: AssetExists, function returns 'true' if an asset with given assetID exist")
	// result, err = contract.EvaluateTransaction("AssetExists", "asset1")
	// if err != nil {
	// 	log.Fatalf("Failed to evaluate transaction: %v\n", err)
	// }
	// log.Println(string(result))

	// log.Println("--> Submit Transaction: TransferAsset asset1, transfer to new owner of Tom")
	// _, err = contract.SubmitTransaction("TransferAsset", "asset1", "Tom")
	// if err != nil {
	// 	log.Fatalf("Failed to Submit transaction: %v", err)
	// }

	// log.Println("--> Evaluate Transaction: ReadAsset, function returns 'asset1' attributes")
	// result, err = contract.EvaluateTransaction("ReadAsset", "asset1")
	// if err != nil {
	// 	log.Fatalf("Failed to evaluate transaction: %v", err)
	// }
	// log.Println(string(result))
	log.Println("============ application-golang ends ============")

	r := gin.Default()

	r.GET("/ReadAsset", func(c *gin.Context) {
		log.Println("--> Evaluate Transaction: ReadAsset, function returns an asset with a given assetID")
		assetID := c.Query("assetID")
		result, err := contract.EvaluateTransaction("ReadAsset", assetID)
		if err != nil {
			log.Println("Failed to evaluate transaction: %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		log.Println(string(result))

		c.JSON(http.StatusOK, gin.H{
			"message": string(result),
		})
	})

	r.GET("/GetAllAssets", func(c *gin.Context) {
		log.Println("--> Evaluate Transaction: GetAllAssets, function returns all the current assets on the ledger")
		result, err := contract.EvaluateTransaction("GetAllAssets")
		if err != nil {
			log.Println("Failed to evaluate transaction: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		log.Println(string(result))

		c.JSON(http.StatusOK, gin.H{
			"message": string(result),
		})
	})

	r.POST("/CreateAsset", func(c *gin.Context) {

		var json Asset
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		log.Println("--> Submit Transaction: CreateAsset, creates new asset with", json)

		result, err := contract.SubmitTransaction("CreateAsset", json.AssetID, json.Color, json.Size, json.Owner, json.AppraisedValue)

		if err != nil {
			log.Println("Failed to Submit transaction: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		log.Println(string(result))
		c.JSON(http.StatusOK, gin.H{
			"message": string(result),
		})
	})

	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

func populateWallet(wallet *gateway.Wallet) error {
	log.Println("============ Populating wallet ============")
	credPath := filepath.Join(
		"..",
		"..",
		"test-network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"users",
		"User1@org1.example.com",
		"msp",
	)

	certPath := filepath.Join(credPath, "signcerts", "cert.pem")
	// read the certificate pem
	cert, err := ioutil.ReadFile(filepath.Clean(certPath))
	if err != nil {
		return err
	}

	keyDir := filepath.Join(credPath, "keystore")
	// there's a single file in this dir containing the private key
	files, err := ioutil.ReadDir(keyDir)
	if err != nil {
		return err
	}
	if len(files) != 1 {
		return fmt.Errorf("keystore folder should have contain one file")
	}
	keyPath := filepath.Join(keyDir, files[0].Name())
	key, err := ioutil.ReadFile(filepath.Clean(keyPath))
	if err != nil {
		return err
	}

	identity := gateway.NewX509Identity("Org1MSP", string(cert), string(key))

	return wallet.Put("appUser", identity)
}
