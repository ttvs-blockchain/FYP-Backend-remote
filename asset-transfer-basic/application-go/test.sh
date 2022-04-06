
# curl -v -X GET "http://localhost:8080/GetAllAssets"


curl -v -X POST http://localhost:8080/CreateAsset   -H 'content-type: application/json' -d '{"CertDetail":{"CertNo": "TEST20220413_01", "ID": "M123456(0)", "Name":"CoronaVac" , "Brand": "SinoVac" ,   "NumOfDose":"2","Time":"2022-04-03T15:50", "Issuer":"vac_center1" } , "PersonInfoHash":"047d1be86e0ed70abfd55540a1d6fe09c91c74eb99f82d4f77926b721a87b0fe"}'


curl -v -X POST http://localhost:8080/CreateAsset   -H 'content-type: application/json' -d '{"CertDetail":{"CertNo": "TEST20220413_02", "ID": "M123456(0)", "Name":"CoronaVac" , "Brand": "SinoVac" ,   "NumOfDose":"2","Time":"2022-04-03T15:50", "Issuer":"vac_center1" } , "PersonInfoHash":"047d1be86e0ed70abfd55540a1d6fe09c91c74eb99f82d4f77926b721a87b0fe"}'
curl -v -X POST http://localhost:8080/CreateAsset   -H 'content-type: application/json' -d '{"CertDetail":{"CertNo": "TEST20220413_03", "ID": "M123456(0)", "Name":"CoronaVac" , "Brand": "SinoVac" ,   "NumOfDose":"2","Time":"2022-04-03T15:50", "Issuer":"vac_center1" } , "PersonInfoHash":"047d1be86e0ed70abfd55540a1d6fe09c91c74eb99f82d4f77926b721a87b0fe"}'
curl -v -X POST http://localhost:8080/CreateAsset   -H 'content-type: application/json' -d '{"CertDetail":{"CertNo": "TEST20220413_04", "ID": "M123456(0)", "Name":"CoronaVac" , "Brand": "SinoVac" ,   "NumOfDose":"2","Time":"2022-04-03T15:50", "Issuer":"vac_center1" } , "PersonInfoHash":"047d1be86e0ed70abfd55540a1d6fe09c91c74eb99f82d4f77926b721a87b0fe"}'
curl -v -X POST http://localhost:8080/CreateAsset   -H 'content-type: application/json' -d '{"CertDetail":{"CertNo": "TEST20220413_05", "ID": "M123456(0)", "Name":"CoronaVac" , "Brand": "SinoVac" ,   "NumOfDose":"2","Time":"2022-04-03T15:50", "Issuer":"vac_center1" } , "PersonInfoHash":"047d1be86e0ed70abfd55540a1d6fe09c91c74eb99f82d4f77926b721a87b0fe"}'


curl -v -X GET "http://localhost:8080/GetAllAssets"

curl -v -X POST http://localhost:8080/Upload