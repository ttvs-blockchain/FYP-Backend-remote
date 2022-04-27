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
	"asset-transfer-basic/middlewares"
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
	r.Use(middlewares.CORSMiddleware())

	r.GET("/ReadAsset", controllers.ReadAsset)

	r.GET("/GetAllAssets", controllers.GetAllAssets)

	r.POST("/CreateAsset", controllers.CreateAsset)

	r.POST("/Upload", controllers.Upload)
	r.RunTLS(":8080", "./tlsCert/cert.pem", "./tlsCert/key.pem") //
	r.Use(middlewares.TlsHandler())

	log.Printf("============ application-golang ends ============")

}
