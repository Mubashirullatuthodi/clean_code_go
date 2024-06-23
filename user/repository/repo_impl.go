package repository

import (
	"fmt"

	"github.com/mubzz/clean/model"
	"github.com/mubzz/clean/user"
	"gorm.io/gorm"
)

type UserRepoImpl struct {
	DB *gorm.DB
}

func NewUserRepo(DB *gorm.DB) user.UserRepository {
	return &UserRepoImpl{DB: DB}
}

func (e *UserRepoImpl) SignUp(user *model.OTP) error {
	err := e.DB.Create(&user).Error
	if err != nil {
		return fmt.Errorf("failed to create otp")
	}
	return nil
}

func (e *UserRepoImpl) PostOtp(user *model.User) (*model.User, error) {
	err := e.DB.Create(&user).Error
	if err != nil {
		return nil, fmt.Errorf("failed to create user")
	}
	return user, nil
}

func (e *UserRepoImpl) OtpVerify(otp string, otpstore *model.OTP) (*model.OTP, error) {
	if err := e.DB.Where("otp=?", otp).First(&otpstore).Error; err != nil {
		return nil, fmt.Errorf("invalid OTP")
	}
	return otpstore, nil
}

func (e *UserRepoImpl) DeleteOtp(email, otp string,otpstore *model.OTP) error {
	if err := e.DB.Where("otp=? AND email=?", otp, email).Delete(&otpstore).Error; err != nil {
		return fmt.Errorf("failed to delete")
	}
	return nil
}
