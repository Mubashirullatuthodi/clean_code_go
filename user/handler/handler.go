package handler

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mubzz/clean/model"
	"github.com/mubzz/clean/pkg/authotp"
	"github.com/mubzz/clean/pkg/hashpassword"
	"github.com/mubzz/clean/user"
	"github.com/mubzz/clean/utils"
)

type UserHandler struct {
	userusecase user.UserUsecase
	dataChannel chan model.UserBind
	dataStore   map[string]model.UserBind
	mu          sync.Mutex
}

func CreateUserHandler(r *gin.Engine, Userusecase user.UserUsecase) {
	UserHandler := UserHandler{
		userusecase: Userusecase,
		dataChannel: make(chan model.UserBind),
		dataStore:   make(map[string]model.UserBind),
	}

	go UserHandler.dataListener()

	r.POST("/user", UserHandler.AddUser)
	r.POST("/user/verifyotp", UserHandler.PostOtp)
}

func (e *UserHandler) dataListener() {
	for user := range e.dataChannel {
		e.mu.Lock()
		// Using email as the key for simplicity, but a UUID or some other identifier would be better.
		e.dataStore[user.Email] = user
		e.mu.Unlock()
	}
}

func (e *UserHandler) AddUser(c *gin.Context) {
	var user = model.UserBind{}
	var exUser = model.User{}
	err := c.Bind(&user)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, "failed to bind")
		return
	}

	hashpassword, err := hashpassword.HashPassword(user.Password)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "failed to hash password")
	}
	user.Password = hashpassword

	otp := authotp.GenerateOtp()

	otpRecord := model.OTP{
		Otp:    otp,
		Email:  user.Email,
		Exp:    time.Now().Add(5 * time.Minute),
		UserID: exUser.ID,
	}

	//sending
	e.dataChannel <- user

	erro := e.userusecase.SignUp(&otpRecord)
	if erro != nil {
		utils.HandleError(c, http.StatusBadRequest, "failed to create user")
		return
	}

	errr := authotp.SendEmail(user.Email, otp)
	if errr != nil {
		utils.HandleError(c, http.StatusInternalServerError, "Failed to send OTP via email")
		return
	}
	message := "email send to your email"

	utils.HandleSuccess(c, message)
}

func (e *UserHandler) PostOtp(c *gin.Context) {
	var input = model.OtpBind{}

	err := c.BindJSON(&input)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, "Failed to bind")
		return
	}

	var otp model.OTP
	fetchedOtp, err := e.userusecase.OtpVerify(input.OTP, &otp)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, "Invalid otp")
		return
	}

	if time.Now().After(fetchedOtp.Exp) {
		utils.HandleError(c, http.StatusBadRequest, "otp expired. Please request new otp.")
		return
	}

	e.mu.Lock()
	exUser, exists := e.dataStore[fetchedOtp.Email]
	e.mu.Unlock()
	if !exists {
		utils.HandleError(c, http.StatusBadRequest, "user data not found")
		return
	}

	new := model.User{
		FirstName: exUser.FirstName,
		LastName:  exUser.LastName,
		Email:     exUser.Email,
		Gender:    exUser.Gender,
		Phone:     exUser.Phone,
		Password:  exUser.Password,
	}

	newUser, err := e.userusecase.PostOtp(&new)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "failed to create user")
		return
	}

	utils.HandleSuccess(c, newUser)

	err = e.userusecase.DeleteOtp(fetchedOtp.Email, fetchedOtp.Otp, fetchedOtp)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "failed to delete otp")
		return
	}

}
