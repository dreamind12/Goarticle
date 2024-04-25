package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        uint `gorm:"primaryKey"`
	User_ID   uint `gorm:"not null"`
	UserUpdate_ID uint `gorm:"default:null"`
	Title     string `gorm:"not null"`
	Descript  string  `gorm:"type:text;not null"`
	Category  string	`gorm:"default:null"`
	Status    string	`gorm:"default:null"`
	CreatedAt time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:datetime;default:null"`
}

func (Post) TableName() string {
	return "posts"
}

func (post *Post) BeforeCreate(tx *gorm.DB) (err error) {
	tx.Statement.SetColumn("CreatedAt", time.Now())
	return
}

func (post *Post) BeforeUpdate(tx *gorm.DB) (err error) {
	tx.Statement.SetColumn("UpdatedAt", time.Now())
	return
}

func MigratePosts(db *gorm.DB) error {
	err := db.AutoMigrate(&Post{})
	if err != nil {
		return err
	}
	return nil
}
