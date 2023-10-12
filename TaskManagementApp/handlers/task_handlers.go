package handlers

import (
	"TaskManagementApp/Config"
	"TaskManagementApp/models"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	//"github.com/go-playground/validator/v10/translations/id"
	"github.com/jinzhu/gorm"
)

//Create Task

func CreateTask(c *gin.Context){
	db,err:=config.InitDB()

	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"Error":"Not connected"})

	}
	defer db.Close()
	var task models.Task

	if err:=c.ShouldBindJSON(&task);err !=nil{
		c.JSON(http.StatusBadRequest,gin.H{"Error":"Invalid input "})
		return
	}
	if err:=db.Create(&task).Error;err !=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"Error":"Not created"})
		return
	}
	c.JSON(http.StatusCreated,task)
}


//Get all Tasks

func GetTask(c *gin.Context){
	db,err:=config.InitDB()

	if err != nil{

		c.JSON(http.StatusNotFound,gin.H{"Error":"Not connected"})
	}
	var task []models.Task

	db.Find(&task)
	c.JSON(http.StatusOK,task)
}


//Delete Tasks

func Deletetask(c * gin.Context){
	db,err:=config.InitDB()

	if err!= nil{
		c.JSON(http.StatusBadRequest,gin.H{"Error":"Not connected"})
		return
	}
	defer db.Close()
	idstr:=c.Param("id")

	id,err:=strconv.Atoi(idstr)

	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"Error":"Invalied id"})
		return
	}
	if err:=db.Where("id=?",id).Delete(models.Task{}).Error;err!=nil{

		if gorm.IsRecordNotFoundError(err){

			c.JSON(http.StatusNotFound,gin.H{"Error":"Not Found"})
			return

		}
	}
	fmt.Println("Message deleted")
	c.JSON(http.StatusOK,gin.H{"Message":"Deleted"})
}


// Update Task

func UpdateTask(c *gin.Context){
	db,err:=config.InitDB()

	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"Erorr":"Not connected"})
		return
	}
	defer db.Close()

	idstr:=c.Param("id")

	id,err:=strconv.Atoi(idstr)

	if err != nil{

		c.JSON(http.StatusBadRequest,gin.H{"Error":"Invalid id"})
	}
	var updatedtask models.Task

	if err:=c.ShouldBindJSON(&updatedtask);err!=nil{

		c.JSON(http.StatusInternalServerError,gin.H{"message":"Not updated"})
		return
	}
	var exsitingtask models.Task

	if err:=db.First(&exsitingtask,id).Error;err!=nil{
		if gorm.IsRecordNotFoundError(err){
			c.JSON(http.StatusNotFound,gin.H{"Error":"Not found"})
			return
		}
	}

	exsitingtask.Title=updatedtask.Title
	exsitingtask.Description =updatedtask.Description
    
	exsitingtask.AssignedTo = updatedtask.AssignedTo

	exsitingtask.Status = updatedtask.Status

	if err:=db.Save(&exsitingtask).Error;err !=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"Error":"Can't find matching Employee"})
		return
	}
	c.JSON(http.StatusOK,exsitingtask)

}


//Search By any Tasks

func SearchByQuery(c *gin.Context) {
    db, err := config.InitDB()
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"Error": "Not Connected"})
        return
    }
	defer db.Close()

    var newtasks []models.Task 
    searchQuery := c.DefaultQuery("q", "")

   
    if _, err := strconv.Atoi(searchQuery); err == nil {
        db.Where("id = ?", searchQuery).Find(&newtasks)
    } else {
        db.Where("LOWER(title) LIKE ? OR LOWER(description) LIKE ? OR LOWER(status) LIKE ?","%"+strings.ToLower(searchQuery)+"%","%"+strings.ToLower(searchQuery)+"%","%"+strings.ToLower(searchQuery)+"%",).Find(&newtasks)
    }

    if len(searchQuery) == 0 {
        c.JSON(http.StatusNotFound, gin.H{"Error": "No matching tasks found"})
        return
    }

    c.JSON(http.StatusOK, newtasks)
}
