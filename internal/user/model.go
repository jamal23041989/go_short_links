package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Password string
	Email    string `gorm:"index"`
}
