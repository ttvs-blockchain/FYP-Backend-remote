
# curl -v -X GET "http://localhost:8080/GetAllAssets"

# curl -v -X POST http://localhost:8080/CreateAsset   -H 'content-type: application/json' -d '{"CertDetail":{ "CertNo": "a1", "ID":"IDTEST", "Name":"SinoVac", "Brand": "TestBrand", "NumOfDose":"3","Time":"2022/02/22", "Issuer":"issuertest", "Remark":"no remark"}, "PersonHash":"personhashsample"}'
# curl -v -X POST http://localhost:8080/CreateAsset   -H 'content-type: application/json' -d '{"CertDetail":{ "CertNo": "a2", "ID":"IDTEST", "Name":"SinoVac", "Brand": "TestBrand", "NumOfDose":"3","Time":"2022/02/22", "Issuer":"issuertest", "Remark":"no remark"}, "PersonHash":"personhashsample"}'
# curl -v -X POST http://localhost:8080/CreateAsset   -H 'content-type: application/json' -d '{"CertDetail":{ "CertNo": "a3", "ID":"IDTEST", "Name":"SinoVac", "Brand": "TestBrand", "NumOfDose":"3","Time":"2022/02/22", "Issuer":"issuertest", "Remark":"no remark"}, "PersonHash":"personhashsample"}'
# curl -v -X POST http://localhost:8080/CreateAsset   -H 'content-type: application/json' -d '{"CertDetail":{ "CertNo": "a4", "ID":"IDTEST", "Name":"SinoVac", "Brand": "TestBrand", "NumOfDose":"3","Time":"2022/02/22", "Issuer":"issuertest", "Remark":"no remark"}, "PersonHash":"personhashsample"}'
# curl -v -X POST http://localhost:8080/CreateAsset   -H 'content-type: application/json' -d '{"CertDetail":{ "CertNo": "a5", "ID":"IDTEST", "Name":"SinoVac", "Brand": "TestBrand", "NumOfDose":"3","Time":"2022/02/22", "Issuer":"issuertest", "Remark":"no remark"}, "PersonHash":"personhashsample"}'
# curl -v -X POST http://localhost:8080/CreateAsset   -H 'content-type: application/json' -d '{"CertDetail":{ "CertNo": "a6", "ID":"IDTEST", "Name":"SinoVac", "Brand": "TestBrand", "NumOfDose":"3","Time":"2022/02/22", "Issuer":"issuertest", "Remark":"no remark"}, "PersonHash":"personhashsample"}'
# curl -v -X POST http://localhost:8080/CreateAsset   -H 'content-type: application/json' -d '{"CertDetail":{ "CertNo": "a7", "ID":"IDTEST", "Name":"SinoVac", "Brand": "TestBrand", "NumOfDose":"3","Time":"2022/02/22", "Issuer":"issuertest", "Remark":"no remark"}, "PersonHash":"personhashsample"}'
curl -v -X POST http://localhost:8080/CreateAsset   -H 'content-type: application/json' -d '{"CertDetail":{ "CertNo": "a11", "ID":"IDTEST", "Name":"SinoVac", "Brand": "TestBrand", "NumOfDose":"3","Time":"2022/02/22", "Issuer":"issuertest", "Remark":"no remark"}, "PersonHash":"personhashsample"}'

# curl -v -X GET "http://localhost:8080/ReadAsset?CertNo=a1"

# curl -v -X GET "http://localhost:8080/GetAllAssets"

curl -v -X POST http://localhost:8080/Upload