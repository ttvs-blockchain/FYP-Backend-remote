

for i in {0..100}
do 
    time curl -X POST http://localhost:8080/CreateAsset   -H 'content-type: application/json' -d '{"CertDetail":{"CertID": "TEST20220206_'$i'", "PersonSysID": "M123456(0)", "Name":"CoronaVac" , "Brand": "SinoVac" ,   "NumOfDose":"2","Time":"2022-04-03T15:50", "Issuer":"vac_center1" } , "PersonInfoHash":"047d1be86e0ed70abfd55540a1d6fe09c91c74eb99f82d4f77926b721a87b0fe","Key":"413hfu234"}' 

done 


# curl -v -X GET "http://localhost:8080/GetAllAssets"

time curl -v -X POST http://localhost:8080/Upload


 time curl -X POST http://localhost:8080/CreateAsset   -H 'content-type: application/json' -d '{"CertDetail":{"CertID": "TEST20220206_fafdfd", "PersonSysID": "M123456(0)", "Name":"CoronaVac" , "Brand": "SinoVac" ,   "NumOfDose":"2","Time":"2022-04-03T15:50", "Issuer":"vac_center1" } , "PersonInfoHash":"047d1be86e0ed70abfd55540a1d6fe09c91c74eb99f82d4f77926b721a87b0fe","Key":"413hfu234"}' 