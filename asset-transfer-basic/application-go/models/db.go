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

func InsertRow(asset Asset, info LocalChainInfo) {
	sqlStr := INSERT_SQL
	ret, err := MyDatabase.Exec(sqlStr, asset.CertNo, asset.ID, asset.Name, asset.Brand, asset.NumOfDose, asset.Time, asset.Issuer, asset.Remark, info.LocalChainID, info.LocalChainTxHash, info.LocalChainBlockNum, info.LocalChainTimeStamp)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}
