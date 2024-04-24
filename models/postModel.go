package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        uint `gorm:"primaryKey"`
	User_ID   uint `gorm:"integer(11);not null"`
	Title     string `gorm:"varchar(50);not null"`
	Descript  string  `gorm:"type:text;not null"`
	Category  string	`gorm:"varchar(20)"`
	Status    string	`gorm:"varchar(10)"`	
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
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
