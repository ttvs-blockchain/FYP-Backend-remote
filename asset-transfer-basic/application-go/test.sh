curl -v -X POST http://localhost:8080/CreateAsset   -H 'content-type: application/json' -d '{ "CertNo": "a1", "ID":"IDTEST", "Name":"jerry", "Brand": "TestBrand", "NumOfDose":"3","Time":"2022/02/22", "Issuer":"issuertest", "Remark":"no remark"}'


curl -v -X POST http://localhost:8080/CreateAsset   -H 'content-type: application/json' -d '{ "CertNo": "a2", "ID":"IDTEST", "Name":"jerry", "Brand": "TestBrand", "NumOfDose":"3","Time":"2022/02/22", "Issuer":"issuertest", "Remark":"no remark"}'

curl -v -X POST http://localhost:8080/CreateAsset   -H 'content-type: application/json' -d '{ "CertNo": "a3", "ID":"IDTEST", "Name":"jerry", "Brand": "TestBrand", "NumOfDose":"3","Time":"2022/02/22", "Issuer":"issuertest", "Remark":"no remark"}'
curl -v -X POST http://localhost:8080/CreateAsset   -H 'content-type: application/json' -d '{ "CertNo": "a4", "ID":"IDTEST", "Name":"jerry", "Brand": "TestBrand", "NumOfDose":"3","Time":"2022/02/22", "Issuer":"issuertest", "Remark":"no remark"}'



curl -v -X POST http://localhost:8080/CreateAsset   -H 'content-type: application/json' -d '{ "CertNo": "a5", "ID":"IDTEST", "Name":"jerry", "Brand": "TestBrand", "NumOfDose":"3","Time":"2022/02/22", "Issuer":"issuertest", "Remark":"no remark"}'
curl -v -X POST http://localhost:8080/CreateAsset   -H 'content-type: application/json' -d '{ "CertNo": "a6", "ID":"IDTEST", "Name":"jerry", "Brand": "TestBrand", "NumOfDose":"3","Time":"2022/02/22", "Issuer":"issuertest", "Remark":"no remark"}'


curl -v -X POST http://localhost:8080/CreateAsset   -H 'content-type: application/json' -d '{ "CertNo": "a7", "ID":"IDTEST", "Name":"jerry", "Brand": "TestBrand", "NumOfDose":"3","Time":"2022/02/22", "Issuer":"issuertest", "Remark":"no remark"}'
curl -v -X POST http://localhost:8080/CreateAsset   -H 'content-type: application/json' -d '{ "CertNo": "a8", "ID":"IDTEST", "Name":"jerry", "Brand": "TestBrand", "NumOfDose":"3","Time":"2022/02/22", "Issuer":"issuertest", "Remark":"no remark"}'

curl -v -X POST http://localhost:8080/Upload

# curl -v -X GET "http://localhost:8080/ReadAsset?ID=1g8443fdd"

curl -v -X GET "http://localhost:8080/GetAllAssets"

