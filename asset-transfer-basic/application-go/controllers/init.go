package controllers

import (
	"asset-transfer-basic/utils"
	"log"
	"os"
	"path/filepath"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

var Contract *gateway.Contract
var GlobalContract *gateway.Contract

func init() {
	err := os.Setenv("DISCOVERY_AS_LOCALHOST", "true")
	if err != nil {
		log.Fatalf("Error setting DISCOVERY_AS_LOCALHOST environment variable: %v\n", err)
	}

	wallet, err := gateway.NewFileSystemWallet("wallet")
	if err != nil {
		log.Fatalf("Failed to create wallet: %v\n", err)
	}

	if !wallet.Exists("appUser") {
		err = utils.PopulateWallet(wallet)
		if err != nil {
			log.Fatalf("Failed to populate wallet contents: %v\n", err)
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
		log.Fatalf("Failed to connect to gateway: %v\n", err)
	}

	defer gw.Close()

	network, err := gw.GetNetwork("mychannel")
	if err != nil {
		log.Fatalf("Failed to get network: %v\n", err)
	}

	globalNetwork, err := gw.GetNetwork("globalchannel")
	if err != nil {
		log.Fatalf("Failed to get global network: %v\n", err)
	}
	Contract = network.GetContract("basic")
	GlobalContract = globalNetwork.GetContract("basic-global")
}
