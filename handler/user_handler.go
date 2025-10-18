package handler

import (
	"github.com/aruncs31s/esdcusermodule/service"
	"github.com/aruncs31s/responsehelper"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	SearchUsers(c *gin.Context)
}

func NewUserHandler(userService service.UserService) UserHandler {
	responseHelper := responsehelper.NewResponseHelper()
	return &userHandler{
		userService:    userService,
		responseHelper: responseHelper,
	}
}

type userHandler struct {
	userService    service.UserService
	responseHelper responsehelper.ResponseHelper
}

// handlers/user_handler.go
func (h *userHandler) SearchUsers(c *gin.Context) {
	query := c.Query("q")

	if query == "" {
		h.responseHelper.BadRequest(c, "bad request", "bad query")
		return
	}
	users, err := h.userService.SearchUsers(query)
	if err != nil {
		h.responseHelper.InternalError(c, "internal server error", err)
		return
	}

	h.responseHelper.Success(c, map[string]interface{}{"users": users})

}
