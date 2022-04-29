/*
Copyright 2020 IBM All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package main

import (
	"asset-transfer-basic/controllers"
	_ "asset-transfer-basic/controllers"
	"asset-transfer-basic/middlewares"
	"asset-transfer-basic/models"

	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func main() {

	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    	if err != nil {
   	     log.Fatal(err)
    	}

    	log.SetOutput(file)
	log.Println("============ application-golang starts ============")

	err = models.InitDB() //
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	log.Printf("============ database connected  ============")

	log.Println("============ application-golang ends ============")

	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())
	//r.Use(middlewares.TlsHandler())

	r.GET("/GetAllAssets", controllers.GetAllAssets)

	r.POST("/VerifyPath", controllers.VerifyPath)

//	r.Run(":8081")
	r.RunTLS(":8081", "./vaxpass.ttommy.tech_apache/vaxpass.ttommy.tech.crt", "./vaxpass.ttommy.tech_apache/vaxpass.ttommy.tech.key") 

}
