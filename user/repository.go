package user

import "github.com/mubzz/clean/model"

type UserRepository interface {
	SignUp(user *model.OTP) error
	PostOtp(user *model.User) (*model.User, error)
	OtpVerify(otp string, otpstore *model.OTP) (*model.OTP, error)
	DeleteOtp(email, otp string,otpstore *model.OTP) error
}
