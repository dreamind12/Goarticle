package migration

import (
	"Gotest/models"
	"gorm.io/gorm"
)

func MigratePosts(db *gorm.DB) error {
	return db.AutoMigrate(&models.Post{})
}