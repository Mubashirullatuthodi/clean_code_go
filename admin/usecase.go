package admin

import "github.com/mubzz/clean/model"

type AdminUsecase interface {
	Create(admin *model.Admin)(*model.Admin,error)
}
