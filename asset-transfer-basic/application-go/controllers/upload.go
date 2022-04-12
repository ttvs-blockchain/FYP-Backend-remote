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
		inputInfo := models.InputInfo{CertDetail: localAssetItem, PersonInfoHash: s[1], Key: s[2]}

		inputInfoArray = append(inputInfoArray, inputInfo)
	}

	batches := utils.CreateBatches(inputInfoArray)

	var infoArray []models.GlobalChainInfo
	for _, batch := range batches {

		globalID := uuid.New().String()

		list, err := utils.ConvertTreeContent(batch)
		if err != nil {
			log.Printf("Failed to ConvertTreeContent: %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		merkleTree, err := utils.GetMerkleTree(list)
		if err != nil {
			log.Printf("Failed to GetMerkleTree: %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = utils.StoreMerklePath(list, merkleTree, globalID, batch)
		if err != nil {
			log.Printf("Failed to StoreMerklePath: %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		merkleTreeRoot := base64.StdEncoding.EncodeToString(merkleTree.MerkleRoot())
		log.Printf("--> Evaluate Transaction: Upload to global chain")

		var batchCertIDList []string
		for i := range batch {
			batchCertIDList = append(batchCertIDList, batch[i].CertDetail.CertID)
		}
		batchCertIDListJson, err := json.Marshal(batchCertIDList)
		if err != nil {
			log.Printf("Failed to Construct cert ID list: %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var info = models.GlobalChainInfo{
			CertIDList:           string(batchCertIDListJson),
			MerkleTreeRoot:       merkleTreeRoot,
			GlobalChainTxHash:    "",
			GlobalChainBlockNum:  1,
			GlobalChainTimeStamp: utils.GetUnixTime()}

		result, err := GlobalContract.SubmitTransaction("CreateAsset",
			globalID,
			strconv.Itoa(int(info.GlobalChainBlockNum)),
			info.MerkleTreeRoot,
		)

		if err != nil {
			log.Printf("Failed to Submit transaction: %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Printf("-->Evaluate Transaction: result of CreateAsset is %s\n", result)
		infoArray = append(infoArray, info)
	}

	log.Printf("--> Start Insert Multiple Rows to DB")

	err = models.InsertMultipleToGlobalHashDB(infoArray)
	if err != nil {
		log.Printf("Failed to Insert Row in DB for transaction: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("--> Finish Transaction: Upload\n")

	c.JSON(http.StatusOK, gin.H{
		"message": "finish upload",
	})
}
