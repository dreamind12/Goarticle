package models

import (
	"time"

	"gorm.io/gorm"

)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Username  string
	Email     string
	Password  string
	Profile   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) Tablename() string {
	return "users"
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	tx.Statement.SetColumn("CreatedAt", time.Now())
	return
}

func (user *User) BeforeUpdate(tx *gorm.DB) (err error) {
	tx.Statement.SetColumn("UpdatedAt", time.Now())
	return
}
