package models

import (
	"time"

	"gorm.io/gorm"

)

type Post struct {
    ID        uint `gorm:"primaryKey"`
    Title     string
    Content   string
    Category  string
    CreatedAt time.Time
    UpdatedAt time.Time
    Status    string
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
