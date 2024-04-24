package migration

import (
	"Gotest/models"
	"gorm.io/gorm"
)

func MigrateUsers(db *gorm.DB) error {
	return db.AutoMigrate(&models.User{})
}
