package database

import (
	"log"

	"Gotest/migrationseeder/migration"

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

	if err := migration.MigrateUsers(DB); err != nil {
		log.Fatal("Error migrating users table:", err)
	}

	if err := migration.MigratePosts(DB); err != nil {
		log.Fatal("Error migrating posts table:", err)
	}

	if err := migration.MigrateFiles(DB); err != nil {
		log.Fatal("Error migrating file table:", err)
	}

	log.Println("Database migration completed successfully")
}
