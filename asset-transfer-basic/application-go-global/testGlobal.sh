

time curl -v -X POST "https://125.59.138.87:8081/VerifyPath" -k -H 'content-type: application/json' -d '{
    "VerifyInputInfo": {
        "CertDetail": {
            "CertID": "TEST20220206_10",
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
        "GlobalRootID": "66048fd0-5bf7-4e04-ab37-5fa7eb607e63",
        "Path": 
        [
            "zjLNBrmixWR6BeJGMdeO8vMXZhaXOQgfhrAOAILV5sA=",
            "abURhtxXORhrYLLk78zIPBS0IhxR2YKxP9sEvjVuVi8=",
            "ji/uMusHTRqmwFV2Z1cjPH6ZdJKhjt/S4euiUt8V0CY=",
            "C7UN8l0e83IZ+7bZxEFeOjLQjbhkfGUPJW0OfNfqpMY="
        ],
        "Indexes": [
            1,
            0,
            1,
            0
        ]
    }
}'

