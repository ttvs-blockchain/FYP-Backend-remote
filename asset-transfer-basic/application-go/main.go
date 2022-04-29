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
//	"asset-transfer-basic/middlewares"
	"asset-transfer-basic/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	cors "github.com/rs/cors/wrapper/gin"
	"os"
)

func main() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	 if err != nil {
       	 log.Fatal(err)
    	}
    	log.SetOutput(file)
	log.Printf("============ application-golang starts ============")

	err = models.InitDB() //
	if err != nil {
		fmt.Printf("init db failed, err:%v\n", err)
		return
	}
	log.Printf("============ database connected  ============")

	r := gin.Default()
	r.Use(cors.Default())
//	r.Use(middlewares.CORSMiddleware())

	r.GET("/ReadAsset", controllers.ReadAsset)

	r.GET("/GetAllAssets", controllers.GetAllAssets)

	r.POST("/CreateAsset", controllers.CreateAsset)

	r.POST("/Upload", controllers.Upload)
//	r.Run(":8080")
	r.RunTLS(":8080", "./vaxpass.ttommy.tech_apache/vaxpass.ttommy.tech.crt", "./vaxpass.ttommy.tech_apache/vaxpass.ttommy.tech.key") //
//	r.Use(middlewares.TlsHandler())
//	r.Use(middlewares.CORSMiddleware())

	log.Printf("============ application-golang ends ============")

}
