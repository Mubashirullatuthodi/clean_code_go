package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mubzz/clean/admin"
	"github.com/mubzz/clean/model"
	"github.com/mubzz/clean/utils"
)

type AdminHandler struct {
	adminUsecase admin.AdminUsecase
}

func CreateAdminHandler(r *gin.Engine, adminUsecase admin.AdminUsecase) {
	UserHandler := AdminHandler{adminUsecase: adminUsecase}

	r.POST("/admin", UserHandler.CreateAdmin)
}

func (e *AdminHandler) CreateAdmin(c *gin.Context) {
	var admin = model.AminBind{}
	err := c.BindJSON(&admin)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, "failed to bind")
		return
	}

	Admin := model.Admin{
		Email:    admin.Email,
		Password: admin.Password,
	}

	newAdmin, err := e.adminUsecase.Create(&Admin)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "failed to create")
		return
	}

	utils.HandleSuccess(c, newAdmin)
}
