package model

import "time"

type OTP struct {
	ID     uint      `gorm:"primarykey" json:"id"`
	Otp    string    `json:"otp"`
	Email  string    `gorm:"unique" json:"email"`
	Exp    time.Time //OTP expiry time
	UserID uint      //Foreign key referencing the user model
}

type OtpBind struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}
