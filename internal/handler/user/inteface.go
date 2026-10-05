package user_handler

import (
	user_service "github.com/bookmark-project-learn/user-service/internal/service/user"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	GetUserInfo(c *gin.Context)
	UpdateUserInfo(c *gin.Context)
}

type userHandler struct {
	svc user_service.UserService
}

func NewUserHandler(svc user_service.UserService) UserHandler {
	return &userHandler{svc: svc}
}

var createSuccessMessage = "User created successfully!"
