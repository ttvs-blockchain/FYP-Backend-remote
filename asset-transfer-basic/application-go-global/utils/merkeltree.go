package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"

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
