package user_handler

import (
	"net/http"

	request_ultils "github.com/bookmark-project-learn/bookmark-common-libs/pkg/request_ultils"
	userModel "github.com/bookmark-project-learn/user-service/internal/models/dto/api/user"
	user_service "github.com/bookmark-project-learn/user-service/internal/service/user"
	"github.com/gin-gonic/gin"
)

// LoginLink Login
// @Summary Login
// @Description Login
// @Tags user
// @Accept json
// @Produce json
// @Param input body userModel.UserLogin true "Input required"
// @Success 200 {object} loginResponse
// @Failure      400  {object}  loginResponse
// @Failure      500  {object}  loginResponse
// @Router /v1/users/login [post]
func (u *userHandler) Login(c *gin.Context) {
	userRequest, err := request_ultils.ModelBindValidation[userModel.UserLogin](c)
	loginResponse := &loginResponse{}
	if err != nil {
		loginResponse.Message = err.Error()
		c.JSON(http.StatusBadRequest, loginResponse)
		return
	}

	token, err := u.svc.Login(c, *userRequest)

	if err != nil {
		if user_service.CheckErrorIsServiceErr(err) {
			loginResponse.Message = err.Error()
			c.JSON(http.StatusBadRequest, loginResponse)
			return
		}
		loginResponse.Message = "Internal server error"
		c.JSON(http.StatusInternalServerError, loginResponse)
		return
	}
	loginResponse.Data = token
	loginResponse.Message = "Login successfully"
	c.JSON(http.StatusOK, loginResponse)
}

type loginResponse struct {
	Data    string `json:"data"`
	Message string `json:"message"`
}
