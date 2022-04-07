package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"gopkg.in/ini.v1"
)

var MyDatabase *sql.DB

func InitDB() (err error) {
	// DSN:Data Source Name
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	dsn := cfg.Section("server").Key("dsn").String()
	MyDatabase, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = MyDatabase.Ping()
	if err != nil {
		return err
	}
	return nil
}

func InsertLocalDBCert(asset Asset, personInfoHash, keyHash string) (err error) {
	sqlStr := INSERT_CERT_SQL
	ret, err := MyDatabase.Exec(sqlStr,
		asset.CertID,
		asset.PersonSysID,
		asset.Name,
		asset.Brand,
		asset.NumOfDose,
		asset.Time,
		asset.Issuer,
		asset.Remark,
		personInfoHash,
		keyHash,
	)
	if err != nil {
		return err
	}
	theID, err := ret.LastInsertId()
	if err != nil {
		return err
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
	return nil
}

func InsertMultipleToGlobalHashDB(infos []GlobalChainInfo) (err error) {
	sqlStr := INSERT_GLOBAL_HASH_SQL
	vals := []interface{}{}
	for _, info := range infos {
		sqlStr += "(?, ?, ?, ?),"
		vals = append(vals, info.CertIDList, info.MerkleTreeRoot, info.GlobalChainBlockNum, info.GlobalChainTimeStamp)
	}
	//trim the last ,
	sqlStr = sqlStr[0 : len(sqlStr)-1]
	stmt, _ := MyDatabase.Prepare(sqlStr)
	res, _ := stmt.Exec(vals...)
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Printf("insert success, there are %v rows affected\n", rowsAffected)

	return nil
}

func UpdateLocalCertDB(info LocalChainInfo, id string) (err error) {

	path, err := json.Marshal(info.MerkleTreePathDetail.Path)
	if err != nil {
		return err
	}
	indexes, err := json.Marshal(info.MerkleTreePathDetail.Indexes)
	if err != nil {
		return err
	}

	sqlStr := UPDATE_SQL
	ret, err := MyDatabase.Exec(sqlStr,
		info.LocalChainID,
		string(path),
		string(indexes),
		info.MerkleTreePathDetail.GlobalRootID,
		info.LocalChainBlockNum,
		info.LocalChainTimeStamp,
		id)
	if err != nil {
		return err
	}
	fmt.Printf("insert success, %d\n", ret)
	return nil
}

// func UpdateMultipleLocalCertDB(infos []LocalChainInfo, dailyRecord []InputInfo) (err error) {

// 	sqlStr := UPDATE_MULTIPLE_SQL_START

// 	vals := []interface{}{}
// 	for i, info := range infos {
// 		sqlStr += "(?, ?, ?, ?, ?, ?, ?),"
// 		vals = append(vals,
// 			info.LocalChainID,
// 			info.MerkleTreePathDetail.Path,
// 			info.MerkleTreePathDetail.Indexes,
// 			info.MerkleTreePathDetail.GlobalRootID,
// 			info.LocalChainBlockNum,
// 			info.LocalChainTimeStamp,
// 			dailyRecord[i].CertDetail.CertID)
// 	}
// 	//trim the last ,
// 	sqlStr = sqlStr[0 : len(sqlStr)-1]

// 	sqlStr += UPDATE_MULTIPLE_SQL_END
// 	log.Println("-->DEBUG SQL STR " + sqlStr)

// 	log.Println(vals)

// 	stmt, _ := MyDatabase.Prepare(sqlStr)
// 	res, _ := stmt.Exec(vals...)

// 	if err != nil {
// 		return err
// 	}
// 	fmt.Printf("insert success, %v\n", res)

// 	return nil

// }

func ReadRowForMKTree() ([][]string, error) {
	sqlStr := READ_ROW_FOR_MKTREE_SQL
	rows, err := MyDatabase.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	fmt.Printf("query success\n")

	var listOfRow [][]string
	for rows.Next() {
		var id, personInfoHash, keyHash string
		err := rows.Scan(&id, &personInfoHash, &keyHash)
		if err != nil {
			log.Fatal(err)
		}
		listOfRow = append(listOfRow, []string{id, personInfoHash, keyHash})
	}
	return listOfRow, err
}

func ReadPath(certID string) (*Asset, string, string, error) {
	sqlStr := READ_PATH_SQL

	var personSysID, name, brand, numOfDose, issueTime, issuer, remark, personInfoHash, merkleTreePath string

	err := MyDatabase.QueryRow(sqlStr, certID).Scan(&personSysID, &name, &brand, &numOfDose, &issueTime,
		&issuer, &remark, &personInfoHash, &merkleTreePath)
	if err != nil {
		fmt.Printf("Error,  %s,\n", err.Error())

		return nil, "", "", err
	}

	fmt.Printf("debug 999, personSysID is %s,\n", personSysID)

	asset := Asset{
		certID, personSysID, name, brand, numOfDose, issueTime, issuer, remark}
	fmt.Printf("debug 0000, asset is %s, personInfoHash is %s, path is %s\n", asset, personInfoHash, merkleTreePath)

	return &asset, personInfoHash, merkleTreePath, nil
}
