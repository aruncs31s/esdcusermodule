// Package middleware provides HTTP middleware for authentication and authorization.
package middleware

import (
	"net/http"
	"strings"

	"github.com/aruncs31s/esdcusermodule/domain"
	"github.com/gin-gonic/gin"
)

// AuthService defines the interface for authentication operations
type AuthService interface {
	// ValidateToken validates a JWT token and returns the claims
	ValidateToken(token string) (*domain.TokenClaims, error)
}

// AuthMiddleware provides authentication middleware for protected routes
type AuthMiddleware struct {
	authService AuthService
}

// NewAuthMiddleware creates a new authentication middleware instance
func NewAuthMiddleware(authService AuthService) *AuthMiddleware {
	return &AuthMiddleware{
		authService: authService,
	}
}

// RequireAuth returns a middleware that requires valid authentication
func (m *AuthMiddleware) RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := extractBearerToken(c)
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "missing or invalid authorization header",
			})
			return
		}

		claims, err := m.authService.ValidateToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})
			return
		}

		if !claims.IsValid() {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "token expired",
			})
			return
		}

		// Set user information in context for downstream handlers
		c.Set("user_id", claims.UserID)
		c.Set("user_email", claims.Email)
		c.Set("auth_provider", claims.Provider)
		c.Set("token_claims", claims)

		c.Next()
	}
}

// OptionalAuth returns a middleware that allows optional authentication
func (m *AuthMiddleware) OptionalAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := extractBearerToken(c)
		if token == "" {
			c.Next()
			return
		}

		claims, err := m.authService.ValidateToken(token)
		if err == nil && claims.IsValid() {
			c.Set("user_id", claims.UserID)
			c.Set("user_email", claims.Email)
			c.Set("auth_provider", claims.Provider)
			c.Set("token_claims", claims)
		}

		c.Next()
	}
}

// RequireProvider returns a middleware that requires authentication from a specific provider
func (m *AuthMiddleware) RequireProvider(provider domain.AuthProvider) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := extractBearerToken(c)
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "missing or invalid authorization header",
			})
			return
		}

		claims, err := m.authService.ValidateToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})
			return
		}

		if !claims.IsValid() {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "token expired",
			})
			return
		}

		if claims.Provider != provider {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "authentication from required provider is needed",
			})
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("user_email", claims.Email)
		c.Set("auth_provider", claims.Provider)
		c.Set("token_claims", claims)

		c.Next()
	}
}

// extractBearerToken extracts the bearer token from the Authorization header
func extractBearerToken(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return ""
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "bearer") {
		return ""
	}

	return parts[1]
}

// GetUserID retrieves the user ID from the context (set by auth middleware)
func GetUserID(c *gin.Context) (domain.UserID, bool) {
	userID, exists := c.Get("user_id")
	if !exists {
		return 0, false
	}
	id, ok := userID.(domain.UserID)
	return id, ok
}

// GetUserEmail retrieves the user email from the context (set by auth middleware)
func GetUserEmail(c *gin.Context) (domain.Email, bool) {
	email, exists := c.Get("user_email")
	if !exists {
		return "", false
	}
	e, ok := email.(domain.Email)
	return e, ok
}

// GetAuthProvider retrieves the auth provider from the context (set by auth middleware)
func GetAuthProvider(c *gin.Context) (domain.AuthProvider, bool) {
	provider, exists := c.Get("auth_provider")
	if !exists {
		return "", false
	}
	p, ok := provider.(domain.AuthProvider)
	return p, ok
}

// GetTokenClaims retrieves the full token claims from the context
func GetTokenClaims(c *gin.Context) (*domain.TokenClaims, bool) {
	claims, exists := c.Get("token_claims")
	if !exists {
		return nil, false
	}
	tc, ok := claims.(*domain.TokenClaims)
	return tc, ok
}
