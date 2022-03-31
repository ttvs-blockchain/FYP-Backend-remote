package utils

import (
	"asset-transfer-basic/models"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"

	"github.com/cbergoon/merkletree"
)

//TestContent implements the Content interface provided by merkletree and represents the content stored in the tree.
type TestContent struct {
	x string
}

//CalculateHash hashes the values of a TestContent
func (t TestContent) CalculateHash() ([]byte, error) {
	h := sha256.New()
	if _, err := h.Write([]byte(t.x)); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

//Equals tests for equality of two Contents
func (t TestContent) Equals(other merkletree.Content) (bool, error) {
	return t.x == other.(TestContent).x, nil
}

func GetMerkelTree(dailyRecord []models.Asset, globalID string) (*merkletree.MerkleTree, error) {
	var list []merkletree.Content

	for i, s := range dailyRecord {
		fmt.Println(i, s)
		jsonS, err := json.Marshal(s)
		if err != nil {
			fmt.Println("failed when converting json")
			return nil, err
		}
		list = append(list, TestContent{x: string(jsonS)})
	}

	//Create a new Merkle Tree from the list of Content
	t, err := merkletree.NewTree(list)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for i, s := range list {
		// fmt.Println(i, s)
		path, indexes, err := t.GetMerklePath(s)
		if err != nil {
			fmt.Println("failed when getting path ")
			return nil, err
		}
		currentHash, err := s.CalculateHash()
		if err != nil {
			fmt.Println("failed when getting currentHash ")
			return nil, err
		}

		var resultPath []string
		for _, s := range path {
			resultPath = append(resultPath, base64.StdEncoding.EncodeToString(s))
		}
		currentHashStr := base64.StdEncoding.EncodeToString(currentHash)
		var merkelTreePath = models.MerkelTreePath{
			globalID,
			currentHashStr,
			resultPath,
			indexes}

		resultPathJson, err := json.Marshal(merkelTreePath)
		if err != nil {
			fmt.Println("failed when getting path in json")
			return nil, err
		}

		info := models.LocalChainInfo{
			LOCAL_CHAIN_ID,
			string(resultPathJson),
			1,
			GetUnixTime()}

		err = models.UpdateRow(info, dailyRecord[i].CertNo)
		if err != nil {
			return nil, err
		}

	}

	fmt.Println("*************get path end")

	return t, nil

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
