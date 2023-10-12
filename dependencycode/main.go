package main

import (
	"log"
	"myproject/database"
	"myproject/service"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	db,err:=database.InitDB()

	if err !=nil{
		log.Println("Not connected")
		return
	}

	//database.InitDB()
	//var db *gorm.DB

	empService:=service.DBUserService{Db: db}
	service.InitializeRoutes(r,empService)

	r.Run(":8080")
}
