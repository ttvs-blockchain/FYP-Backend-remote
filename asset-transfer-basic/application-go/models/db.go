package models

import (
	"database/sql"
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

func InsertLocalDBCert(asset Asset, personHash string) (err error) {
	sqlStr := INSERT_CERT_SQL
	ret, err := MyDatabase.Exec(sqlStr,
		asset.CertNo,
		asset.ID,
		asset.Name,
		asset.Brand,
		asset.NumOfDose,
		asset.Time,
		asset.Issuer,
		asset.Remark,
		personHash,
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

func InsertGlobalHashDB(info GlobalChainInfo) (err error) {
	sqlStr := INSERT_GLOBAL_HASH_SQL
	ret, err := MyDatabase.Exec(sqlStr,
		info.CertIDList,
		info.MerkleTreeRoot,
		info.GlobalChainBlockNum,
		info.GlobalChainTimeStamp)
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

func UpdateLocalCertDB(info LocalChainInfo, id string) (err error) {
	sqlStr := UPDATE_SQL
	ret, err := MyDatabase.Exec(sqlStr,
		info.LocalChainID,
		info.MerkleTreePathDetail,
		info.LocalChainBlockNum,
		info.LocalChainTimeStamp,
		id)
	if err != nil {
		return err
	}
	fmt.Printf("insert success, %d\n", ret)
	return nil
}

func ReadRowForMKTree() ([][]string, error) {
	sqlStr := READ_ROW_FOR_MKTREE_SQL
	rows, err := MyDatabase.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	fmt.Printf("query success\n")

	var listOfRow [][]string
	for rows.Next() {
		var id string
		var personHash string
		err := rows.Scan(&id, &personHash)
		if err != nil {
			log.Fatal(err)
		}
		idAndPersonHash := []string{id, personHash}
		listOfRow = append(listOfRow, idAndPersonHash)
	}
	return listOfRow, err
}

func ReadPath(certID string) (*Asset, string, string, error) {
	sqlStr := READ_PATH_SQL

	var personID, name, brand, numOfDose, issueTime, issuer, remark, personHash, merkleTreePath string

	err := MyDatabase.QueryRow(sqlStr, certID).Scan(&personID, &name, &brand, &numOfDose, &issueTime,
		&issuer, &remark, &personHash, &merkleTreePath)
	if err != nil {
		fmt.Printf("Error,  %s,\n", err.Error())

		return nil, "", "", err
	}

	fmt.Printf("debug 999, personID is %s,\n", personID)

	asset := Asset{
		certID, personID, name, brand, numOfDose, issueTime, issuer, remark}
	fmt.Printf("debug 0000, asset is %s, personHash is %s, path is %s\n", asset, personHash, merkleTreePath)

	return &asset, personHash, merkleTreePath, nil
}
