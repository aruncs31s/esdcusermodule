package routes

import (
	"github.com/aruncs31s/esdcusermodule/handler"
	"github.com/gin-gonic/gin"
)

func RegisterPublicUserRoutes(r *gin.Engine, userHandler handler.UserHandler) {
	userGroup := r.Group("/api/users")
	{

		userGroup.GET("/search", userHandler.SearchUsers)
		userGroup.GET("", userHandler.ListAllUsers)
	}
}
