package migration

import (
	"Gotest/models"

	"gorm.io/gorm"
)

func MigrateFiles(db *gorm.DB) error {
	return db.AutoMigrate(&models.File{})
}
