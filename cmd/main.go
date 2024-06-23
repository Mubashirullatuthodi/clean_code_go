package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	adminHandler "github.com/mubzz/clean/admin/handler"
	adminRepo "github.com/mubzz/clean/admin/repository"
	adminUsecase "github.com/mubzz/clean/admin/usecase"
	"github.com/mubzz/clean/config"
	"github.com/mubzz/clean/model"
	userHandler "github.com/mubzz/clean/user/handler"
	userRepo "github.com/mubzz/clean/user/repository"
	userUsecase "github.com/mubzz/clean/user/usecase"
)

func main() {
	r := gin.Default()

	db := config.ConnectDB()
	if err := config.LoadEnvVariables(); err != nil {
		fmt.Println("counldn't load env variables err:", err)
	}

	userRepo := userRepo.NewUserRepo(db)
	userUsecase := userUsecase.CreateUserUsecase(userRepo)
	userHandler.CreateUserHandler(r, userUsecase)

	adminRepo := adminRepo.NewAdminRepo(db)
	adminUsecase := adminUsecase.NewAdmin(adminRepo)
	adminHandler.CreateAdminHandler(r, adminUsecase)

	db.AutoMigrate(&model.User{}, &model.Admin{}, &model.OTP{})

	r.Run()
}
