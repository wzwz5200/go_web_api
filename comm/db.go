package comm

import (
	"fmt"
	"hellow/model"
	"log"
	"net/url"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var DB *gorm.DB
var PgDB *gorm.DB


func InitDB() *gorm.DB {

	host := "192.168.10.240"
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

func GetPG() *gorm.DB {
	dsn := "host=127.0.0.1 port=5432 user=postgres dbname=golang password=123456 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	db.AutoMigrate(&model.User{})

	PgDB = db
	return db
}

func GetPGDB() *gorm.DB {
	return PgDB
}