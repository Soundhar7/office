package main

import (
	"myproject/database"
	"myproject/routes"

	"github.com/gin-gonic/gin"
	//"github.com/go-delve/delve/service"
)

func main(){

	r:=gin.Default()

	db,err:=database.InitDB()

	if err!=nil{

		panic(err)
	}

	service :=routes.DBCustomerservice{Db :db}
	routes.SetupRoutes(r,service)
	r.Run(":8080")

}