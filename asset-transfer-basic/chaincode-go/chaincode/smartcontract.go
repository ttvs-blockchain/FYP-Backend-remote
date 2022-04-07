package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

// Asset describes basic details of what makes up a simple asset
type Asset struct {
	CertID      string `form:"certID" json:"certID" xml:"certID"  binding:"required"`
	PersonSysID string `form:"personSysID" json:"personSysID" xml:"personSysID"  binding:"required"`
	Name        string `form:"name" json:"name" xml:"name"  binding:"required"`
	Brand       string `form:"brand" json:"brand" xml:"brand"  binding:"required"`
	NumOfDose   string `form:"numOfDose" json:"numOfDose" xml:"numOfDose"  binding:"required"`
	Time        string `form:"time" json:"time" xml:"time"  binding:"required"`
	Issuer      string `form:"issuer" json:"issuer" xml:"issuer"  binding:"required"`
	Remark      string `form:"remark" json:"remark" xml:"remark"  binding:""`
}

// InitLedger adds a base set of assets to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	assets := []Asset{
		{
			CertID:      "testCertNo",
			PersonSysID: "testID",
			Name:        "Alice",
			Brand:       "TestBrand",
			NumOfDose:   "1",
			Time:        "2021-12-21",
			Issuer:      "test Issuer",
			Remark:      ""},
	}

	for _, asset := range assets {
		assetJSON, err := json.Marshal(asset)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(asset.CertID, assetJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state. %v", err)
		}
	}

	return nil
}

// CreateAsset issues a new asset to the world state with given details.
func (s *SmartContract) CreateAsset(ctx contractapi.TransactionContextInterface,
	certID string,
	personSysID string,
	name string,
	brand string,
	numOfDose string,
	time string,
	issuer string,
	remark string) error {
	exists, err := s.AssetExists(ctx, certID)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the asset %s already exists", certID)
	}

	asset := Asset{
		CertID:      certID,
		PersonSysID: personSysID,
		Name:        name,
		Brand:       brand,
		NumOfDose:   numOfDose,
		Time:        time,
		Issuer:      issuer,
		Remark:      remark}

	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(certID, assetJSON)
}

// ReadAsset returns the asset stored in the world state with given certID.
func (s *SmartContract) ReadAsset(ctx contractapi.TransactionContextInterface, certID string) (*Asset, error) {
	assetJSON, err := ctx.GetStub().GetState(certID)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("the asset %s does not exist", certID)
	}

	var asset Asset
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, err
	}

	return &asset, nil
}

// UpdateAsset updates an existing asset in the world state with provided parameters.
func (s *SmartContract) UpdateAsset(ctx contractapi.TransactionContextInterface,
	certID string,
	personSysID string,
	name string,
	brand string,
	numOfDose string,
	time string,
	issuer string,
	remark string) error {
	exists, err := s.AssetExists(ctx, certID)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", certID)
	}

	// overwriting original asset with new asset
	asset := Asset{
		CertID:      certID,
		PersonSysID: personSysID,
		Name:        name,
		Brand:       brand,
		NumOfDose:   numOfDose,
		Time:        time,
		Issuer:      issuer,
		Remark:      remark}
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(certID, assetJSON)
}

// DeleteAsset deletes an given asset from the world state.
func (s *SmartContract) DeleteAsset(ctx contractapi.TransactionContextInterface, certID string) error {
	exists, err := s.AssetExists(ctx, certID)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", certID)
	}

	return ctx.GetStub().DelState(certID)
}

// AssetExists returns true when asset with given ID exists in world state
func (s *SmartContract) AssetExists(ctx contractapi.TransactionContextInterface, certID string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(certID)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return assetJSON != nil, nil
}

// GetAllAssets returns all assets found in world state
func (s *SmartContract) GetAllAssets(ctx contractapi.TransactionContextInterface) ([]*Asset, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all assets in the chaincode namespace.
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var assets []*Asset
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset Asset
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}

	return assets, nil
}
