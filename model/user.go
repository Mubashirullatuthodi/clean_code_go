package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `gorm:"not null" json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `gorm:"unique;not null" json:"email"`
	Gender    string `gorm:"check:gender IN ('male','MALE','female','FEMALE','')" json:"gender"`
	Phone     string `gorm:"not null" json:"phone_no"`
	Password  string `gorm:"not null" json:"password"`
	Status    string `gorm:"default:Active" json:"status"`
}

type UserBind struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	Phone     string `json:"phone_no"`
	Password  string `json:"password"`
	Status    string `json:"status"`
}
