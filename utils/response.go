package utils

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mubzz/clean/model"
)

func HandleSuccess(c *gin.Context, data interface{}) {
	responsedata := model.Response{
		Status:  "200",
		Message: "success",
		Data:    data,
	}
	c.JSON(http.StatusOK, responsedata)
}

func HandleError(c *gin.Context, status int, message string) {
	responsedata := model.Response{
		Status:  strconv.Itoa(status),
		Message: message,
	}
	c.JSON(status, responsedata)
}
