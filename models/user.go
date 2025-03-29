package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       string `json:"ID"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}
