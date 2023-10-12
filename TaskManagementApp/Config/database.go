package config

import(

	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "TaskManagementApp/models"
)

var db  *gorm.DB

var err error

func InitDB() (*gorm.DB, error) {

	db,err=gorm.Open("postgres","user=postgres password=Soundhar@10 dbname=task sslmode=disable")
	
if err !=nil{

	return nil,err


}


db.AutoMigrate(&models.Task{})

return db,nil
}