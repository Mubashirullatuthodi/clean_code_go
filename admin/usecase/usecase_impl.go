package usecase

import (
	"github.com/mubzz/clean/admin"
	"github.com/mubzz/clean/model"
)

type UseCaseImpl struct {
	adminRepo admin.AdminRepository
}

func NewAdmin(adminRepo admin.AdminRepository) admin.AdminUsecase {
	return &UseCaseImpl{adminRepo: adminRepo}
}

func (e *UseCaseImpl) Create(admin *model.Admin) (*model.Admin, error) {
	return e.adminRepo.Create(admin)
}
