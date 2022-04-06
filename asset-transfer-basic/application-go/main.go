/*
Copyright 2020 IBM All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package main

import (
	"fmt"
	"log"

	"asset-transfer-basic/controllers"
	_ "asset-transfer-basic/controllers"
	"asset-transfer-basic/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log.Printf("============ application-golang starts ============")

	err := models.InitDB() //
	if err != nil {
		fmt.Printf("init db failed, err:%v\n", err)
		return
	}
	log.Printf("============ database connected  ============")

	r := gin.Default()

	r.GET("/ReadAsset", controllers.ReadAsset)

	r.GET("/GetAllAssets", controllers.GetAllAssets)

	r.POST("/CreateAsset", controllers.CreateAsset)

	r.POST("/Upload", controllers.Upload)

	r.GET("/GetPath", controllers.GetPath)

	r.Run(":8080") //

	log.Printf("============ application-golang ends ============")

}
