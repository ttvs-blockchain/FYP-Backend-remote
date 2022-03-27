# curl -v -X POST http://localhost:8081/CreateAsset   -H 'content-type: application/json' -d '{ "LocalChainID": "LocalChainIDtest", "GlobalChainTxHash":"hashtest"}'

# curl -v -X GET "http://localhost:8080/ReadAsset?assetID=123"

curl -v -X GET "http://localhost:8081/GetAllAssets"