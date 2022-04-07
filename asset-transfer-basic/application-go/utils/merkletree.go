package utils

import (
	"asset-transfer-basic/models"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"

	mt "github.com/cbergoon/merkletree"
)

//TestContent implements the Content interface provided by merkle tree and represents the content stored in the tree.
type TestContent struct {
	x string
}

//CalculateHash hashes the values of a TestContent
func (t TestContent) CalculateHash() ([]byte, error) {
	h := sha256.New()
	fmt.Printf("-->Debug get string before hash %s\n", t.x)
	if _, err := h.Write([]byte(t.x)); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

//Equals tests for equality of two Contents
func (t TestContent) Equals(other mt.Content) (bool, error) {
	return t.x == other.(TestContent).x, nil
}

func ConvertTreeContent(dailyRecord []models.InputInfo) ([]mt.Content, error) {
	var list []mt.Content

	for i, s := range dailyRecord {
		fmt.Println(i, s)
		jsonS, err := json.Marshal(s)
		if err != nil {
			fmt.Println("failed when converting json")
			return nil, err
		}
		list = append(list, TestContent{x: string(jsonS)})
	}
	return list, nil

}

func GetMerkleTree(list []mt.Content) (*mt.MerkleTree, error) {

	//Create a new Merkle Tree from the list of Content
	t, err := mt.NewTree(list)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return t, nil

}

func StoreMerklePath(list []mt.Content, t *mt.MerkleTree, globalID string, dailyRecord []models.InputInfo) error {
	for i, s := range list {
		// fmt.Println(i, s)
		path, indexes, err := t.GetMerklePath(s)
		if err != nil {
			fmt.Println("failed when getting path ")
			return err
		}
		if err != nil {
			fmt.Println("failed when getting currentHash ")
			return err
		}

		var resultPath []string
		for _, s := range path {
			resultPath = append(resultPath, base64.StdEncoding.EncodeToString(s))
		}
		var merkleTreePath = models.MerkleTreePath{
			GlobalRootID: globalID,
			Path:         resultPath,
			Indexes:      indexes}

		info := models.LocalChainInfo{
			LocalChainID:         LOCAL_CHAIN_ID,
			MerkleTreePathDetail: merkleTreePath,
			LocalChainTxHash:     "",
			LocalChainBlockNum:   1,
			LocalChainTimeStamp:  GetUnixTime()}

		err = models.UpdateLocalCertDB(info, dailyRecord[i].CertDetail.CertID)
		if err != nil {
			return err
		}
	}

	return nil
}

func reverseArray(arr [][]byte) [][]byte {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func getHash(a []byte, b []byte) []byte {

	h := sha256.New()
	fmt.Printf("the input is %s,    %s\n",
		base64.StdEncoding.EncodeToString(a),
		base64.StdEncoding.EncodeToString(b))

	if _, err := h.Write(append(a, b...)); err != nil {
		// return nil, err
		fmt.Printf("GG")
	}

	fmt.Printf("the out is %s\n",
		base64.StdEncoding.EncodeToString(h.Sum(nil)))
	return h.Sum(nil)
}

func CreateBatches(inputInfos []models.InputInfo) [][]models.InputInfo {
	batchSize := MAX_BATCH_SIZE_FOR_MKTREE
	batches := make([][]models.InputInfo, 0, (len(inputInfos)+batchSize-1)/batchSize)
	for batchSize < len(inputInfos) {
		inputInfos, batches = inputInfos[batchSize:], append(batches, inputInfos[0:batchSize:batchSize])
	}
	batches = append(batches, inputInfos)
	return batches
}
