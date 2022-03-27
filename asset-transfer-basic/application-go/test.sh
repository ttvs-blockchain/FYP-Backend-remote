curl -v -X POST http://localhost:8080/CreateAsset   -H 'content-type: application/json' -d '{ "CertNo": "45465465", "ID":"IDTEST", "Name":"jerry", "Brand": "TestBrand", "NumOfDose":"3","Time":"2022/02/22", "Issuer":"issuertest", "Remark":"no remark"}'

# curl -v -X GET "http://localhost:8080/ReadAsset?assetID=123"

curl -v -X GET "http://localhost:8080/GetAllAssets"

curl -v -X POST http://localhost:8080/Upload