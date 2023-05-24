package models

import "gorm.io/gorm"

type User struct {
	ID       *int   `gorm:"primaryKey;autoIncrement"`
	Username string `gorm:"size:100;unique"`
	Email    string `gorm:"size:100;unique"`
}

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&User{}); err != nil {
		return err
	}
	return nil
}
