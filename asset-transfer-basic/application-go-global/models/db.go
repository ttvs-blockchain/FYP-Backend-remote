package models

import (
	"database/sql"
	"fmt"
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

func InsertGlobal(info GlobalChainInfo) error {
	sqlStr := INSERT_GLOBAL_HASH_SQL

	ret, err := MyDatabase.Exec(
		sqlStr,
		info.CertIDList,
		info.MerkleTreeRoot,
		info.GlobalChainBlockNum,
		info.GlobalChainTimeStamp)
	if err != nil {
		return err
	}
	theID, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return err

	}
	fmt.Printf("insert success, the id is %d.\n", theID)
	return nil
}
