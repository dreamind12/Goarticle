package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

)

var DB *gorm.DB

func InitDB() {
    dsn := "root:@tcp(localhost:3306)/article?charset=utf8mb4&parseTime=True&loc=Local"
    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Connected to the database using GORM")
}
