# curl -v -X POST http://localhost:8080/CreateAsset   -H 'content-type: application/json' -d '{ "assetID": "123", "color":"red", "owner":"jerry", "size": "3", "appraisedValue":"3"}'

# curl -v -X GET "http://localhost:8080/ReadAsset?assetID=123"

curl -v -X GET "http://localhost:8080/GetAllAssets"