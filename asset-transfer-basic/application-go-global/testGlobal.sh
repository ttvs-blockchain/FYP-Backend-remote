# curl -v -X GET "http://localhost:8081/GetAllAssets"

# curl -v -X POST http://localhost:8081/Verify -H 'content-type: application/json' -d ' {"GlobalID":"0db1a2ef-7dc3-426d-950a-c710ee603f31","CurrentHash":"jk2gH/emOEAS4L6rD1umcwgjih/PatbazARXH8lKzWo=","Path":["+r0JSJPWaRVOEmkrllfsBocDealHPBrm+6jtg7Mwkyk=","0ac2gDZrYRNX71XSM0dz2SpS5DMcvDbCaUkbUiGxUlw=","cyLIAbjOGKv/iUKQmibxdhw75+V+3AA5Grx13v/Jjcc="],"Indexes":[1,1,1]}'

curl -v -X POST "http://localhost:8081/VerifyPath" -H 'content-type: application/json' -d '{"VerifyInputInfo":{"CertDetail":{"CertNo": "TEST20220413_5", "ID": "M123456(0)", "Name":"CoronaVac" , "Brand": "SinoVac" ,   "NumOfDose":"2","Time":"2022-04-03T15:50", "Issuer":"vac_center1" } , "PersonInfoHash":"047d1be86e0ed70abfd55540a1d6fe09c91c74eb99f82d4f77926b721a87b0fe","Key":"413hfu234"}, "VerifyPath":{"GlobalID":"c5fe5eb8-c4aa-49b6-9422-2785dbd0e71a","CurrentHash":"","Path":["MT7xnxMjJWim0HU6QWjqN1jK4flOePPDLFtUkbQBRpE=","b5Ws2KFsegb0JIj292cLBZTc7BPkNBpfVGN8xtZ/dYs=","LtVsvHo6YIcXjMvbPc2meOnKegTNo4n2o8S6KWue6JU="],"Indexes":[0,1,0]}}'

