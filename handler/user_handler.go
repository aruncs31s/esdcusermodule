package handler

import (
	"strconv"

	"github.com/aruncs31s/esdcusermodule/service"
	"github.com/aruncs31s/esdcusermodule/utils"
	"github.com/aruncs31s/responsehelper"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	SearchUsers(c *gin.Context)
	ListAllUsers(c *gin.Context)
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

func (h *userHandler) ListAllUsers(c *gin.Context) {
	rc := getEssential(c)
	if rc == nil {
		h.responseHelper.BadRequest(c, "bad request", "invalid pagination parameters")
		return
	}
	users, err := h.userService.ListAllUsers(rc)
	if err != nil {
		h.responseHelper.InternalError(c, "internal server error", err)
		return
	}

	h.responseHelper.Success(c, users)
}

func getEssential(c *gin.Context) *utils.HTTPRequestContext {
	// Extract pagination parameters from query
	limitParam := c.DefaultQuery("limit", "10")
	offsetParam := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit < 0 {
		return nil
	}

	offset, err := strconv.Atoi(offsetParam)
	if err != nil || offset < 0 {
		return nil
	}

	rc := &utils.HTTPRequestContext{
		Parameters: &utils.Parameters{
			Limit:  limit,
			Offset: offset,
		},
	}
	return rc
}
