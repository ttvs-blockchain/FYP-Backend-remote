# curl -v -X GET "http://localhost:8081/GetAllAssets"

# curl -v -X POST http://localhost:8081/Verify -H 'content-type: application/json' -d ' {"GlobalID":"0db1a2ef-7dc3-426d-950a-c710ee603f31","CurrentHash":"jk2gH/emOEAS4L6rD1umcwgjih/PatbazARXH8lKzWo=","Path":["+r0JSJPWaRVOEmkrllfsBocDealHPBrm+6jtg7Mwkyk=","0ac2gDZrYRNX71XSM0dz2SpS5DMcvDbCaUkbUiGxUlw=","cyLIAbjOGKv/iUKQmibxdhw75+V+3AA5Grx13v/Jjcc="],"Indexes":[1,1,1]}'

curl -v -X POST "http://localhost:8081/VerifyCert" -H 'content-type: application/json' -d '{"CertNo":"TEST20220413_06","LocalChainID":"1", "PersonHash":"047d1be86e0ed70abfd55540a1d6fe09c91c74eb99f82d4f77926b721a87b0fe"}'

