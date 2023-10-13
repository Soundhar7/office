package database

import (
	"fmt"
	"myproject/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {



	dsn:="user=postgres password=Soundhar@10 dbname=db sslmode=disable"
  var err error
	db, err = gorm.Open(postgres.Open(dsn),&gorm.Config{})

	if err != nil {

		fmt.Println(err)
		return db, err
	}

	db.AutoMigrate(&models.Customer{}, &models.Orders{})
	return db,nil


}
