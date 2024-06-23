package repository

import (
	"fmt"

	"github.com/mubzz/clean/admin"
	"github.com/mubzz/clean/model"
	"gorm.io/gorm"
)

type AdminRepoImpl struct {
	DB *gorm.DB
}

func NewAdminRepo(DB *gorm.DB) admin.AdminRepository {
	return &AdminRepoImpl{DB: DB}
}

func (e *AdminRepoImpl) Create(admin *model.Admin) (*model.Admin, error) {
	err := e.DB.Create(&admin).Error
	if err != nil {
		return nil, fmt.Errorf("failed to create admin")
	}
	return admin, nil
}
