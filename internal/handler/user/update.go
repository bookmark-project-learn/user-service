package user_handler

import (
	"net/http"

	"github.com/bookmark-project-learn/bookmark-common-libs/pkg/request_ultils"
	userModel "github.com/bookmark-project-learn/user-service/internal/models/dto/api/user"
	"github.com/gin-gonic/gin"
)

type updateUserReposonse struct {
	Message string `json:"message"`
}

// UpdateUserInfo godoc
// @Summary Edit current user
// @Description Edit current user
// @Tags user
// @Accept json
// @Produce json
// @Param input body userModel.UpdateUserInput true "Input required"
// @Success 200 {object} updateUserReposonse
// @Failure      400  {object}  updateUserReposonse
// @Failure      500  {object}  updateUserReposonse
// @Router /v1/users [put]
// @Security BearerAuth
func (u *userHandler) UpdateUserInfo(c *gin.Context) {
	request := &userModel.UpdateUserInput{}
	c.ShouldBindJSON(request)
	response := &updateUserReposonse{Message: "Edit current user successfully"}
	userId, err := request_ultils.GetSubjectFromClaims(c)
	if err != nil {
		response.Message = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}
	err = u.svc.UpdateUserInfo(c, userId, request)
	if err != nil {
		response.Message = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}
	c.JSON(http.StatusOK, response)
}
