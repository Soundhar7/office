package routes
import (
    "github.com/gin-gonic/gin"
     "TaskManagementApp/handlers"
)

func TaskRouters(r * gin.Engine){

	taskrouter:=r.Group("/task")
	{
		taskrouter.POST("/create",handlers.CreateTask)

		taskrouter.GET("/getall",handlers.GetTask)

		taskrouter.DELETE("/deletetask/:id",handlers.Deletetask)

		taskrouter.PUT("/updatetask/:id",handlers.UpdateTask)
		taskrouter.GET("/search",handlers.SearchByQuery)

		
	}
}