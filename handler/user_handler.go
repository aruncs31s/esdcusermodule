package handler

import (
	"github.com/aruncs31s/esdcusermodule/middleware"
	"github.com/aruncs31s/esdcusermodule/service"
	"github.com/aruncs31s/responsehelper"
	"github.com/gin-gonic/gin"
)

// UserHandler defines the interface for user-related HTTP handlers
type UserHandler interface {
	// SearchUsers handles user search requests
	SearchUsers(c *gin.Context)
	// GetCurrentUser returns the currently authenticated user's profile
	GetCurrentUser(c *gin.Context)
}

// NewUserHandler creates a new UserHandler instance
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

// SearchUsers handles user search requests
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

// GetCurrentUser returns the currently authenticated user's profile
func (h *userHandler) GetCurrentUser(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		h.responseHelper.Unauthorized(c, "user not authenticated")
		return
	}

	user, err := h.userService.GetUserByID(uint(userID))
	if err != nil {
		h.responseHelper.InternalError(c, "failed to get user", err)
		return
	}

	h.responseHelper.Success(c, map[string]interface{}{"user": user})
}
