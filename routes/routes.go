package routes

import (
	"github.com/aruncs31s/esdcusermodule/handler"
	"github.com/aruncs31s/esdcusermodule/middleware"
	"github.com/gin-gonic/gin"
)

// RegisterPublicUserRoutes registers public user routes (no authentication required)
func RegisterPublicUserRoutes(r *gin.Engine, userHandler handler.UserHandler) {
	userGroup := r.Group("/api/users")
	{
		userGroup.GET("/search", userHandler.SearchUsers)
	}
}

// RegisterProtectedUserRoutes registers protected user routes (authentication required)
func RegisterProtectedUserRoutes(r *gin.Engine, userHandler handler.UserHandler, authMiddleware *middleware.AuthMiddleware) {
	userGroup := r.Group("/api/users")
	userGroup.Use(authMiddleware.RequireAuth())
	{
		userGroup.GET("/me", userHandler.GetCurrentUser)
	}
}

// RegisterOAuthRoutes registers OAuth/SSO authentication routes
func RegisterOAuthRoutes(r *gin.Engine, oauthHandler handler.OAuthHandler) {
	authGroup := r.Group("/api/auth")
	{
		// OAuth provider endpoints
		authGroup.GET("/providers", oauthHandler.GetProviders)
		authGroup.GET("/oauth/:provider", oauthHandler.InitiateAuth)
		authGroup.GET("/oauth/:provider/callback", oauthHandler.HandleCallback)
		authGroup.POST("/token/refresh", oauthHandler.RefreshToken)
	}
}

	