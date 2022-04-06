package models

const (
	INSERT_CERT_SQL string = "insert into localCertificate(certID, personID, name, brand, numOfDose, issueTime, issuer, remark, personInfoHash) values (?,?,?,?,?,?,?,?,?)"

	INSERT_GLOBAL_HASH_SQL string = "insert into globalChainInfo(certIDList, merkleTreeRoot,  globalChainBlockNum, globalChainTimeStamp) values (?,?,?,?)"

	UPDATE_SQL string = "update  localCertificate set localChainID = ?, merkleTreePath = ?, localChainBlockNum = ?, localChainTimeStamp =?  where certID = ?"

	READ_ROW_FOR_MKTREE_SQL string = "select certID, personInfoHash from localCertificate where localChainTimeStamp = 0 "

	READ_PATH_SQL string = "select personID, name, brand, numOfDose, issueTime, issuer, remark, personInfoHash, merkleTreePath from localCertificate where certID = ? "
)
