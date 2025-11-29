package user

import (
	"github.com/aruncs31s/esdcusermodule/config"
	"github.com/aruncs31s/esdcusermodule/domain"
	"github.com/aruncs31s/esdcusermodule/handler"
	"github.com/aruncs31s/esdcusermodule/middleware"
	"github.com/aruncs31s/esdcusermodule/repository"
	"github.com/aruncs31s/esdcusermodule/routes"
	"github.com/aruncs31s/esdcusermodule/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserModule encapsulates the user module components following DDD principles
type UserModule struct {
	userHandler    handler.UserHandler
	oauthHandler   handler.OAuthHandler
	authMiddleware *middleware.AuthMiddleware
	authService    service.AuthService
	oauthService   service.OAuthService
}

var userModule *UserModule

// GetUserRepo returns a new UserRepository instance
func GetUserRepo(db *gorm.DB) repository.UserRepository {
	return repository.NewUserRepository(db)
}

// InitUserModule initializes the user module with basic functionality
func InitUserModule(db *gorm.DB) error {
	userRepo := GetUserRepo(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	userModule = &UserModule{
		userHandler: userHandler,
	}
	return nil
}

// InitUserModuleWithOAuth initializes the user module with OAuth/SSO support
func InitUserModuleWithOAuth(db *gorm.DB, oauthCfg *config.OAuthConfig) error {
	userRepo := GetUserRepo(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Initialize auth services
	authService := service.NewAuthService(userRepo, oauthCfg)
	oauthService := service.NewOAuthService(userRepo, oauthCfg, authService)

	// Create auth middleware
	authMiddleware := middleware.NewAuthMiddleware(authService)

	// Get enabled providers from config
	enabledProviders := getEnabledProviders(oauthCfg)

	// Create OAuth handler
	oauthHandler := handler.NewOAuthHandler(oauthService, authService, enabledProviders)

	userModule = &UserModule{
		userHandler:    userHandler,
		oauthHandler:   oauthHandler,
		authMiddleware: authMiddleware,
		authService:    authService,
		oauthService:   oauthService,
	}
	return nil
}

// getEnabledProviders returns a list of enabled OAuth providers from the config
func getEnabledProviders(cfg *config.OAuthConfig) []domain.AuthProvider {
	providers := make([]domain.AuthProvider, 0)
	for provider, providerCfg := range cfg.Providers {
		if providerCfg.Enabled {
			providers = append(providers, provider)
		}
	}
	return providers
}

// RegisterPublicUserRoutes registers public user routes (no authentication required)
func RegisterPublicUserRoutes(r *gin.Engine) {
	if userModule == nil {
		panic("User module not initialized")
	}
	routes.RegisterPublicUserRoutes(r, userModule.userHandler)
}

// RegisterProtectedUserRoutes registers protected user routes (authentication required)
func RegisterProtectedUserRoutes(r *gin.Engine) {
	if userModule == nil {
		panic("User module not initialized")
	}
	if userModule.authMiddleware == nil {
		panic("Auth middleware not initialized. Use InitUserModuleWithOAuth to enable protected routes")
	}
	routes.RegisterProtectedUserRoutes(r, userModule.userHandler, userModule.authMiddleware)
}

// RegisterOAuthRoutes registers OAuth/SSO authentication routes
func RegisterOAuthRoutes(r *gin.Engine) {
	if userModule == nil {
		panic("User module not initialized")
	}
	if userModule.oauthHandler == nil {
		panic("OAuth handler not initialized. Use InitUserModuleWithOAuth to enable OAuth routes")
	}
	routes.RegisterOAuthRoutes(r, userModule.oauthHandler)
}

// GetAuthMiddleware returns the auth middleware for use in other modules
func GetAuthMiddleware() *middleware.AuthMiddleware {
	if userModule == nil {
		panic("User module not initialized")
	}
	return userModule.authMiddleware
}

// GetAuthService returns the auth service for use in other modules
func GetAuthService() service.AuthService {
	if userModule == nil {
		panic("User module not initialized")
	}
	return userModule.authService
}
