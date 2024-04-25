package models

import (
	"time"
	"gorm.io/gorm"
)

type File struct {
	ID uint `gorm:";primaryKey"`
	Post_ID uint `gorm:""`
	File_Name string `gorm:"not null"`
	File_Path string `gorm:"not null"`
	CreatedAt time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:datetime;default:null"`
}

func (file *File) TableName() string {
	return "files"
}

func (file *File) BeforeCreate(tx *gorm.DB) (err error) {
	tx.Statement.SetColumn("CreatedAt", time.Now())
	return
}

func (file *File) BeforeUpdate(tx *gorm.DB) (err error) {
	tx.Statement.SetColumn("UpdatedAt", time.Now())
	return
}

func MigrateFiles(db *gorm.DB) error {
	err := db.AutoMigrate(&File{})
	if err != nil {
		return err
	}
	return nil
}
