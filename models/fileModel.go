package models

import (
	"time"
	"gorm.io/gorm"
)

type File struct {
	ID uint `gorm:"primaryKey"`
	Post_ID uint `gorm:"integer(11)"`	
	File_Name string `gorm:"varchar(50);not null"`
	File_Path string `gorm:"varchar(50);not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
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
