curl -v -X POST http://localhost:8080/CreateAsset   -H 'content-type: application/json' -d '{ "CertNo": "ggfg1g1gg34dd", "ID":"IDTEST", "Name":"jerry", "Brand": "TestBrand", "NumOfDose":"3","Time":"2022/02/22", "Issuer":"issuertest", "Remark":"no remark"}'


curl -v -X POST http://localhost:8080/CreateAsset   -H 'content-type: application/json' -d '{ "CertNo": "fg1g2gggd", "ID":"IDTEST", "Name":"jerry", "Brand": "TestBrand", "NumOfDose":"3","Time":"2022/02/22", "Issuer":"issuertest", "Remark":"no remark"}'

curl -v -X POST http://localhost:8080/CreateAsset   -H 'content-type: application/json' -d '{ "CertNo": "fg1g34ff5dd", "ID":"IDTEST", "Name":"jerry", "Brand": "TestBrand", "NumOfDose":"3","Time":"2022/02/22", "Issuer":"issuertest", "Remark":"no remark"}'
curl -v -X POST http://localhost:8080/CreateAsset   -H 'content-type: application/json' -d '{ "CertNo": "fg1g46ffdd", "ID":"IDTEST", "Name":"jerry", "Brand": "TestBrand", "NumOfDose":"3","Time":"2022/02/22", "Issuer":"issuertest", "Remark":"no remark"}'
curl -v -X POST http://localhost:8080/CreateAsset   -H 'content-type: application/json' -d '{ "CertNo": "fg1g5ffdd", "ID":"IDTEST", "Name":"jerry", "Brand": "TestBrand", "NumOfDose":"3","Time":"2022/02/22", "Issuer":"issuertest", "Remark":"no remark"}'
curl -v -X POST http://localhost:8080/CreateAsset   -H 'content-type: application/json' -d '{ "CertNo": "fg1g6ff98dd", "ID":"IDTEST", "Name":"jerry", "Brand": "TestBrand", "NumOfDose":"3","Time":"2022/02/22", "Issuer":"issuertest", "Remark":"no remark"}'
curl -v -X POST http://localhost:8080/CreateAsset   -H 'content-type: application/json' -d '{ "CertNo": "fg1g7234fdd", "ID":"IDTEST", "Name":"jerry", "Brand": "TestBrand", "NumOfDose":"3","Time":"2022/02/22", "Issuer":"issuertest", "Remark":"no remark"}'
curl -v -X POST http://localhost:8080/CreateAsset   -H 'content-type: application/json' -d '{ "CertNo": "fg1g8443fdd", "ID":"IDTEST", "Name":"jerry", "Brand": "TestBrand", "NumOfDose":"3","Time":"2022/02/22", "Issuer":"issuertest", "Remark":"no remark"}'

curl -v -X GET "http://localhost:8080/ReadAsset?ID=1g8443fdd"

curl -v -X GET "http://localhost:8080/GetAllAssets"

curl -v -X POST http://localhost:8080/Upload