package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Username  string `gorm:"varchar(50);not null"`
	Email     string `gorm:"varchar(50);not null"`
	Password  string `gorm:"varchar(50);not null"`
	Profile   string `gorm:"varchar(50); null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
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

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	if err != nil {
		return err
	}
	return nil
}
