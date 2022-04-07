package models

const (
	INSERT_CERT_SQL string = "insert into localCertificate(certID, personSysID, name, brand, numOfDose, issueTime, issuer, remark, personInfoHash, keyHash) values (?,?,?,?,?,?,?,?,?,?)"

	INSERT_GLOBAL_HASH_SQL string = "insert into globalChainInfo(certIDList, merkleTreeRoot,  globalChainBlockNum, globalChainTimeStamp) values "

	UPDATE_SQL string = "update  localCertificate set localChainID = ?, merkleTreePath = ?, merkleTreeIndexes = ?, globalRootID = ?, localChainBlockNum = ?, localChainTimeStamp = ?  where certID = ?"

	READ_ROW_FOR_MKTREE_SQL string = "select certID, personInfoHash, keyHash from localCertificate where localChainTimeStamp = 0 "

	READ_PATH_SQL string = "select personSysID, name, brand, numOfDose, issueTime, issuer, remark, personInfoHash, merkleTreePath from localCertificate where certID = ? "

	UPDATE_MULTIPLE_SQL_START string = "insert into localCertificate(localChainID, merkleTreePath,  merkleTreeIndexes, globalRootID, localChainBlockNum, localChainTimeStamp, certID) values"

	UPDATE_MULTIPLE_SQL_END string = "ON DUPLICATE KEY UPDATE localChainID=VALUES(localChainID), merkleTreePath=VALUES(merkleTreePath), merkleTreeIndexes=VALUES(merkleTreeIndexes), globalRootID=VALUES(globalRootID), localChainBlockNum=VALUES(localChainBlockNum), localChainTimeStamp=VALUES(localChainTimeStamp)"
)
