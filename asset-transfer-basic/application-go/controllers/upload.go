package controllers

import (
	"asset-transfer-basic/models"
	"asset-transfer-basic/utils"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Upload(c *gin.Context) {

	log.Printf("--> Evaluate Transaction: ReadRowForMKTree in db")

	idList, err := models.ReadRowForMKTree()
	if err != nil {
		log.Printf("Failed to find on local db: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(idList) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "no new certificate for global chain",
		})
		return
	}

	fmt.Printf("--> Evaluate Transaction: Get id from db for upload, %s\n", idList)
	var inputInfoArray []models.InputInfo
	for _, s := range idList {
		localAsset, err := Contract.EvaluateTransaction("ReadAsset", s[0])
		if err != nil {
			log.Printf("Failed to evaluate transaction: %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var localAssetItem models.Asset
		err = json.Unmarshal(localAsset, &localAssetItem)

		if err != nil {
			log.Printf("Failed to evaluate transaction: %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		inputInfo := models.InputInfo{localAssetItem, s[1]}

		inputInfoArray = append(inputInfoArray, inputInfo)
	}

	// localRecord, err := json.Marshal(InputInfoArray)

	// if err != nil {
	// 	log.Printf("Failed to evaluate transaction: %v\n", err)
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// var inputInfoArray []models.Asset
	// err = json.Unmarshal(localRecord, &inputInfoArray)
	// if err != nil {
	// 	log.Printf("Failed to evaluate transaction: %v\n", err)
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	// inputInfoArray := inputInfoArray

	batchSize := utils.MAX_BATCH_SIZE_FOR_MKTREE
	batches := make([][]models.InputInfo, 0, (len(inputInfoArray)+batchSize-1)/batchSize)
	for batchSize < len(inputInfoArray) {
		inputInfoArray, batches = inputInfoArray[batchSize:], append(batches, inputInfoArray[0:batchSize:batchSize])
	}
	batches = append(batches, inputInfoArray)

	for _, s := range batches {
		dailyRecord := s

		globalID := uuid.New().String()

		merkelTree, err := utils.GetMerkelTree(dailyRecord, globalID)

		if err != nil {
			log.Printf("Failed to Create Merkel Tree: %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		merkelTreeRoot := merkelTree.MerkleRoot()

		merkelTreeRootStr := base64.StdEncoding.EncodeToString(merkelTreeRoot)
		log.Printf("--> Evaluate Transaction: Upload to global chain")

		var certIDList []string
		for i, _ := range dailyRecord {
			certIDList = append(certIDList, dailyRecord[i].CertDetail.CertNo)
		}
		certIDListJson, err := json.Marshal(certIDList)

		if err != nil {
			log.Printf("Failed to Submit transaction: %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var info = models.GlocalChainInfo{
			string(certIDListJson),
			merkelTreeRootStr,
			"",
			1,
			utils.GetUnixTime()}

		result, err := GlobalContract.SubmitTransaction("CreateAsset",
			globalID,
			string(info.GlobalChainBlockNum),
			info.MerkelTreeRoot,
		)
		fmt.Printf("-->Evaluate Transaction: result of CreateAsset is %s\n", result)
		if err != nil {
			log.Printf("Failed to Submit transaction: %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = models.InsertGlobalHashDB(info)

		if err != nil {
			log.Printf("Failed to Insert Row in DB for transaction: %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	log.Printf("--> Finish Transaction: Upload\n")

	c.JSON(http.StatusOK, gin.H{
		"message": "finish upload",
	})
}
