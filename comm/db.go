package comm

import (
	"fmt"
	"hellow/model"
	"net/url"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var DB *gorm.DB


func InitDB() *gorm.DB {

	host := "127.0.0.1"
	port := "3306"
	database := "golang"
	username := "golang"
	password := "golang"

	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&loc=%s&parseTime=true", username, password, host, port, database, url.QueryEscape("Asia/Shanghai"))
	db, err := gorm.Open(mysql.Open(args))

	if err != nil {

		panic("err!")
	}
	db.AutoMigrate(&model.User{})

	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}