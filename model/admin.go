package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
}

type AminBind struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
