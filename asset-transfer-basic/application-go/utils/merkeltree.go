package utils

import (
	"asset-transfer-basic/models"
	"crypto/sha256"
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

func GetMerkelTreeRoot(localRecord []byte) (string, error) {
	var dailyRecord []models.Asset
	var list []merkletree.Content
	err := json.Unmarshal(localRecord, &dailyRecord)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	for i, s := range dailyRecord {
		fmt.Println(i, s)
		jsonS, err := json.Marshal(s)
		if err != nil {
			fmt.Println("failed when converting json")
			return "", err
		}

		list = append(list, TestContent{x: string(jsonS)})
	}

	//Create a new Merkle Tree from the list of Content
	t, err := merkletree.NewTree(list)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	//Get the Merkle Root of the tree
	mr := t.MerkleRoot()

	return string(mr), nil
}
