[//]: # (SPDX-License-Identifier: CC-BY-4.0)

# Vaxpass - A Two-tier Permission Blockchain for COVID-19 Vaccine Certificate Verification 

DIP1

Final Year Project 2022
## File Architecture

The smart contract of  [local chain](/asset-transfer-basic/chaincode-go) and [global chain](/asset-transfer-basic/chaincode-go-global) are developed based on (asset-transfer-basic)[https://github.com/hyperledger/fabric-samples/tree/main/asset-transfer-basic] smart contract from fabric-samples of Hyperledger Fabric.

The [backend for local chain](/asset-transfer-basic/application-go) and [backend for global chain]/(asset-transfer-basic/application-go-global) are based on application from (asset-transfer-basic)[https://github.com/hyperledger/fabric-samples/tree/main/asset-transfer-basic].

The folder of [tlsCert at local](/asset-transfer-basic/application-go/tlsCert/) and [tlsCert at global](/asset-transfer-basic/application-go-global/tlsCert/) stores a self-signed certificates for testing purpose and is gitignored.

## Getting started 
To install all executables, please refers to the Hyperledger Fabric Official Website [test network tutorial](https://hyperledger-fabric.readthedocs.io/en/latest/test_network.html) to install.

To launch two peers and deploy local chain and global chain on these peers, run the following

```
cd test-network
./deploy.sh
```

## Test network

The [Fabric test network](test-network) in the samples repository provides a Docker Compose based test network with two
Organization peers and an ordering service node. You can use it on your local machine to run the samples listed below.
You can also use it to deploy and test your own Fabric chaincodes and applications. To get started, see
the 

## Asset transfer samples and tutorials

The asset transfer series provides a series of sample smart contracts and applications to demonstrate how to store and transfer assets using Hyperledger Fabric.
Each sample and associated tutorial in the series demonstrates a different core capability in Hyperledger Fabric. The **Basic** sample provides an introduction on how
to write smart contracts and how to interact with a Fabric network using the Fabric SDKs. The **Ledger queries**, **Private data**, and **State-based endorsement**
samples demonstrate these additional capabilities. Finally, the **Secured agreement** sample demonstrates how to bring all the capabilities together to securely
transfer an asset in a more realistic transfer scenario.


## License <a name="license"></a>

Hyperledger Project source code files are made available under the Apache
License, Version 2.0 (Apache-2.0), located in the [LICENSE](LICENSE) file.
Hyperledger Project documentation files are made available under the Creative
Commons Attribution 4.0 International License (CC-BY-4.0), available at http://creativecommons.org/licenses/by/4.0/.
