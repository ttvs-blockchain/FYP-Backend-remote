package models

const (
	DSN string = "tommy:mysql123456@tcp(167.179.77.244:3306)/certificate?charset=utf8mb4&parseTime=True"

	INSERT_SQL string = "insert into localCertificate(id, certID, personID, name, brand, numOfDose, issueTime, issuer, remark, localChainID, localChainTxHash, localChainBlockNum, localChainTimeStamp) values (?,?,?,?,?,?,?,?,?,?,?,?,?)"
)
