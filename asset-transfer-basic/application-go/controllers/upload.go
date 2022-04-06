package controllers

import (
	"asset-transfer-basic/models"
	"asset-transfer-basic/utils"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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
		inputInfo := models.InputInfo{CertDetail: localAssetItem, PersonInfoHash: s[1]}

		inputInfoArray = append(inputInfoArray, inputInfo)
	}

	batchSize := utils.MAX_BATCH_SIZE_FOR_MKTREE
	batches := make([][]models.InputInfo, 0, (len(inputInfoArray)+batchSize-1)/batchSize)
	for batchSize < len(inputInfoArray) {
		inputInfoArray, batches = inputInfoArray[batchSize:], append(batches, inputInfoArray[0:batchSize:batchSize])
	}
	batches = append(batches, inputInfoArray)

	for _, s := range batches {
		dailyRecord := s

		globalID := uuid.New().String()

		merkleTree, err := utils.GetMerkleTree(dailyRecord, globalID)

		if err != nil {
			log.Printf("Failed to Create Merkle Tree: %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		merkleTreeRoot := merkleTree.MerkleRoot()

		merkleTreeRootStr := base64.StdEncoding.EncodeToString(merkleTreeRoot)
		log.Printf("--> Evaluate Transaction: Upload to global chain")

		var certIDList []string
		for i := range dailyRecord {
			certIDList = append(certIDList, dailyRecord[i].CertDetail.CertNo)
		}
		certIDListJson, err := json.Marshal(certIDList)

		if err != nil {
			log.Printf("Failed to Submit transaction: %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var info = models.GlobalChainInfo{
			CertIDList:           string(certIDListJson),
			MerkleTreeRoot:       merkleTreeRootStr,
			GlobalChainTxHash:    "",
			GlobalChainBlockNum:  1,
			GlobalChainTimeStamp: utils.GetUnixTime()}

		result, err := GlobalContract.SubmitTransaction("CreateAsset",
			globalID,
			strconv.Itoa(int(info.GlobalChainBlockNum)),
			info.MerkleTreeRoot,
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
