package database

import (
	"gorm.io/driver/postgres"
    "gorm.io/gorm"
)


type User struct{

	ID  uint `gorm:"primarykey"`

	Name string `gorm:"not null"`
} 

var db *gorm.DB 

func InitDB() (*gorm.DB, error) {
	dsn := "user=postgres password=Soundhar@10 dbname=mydb sslmode=disable"

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	
	 if err := db.AutoMigrate(&User{}); err != nil {
	 	return nil, err
	 }

	return db, nil
}

     
