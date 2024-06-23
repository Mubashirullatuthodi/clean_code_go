package admin

import "github.com/mubzz/clean/model"

type AdminRepository interface {
	Create(admin *model.Admin) (*model.Admin, error)
}
