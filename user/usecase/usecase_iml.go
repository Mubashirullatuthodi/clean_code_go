package usecase

import (
	"github.com/mubzz/clean/model"
	"github.com/mubzz/clean/user"
)

type UserUsecaseIml struct {
	userRepo user.UserRepository
}

func CreateUserUsecase(userRepo user.UserRepository) user.UserUsecase {
	return &UserUsecaseIml{userRepo: userRepo}
}

func (e *UserUsecaseIml) SignUp(user *model.OTP) error {
	return e.userRepo.SignUp(user)
}

func (e *UserUsecaseIml) PostOtp(user *model.User) (*model.User, error) {
	return e.userRepo.PostOtp(user)
}

func (e *UserUsecaseIml) OtpVerify(otp string, otpstore *model.OTP) (*model.OTP, error) {
	return e.userRepo.OtpVerify(otp, otpstore)
}

func (e *UserUsecaseIml) DeleteOtp(email, otp string,otpstore *model.OTP) error {
	return e.userRepo.DeleteOtp(email, otp,otpstore)
}
