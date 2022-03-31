package models

import (
	"database/sql"
	"fmt"
	"log"
)

var MyDatabase *sql.DB

func InitDB() (err error) {
	// DSN:Data Source Name

	MyDatabase, err = sql.Open("mysql", DSN)
	if err != nil {
		return err
	}
	err = MyDatabase.Ping()
	if err != nil {
		return err
	}
	return nil
}

func InsertCert(asset Asset, info LocalChainInfo) (err error) {
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
		info.LocalChainID,
		info.LocalChainTxHash,
		info.LocalChainBlockNum,
		info.LocalChainTimeStamp)
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

func InsertGlobalHash(info GlocalChainInfo) (err error) {
	sqlStr := INSERT_GLOBAL_HASH_SQL
	ret, err := MyDatabase.Exec(sqlStr,
		info.CertIDList,
		info.GlobalChainTxHash,
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

func UpdateRow(info LocalChainInfo, id string) (err error) {
	sqlStr := UPDATE_SQL
	ret, err := MyDatabase.Exec(sqlStr, info.LocalChainID, info.LocalChainTxHash, info.LocalChainBlockNum, info.LocalChainTimeStamp, id)
	if err != nil {
		return err
	}
	fmt.Printf("insert success, ", ret)
	return nil
}

func ReadRowForMKTree() ([]string, error) {
	sqlStr := READ_ROW_FOR_MKTREE_SQL
	rows, err := MyDatabase.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	fmt.Printf("query success\n")

	var ids []string
	for rows.Next() {
		var id string

		err := rows.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id)
		ids = append(ids, id)
	}
	return ids, err
}
