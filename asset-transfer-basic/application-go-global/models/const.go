package models

const (
	DSN string = "tommy:mysql123456@tcp(167.179.77.244:3306)/certificate?charset=utf8mb4&parseTime=True"

	INSERT_GLOBAL_HASH_SQL string = "insert into globalChainInfo(certIDList, globalChainTxHash, globalChainBlockNum, globalChainTimeStamp) values (?,?,?,?)"
)
