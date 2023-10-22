package handler

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/haakaashs/antino-labs/models"
	"github.com/haakaashs/antino-labs/resources"
	"github.com/haakaashs/antino-labs/service"
)

type UserHandler interface {
	CreateUser(*gin.Context)
	GetUserById(*gin.Context)
	GetUsers(*gin.Context)
	UserLogin(*gin.Context)
	DeleteUserById(*gin.Context)
}

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *userHandler {
	return &userHandler{
		userService: userService,
	}
}

func (h *userHandler) CreateUser(ctx *gin.Context) {
	funcdesc := "CreateUser"
	log.Println("enter handeler" + funcdesc)

	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := h.userService.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("exit " + funcdesc)
	ctx.JSON(http.StatusOK, gin.H{"user_id": userId})
}

func (h *userHandler) GetUserById(ctx *gin.Context) {
	funcdesc := "GetUserById"
	log.Println("enter handeler" + funcdesc)

	userId, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.GetUserById(uint64(userId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("exit " + funcdesc)
	ctx.JSON(http.StatusOK, user)
}

func (h *userHandler) GetUsers(ctx *gin.Context) {
	funcdesc := "GetUsers"
	log.Println("enter handeler" + funcdesc)

	users, err := h.userService.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("exit " + funcdesc)
	ctx.JSON(http.StatusOK, users)
}

func (h *userHandler) UserLogin(ctx *gin.Context) {
	funcdesc := "UserLogin"
	log.Println("enter handeler" + funcdesc)

	var credentials resources.UserCredential
	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.userService.UserLogin(credentials)
	if err != nil {
		if strings.Contains(err.Error(), "incorrect") {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("exit " + funcdesc)
	ctx.JSON(http.StatusOK, gin.H{"Status": "Login Successful"})
}

func (h *userHandler) DeleteUserById(ctx *gin.Context) {
	funcdesc := "DeleteUserById"
	log.Println("enter handeler" + funcdesc)

	userId, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.userService.DeleteUserById(uint64(userId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("exit " + funcdesc)
	ctx.JSON(http.StatusOK, gin.H{"Status": "Deletion Successful"})
}
