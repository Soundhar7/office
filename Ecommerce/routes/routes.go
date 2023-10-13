package routes

import (
	"myproject/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	//"github.com/pelletier/go-toml/query"
	"gorm.io/gorm"
)

type Repository interface {
	Createnewcustomer(customer models.Customer) error
	GetAllCustomers() ([]models.Customer, error)
	Updatecustomer(customer models.Customer) error
	Deletecustomer(id int)error
	GetById(id int)(models.Customer,error)
	SearchByAny(query string)([]models.Customer,error)
}

type DBCustomerservice struct {
	Db *gorm.DB
}

func (d DBCustomerservice) Createnewcustomer(customer models.Customer) error {

	db := d.Db.Begin()

	if err := db.Create(&customer).Error; err != nil {

		db.Rollback()
		return err
	}
	if err := db.Commit().Error; err != nil {

		db.Rollback()
		return err
	}
	return nil
}

func (d DBCustomerservice) GetAllCustomers() ([]models.Customer, error) {

	var allcustomer []models.Customer

	db := d.Db.Begin()

	if err := db.Find(&allcustomer).Error; err != nil {

		db.Rollback()
		return nil, err
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return nil, err
	}
	return allcustomer, nil
}

func (d DBCustomerservice) Updatecustomer(customer models.Customer) error {

	db := d.Db.Begin()

	if err := db.Save(&customer).Error; err != nil {
		db.Rollback()
		return err
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}
	return nil
}

func(d DBCustomerservice)Deletecustomer(id int)error{

	db:=d.Db.Begin()

	if err:=db.Delete(&models.Customer{},id).Error;err!=nil{

		db.Rollback()
		return err
	}
	if err:=db.Commit().Error;err !=nil{

		db.Rollback()
		return err
	}
	return nil
}

func(d DBCustomerservice)GetById(id int)(models.Customer,error){

	var customer models.Customer

	db:=d.Db.Begin()

	if err:=db.First(&customer,id).Error;err!=nil{

		db.Rollback()

		return models.Customer{},err
	}
	if err:=db.Commit().Error;err != nil{

		db.Rollback()
		return models.Customer{},err
	}
	return customer,nil
}

func(d DBCustomerservice)SearchByAny(query string)([]models.Customer,error){
      
	

	db:=d.Db.Begin()
	var customer []models.Customer

	 id,err:=strconv.Atoi(query)
	if err==nil{

		db.Where("id=?",id).Find(&customer)

	}else{
		db.Where("name LIKE ?","%"+query+"%").Find(&customer)
	}
	if db.Error !=nil{
      db.Rollback()
		return nil,err

	}
	db.Commit()
	return customer,nil


}

func SetupRoutes(r *gin.Engine, service Repository) {

	r.POST("/create", func(c *gin.Context) {

		var newcustomer models.Customer

		if err := c.ShouldBindJSON(&newcustomer); err != nil {

			c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request"})
			return
		}

		if err := service.Createnewcustomer(newcustomer); err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Internal Server Error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"Message": "Created"})
	})

	r.GET("/getall", func(c *gin.Context) {

		allcustomer, err := service.GetAllCustomers()
		if err != nil {

			c.JSON(http.StatusBadRequest, gin.H{"Error": "Inavaild Request"})
			return
		}
		c.JSON(http.StatusOK, allcustomer)

	})

	r.PUT("/updatecus/:id", func(c *gin.Context) {

		idstr := c.Param("id")

		 id,err := strconv.ParseInt(idstr,10,0)
		 if err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{"Error":"Invalid Id"})
			return
		}

		var updatedcustomer models.Customer

		if err:=c.ShouldBindJSON(&updatedcustomer);err!=nil{

			c.JSON(http.StatusBadRequest,gin.H{"Error":"Invalid Request"})
			return
		}

		updatedcustomer.ID= int(id)

		if err:=service.Updatecustomer(updatedcustomer);err!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{"Error":"Internal Server Error"})
			return
		}
		c.JSON(http.StatusOK,updatedcustomer)

	})

	r.DELETE("/deletecustomer/:id",func(c *gin.Context){

		idstr:=c.Param("id")

		 id,err:=strconv.ParseInt(idstr,10,0);
		 if err !=nil{

			c.JSON(http.StatusBadRequest,gin.H{"Error":"Invalid Id"})
			return
		}
		if err:=service.Deletecustomer(int(id));err!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{"Error":"Internal Server Error"})
			return
		}
		c.JSON(http.StatusOK,gin.H{"Message":"deleted"})


	})

	r.GET("/getbyid/:id",func(c *gin.Context){

		idstr:=c.Param("id")

		id,err:=strconv.ParseInt(idstr,10,0)

		if err !=nil{

			c.JSON(http.StatusBadRequest,gin.H{"Error":"Invalid id"})
			return
		}

		 user,err:=service.GetById(int(id))
		 if err!=nil{

			c.JSON(http.StatusInternalServerError,gin.H{"Error":"Internal server Error"})
			return
		}
		c.JSON(http.StatusOK,user)
	})

	r.GET("/search", func(c *gin.Context) {
		query := c.DefaultQuery("q", "")
	
		if query == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
			return
		}
	
		result, err := service.SearchByAny(query)
	
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
	
		c.JSON(http.StatusOK, result)
	})
	
}
