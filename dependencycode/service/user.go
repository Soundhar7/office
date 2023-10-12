package service

import (
	"fmt"
	"myproject/database"
	"net/http"
	"strconv"

	//"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EmployeeService interface {
	GetAlluser() ([]database.User, error)
	Createuser(user database.User) error
	Updateuser(user database.User) error
	DeleteUser(id uint) error

	GetUserbyid(id uint) (database.User, error)
}

type DBUserService struct {
	Db *gorm.DB
}

func (d DBUserService) GetAlluser() ([]database.User, error) {

	var alluser []database.User

	// if err := d.Db.Find(&alluser).Error; err != nil {

	// 	return nil, err
	// }
	// return alluser, nil

	db:=d.Db.Begin()

	if err:=db.Find(&alluser).Error;err!=nil{

		db.Rollback()
		return nil,err
	}
	if err:=db.Commit().Error;err !=nil{

		db.Rollback()
		return nil,err
	}
	return alluser,nil
}

func (d DBUserService) Createuser(user database.User) error {

	// if err := d.Db.Create(&user).Error; err != nil {

	// 	return err
	// }
	// return nil
	db:=d.Db.Begin()

	if err:=db.Create(&user).Error;err !=nil{

		db.Rollback()
		return err

	}
	if err:=db.Commit().Error;err!=nil{

		db.Rollback()
		return err
	}
	return nil
}

func (d DBUserService) Updateuser(user database.User) error {

	// if err := d.Db.Save(&user).Error; err != nil {

	// 	return err
	// }
	// return nil

	db:=d.Db.Begin()

	if err:=db.Save(&user).Error;err!=nil{
		db.Rollback()
		return err
	}
	if err:=db.Commit().Error;err!=nil{
		db.Rollback()
		return err
	}
	return nil
}

func (d DBUserService) GetUserbyid(id uint) (database.User, error) {
	var user database.User

	// if err := d.Db.First(&user, id).Error; err != nil {
	// 	return database.User{}, err
	// }
	// return user, nil

	db:=d.Db.Begin()

	if err:=db.First(&user,id).Error;err!=nil{

		db.Rollback()
		return database.User{},err
	}
	if err:=db.Commit().Error;err!=nil{

		db.Rollback()
		return database.User{},nil
	}
	return user,nil
}

func (d DBUserService) DeleteUser(id uint) error {

	// if err := d.Db.Delete(&database.User{}, id).Error; err != nil {

	// 	return err

	// }
	// return nil

	db:=d.Db.Begin()

	if err:=db.Delete(&database.User{},id).Error;err!=nil{

		db.Rollback()
		return err
	}
	if err:=db.Commit().Error;err!=nil{

		db.Commit()
		return err
	}
	return nil
}

func InitializeRoutes(r *gin.Engine, empService EmployeeService) {

	r.GET("/alluser", func(c *gin.Context) {

		allusers, err := empService.GetAlluser()

		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"Error ": "Internal Server Error"})
			return
		}
		c.JSON(http.StatusOK, allusers)

	})

	r.POST("/create", func(c *gin.Context) {

		var newuser database.User

		if err := c.ShouldBindJSON(&newuser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request"})

			return                               
		}

		if err := empService.Createuser(newuser); err != nil {

			c.JSON(http.StatusNotFound, gin.H{"Error": "Not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"Message": "Created"})
	})

	r.PUT("/updateuser/:id", func(c *gin.Context) {

		idstr := c.Param("id")

		id, err := strconv.ParseUint(idstr, 10, 0)

		if err != nil {

			c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid id"})
			return
		}

		var updateuser database.User

		if err := c.ShouldBindJSON(&updateuser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Request"})
			return
		}

		updateuser.ID = uint(id)

		if err := empService.Updateuser(updateuser); err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Internal server error"})
			return
		}
		c.JSON(http.StatusOK, updateuser)

	})

	r.DELETE("/delete/:id", func(c *gin.Context) {

		idstr := c.Param("id")

		id, err := strconv.ParseUint(idstr, 10, 0)

		if err != nil {

			c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Id"})
			return
		}
		err = empService.DeleteUser(uint(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Internal Server Error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"Message": "User deleted"})

	})

	r.GET("/getuesrbyid/:id", func(c *gin.Context) {

		idstr := c.Param("id")

		userid, err := strconv.ParseUint(idstr, 10, 0)
		if err != nil {

			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid id"})
			return

		}
		user, err := empService.GetUserbyid(uint(userid))

		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"Error": "User Not found"})
			return
		}
		c.JSON(http.StatusOK, user)

	})
}
