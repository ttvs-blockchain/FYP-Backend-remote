

curl -v -X POST "http://localhost:8081/VerifyPath" -H 'content-type: application/json' -d '{
    "VerifyInputInfo": {
        "CertDetail": {
            "CertID": "TEST20220300_5",
            "PersonSysID": "M123456(0)",
            "Name": "CoronaVac",
            "Brand": "SinoVac",
            "NumOfDose": "2",
            "Time": "2022-04-03T15:50",
            "Issuer": "vac_center1"
        },
        "PersonInfoHash": "047d1be86e0ed70abfd55540a1d6fe09c91c74eb99f82d4f77926b721a87b0fe",
        "Key": "413hfu234"
    },
    "VerifyPath": {
        "GlobalRootID": "955b658e-b3b1-4b7e-a9f4-e4b4059ed7a7",
        "Path": [
            "SUN83LhrpECne/8cb73N8MhbehBrsX7bTdaDdM748jM=",
            "KPjW2LycxjjpPuINv/jirDlLBuzE7tlmnFpTaCa/WtU=",
            "youFUQeKzNcHOgHXvTo52A1xknt1FdCPifWygccarAI="
        ],
        "Indexes": [
            0,
            1,
            0
        ]
    }
}'

