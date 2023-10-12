package main

import (
    "github.com/gin-gonic/gin"
    "TaskManagementApp/routes"
)

func main(){

	r:=gin.Default()

	routes.TaskRouters(r)

	r.Run(":8080")
}