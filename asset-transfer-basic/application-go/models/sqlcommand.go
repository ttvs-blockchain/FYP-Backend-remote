package models

const (
	DSN string = "tommy:mysql123456@tcp(167.179.77.244:3306)/certificate?charset=utf8mb4&parseTime=True"

	INSERT_CERT_SQL string = "insert into localCertificate(certID, personID, name, brand, numOfDose, issueTime, issuer, remark, personHash) values (?,?,?,?,?,?,?,?,?)"

	INSERT_GLOBAL_HASH_SQL string = "insert into globalChainInfo(certIDList, merkelTreeRoot,  globalChainBlockNum, globalChainTimeStamp) values (?,?,?,?)"

	UPDATE_SQL string = "update  localCertificate set localChainID = ?, merkelTreePath = ?, localChainBlockNum = ?, localChainTimeStamp =?  where certID = ?"

	READ_ROW_FOR_MKTREE_SQL string = "select certID, personHash from localCertificate where localChainTimeStamp = 0 "

	READ_PATH_SQL string = "select personID, name, brand, numOfDose, issueTime, issuer, remark, personHash, merkelTreePath from localCertificate where certID = ? "
)
