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

	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
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

var db *sql.DB

func initDB() (err error) {
	// DSN:Data Source Name
	dsn := "tommy:mysql123456@tcp(167.179.77.244:3306)/certificate?charset=utf8mb4&parseTime=True"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func insertRow(asset Asset) {
	sqlStr := "insert into certificate(certno, personid, name, brand, numofdose, time, issuer, remark) values (?,?,?,?,?,?,?,?)"
	ret, err := db.Exec(sqlStr, asset.CertNo, asset.ID, asset.Name, asset.Brand, asset.NumOfDose, asset.Time, asset.Issuer, asset.Remark)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

func main() {

	err := initDB() //
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}

	log.Println("============ application-golang starts ============")

	err = os.Setenv("DISCOVERY_AS_LOCALHOST", "true")
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

	log.Println("============ application-golang ends ============")

	r := gin.Default()

	r.GET("/ReadAsset", func(c *gin.Context) {
		log.Println("--> Evaluate Transaction: ReadAsset, function returns an asset with a given assetID")
		CertNo := c.Query("CertNo")
		result, err := contract.EvaluateTransaction("ReadAsset", CertNo)
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

		result, err := contract.SubmitTransaction("CreateAsset",
			json.CertNo, json.ID, json.Name, json.Brand, json.NumOfDose, json.Time, json.Issuer, json.Remark)

		if err != nil {
			log.Println("Failed to Submit transaction: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		log.Println(string(result))

		insertRow(json)

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

// func getNetworkForOrgUser(userId string) {
// 	wallet = getWa
// }
