package handler

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"github.com/aruncs31s/esdcusermodule/domain"
	"github.com/aruncs31s/esdcusermodule/service"
	"github.com/aruncs31s/responsehelper"
	"github.com/gin-gonic/gin"
)

// OAuthHandler handles OAuth authentication endpoints
type OAuthHandler interface {
	// InitiateAuth initiates the OAuth flow for a provider
	InitiateAuth(c *gin.Context)
	// HandleCallback handles the OAuth callback from a provider
	HandleCallback(c *gin.Context)
	// RefreshToken handles token refresh requests
	RefreshToken(c *gin.Context)
	// GetProviders returns available OAuth providers
	GetProviders(c *gin.Context)
}

type oauthHandler struct {
	oauthService   service.OAuthService
	authService    service.AuthService
	responseHelper responsehelper.ResponseHelper
	enabledProviders []domain.AuthProvider
}

// NewOAuthHandler creates a new OAuth handler instance
func NewOAuthHandler(oauthService service.OAuthService, authService service.AuthService, enabledProviders []domain.AuthProvider) OAuthHandler {
	return &oauthHandler{
		oauthService:     oauthService,
		authService:      authService,
		responseHelper:   responsehelper.NewResponseHelper(),
		enabledProviders: enabledProviders,
	}
}

// InitiateAuth initiates the OAuth flow for a provider
func (h *oauthHandler) InitiateAuth(c *gin.Context) {
	providerStr := c.Param("provider")
	provider := domain.AuthProvider(providerStr)

	if !provider.IsValid() {
		h.responseHelper.BadRequest(c, "invalid provider", "unsupported OAuth provider")
		return
	}

	if !h.isProviderEnabled(provider) {
		h.responseHelper.BadRequest(c, "provider not available", "this OAuth provider is not enabled")
		return
	}

	// Generate state parameter for CSRF protection
	state := generateState()

	authURL, err := h.oauthService.GetAuthorizationURL(provider, state)
	if err != nil {
		h.responseHelper.InternalError(c, "failed to generate authorization URL", err)
		return
	}

	// Set state in session or cookie for verification in callback
	c.SetCookie("oauth_state", state, 600, "/", "", true, true)

	c.Redirect(http.StatusTemporaryRedirect, authURL)
}

// HandleCallback handles the OAuth callback from a provider
func (h *oauthHandler) HandleCallback(c *gin.Context) {
	providerStr := c.Param("provider")
	provider := domain.AuthProvider(providerStr)

	if !provider.IsValid() {
		h.responseHelper.BadRequest(c, "invalid provider", "unsupported OAuth provider")
		return
	}

	code := c.Query("code")
	state := c.Query("state")

	if code == "" {
		errorCode := c.Query("error")
		errorDesc := c.Query("error_description")
		h.responseHelper.BadRequest(c, "OAuth error", errorCode+": "+errorDesc)
		return
	}

	// Verify state parameter
	storedState, err := c.Cookie("oauth_state")
	if err != nil || storedState != state {
		h.responseHelper.BadRequest(c, "invalid state", "OAuth state mismatch")
		return
	}

	// Clear the state cookie
	c.SetCookie("oauth_state", "", -1, "/", "", true, true)

	authResult, err := h.oauthService.HandleCallback(provider, code, state)
	if err != nil {
		h.responseHelper.InternalError(c, "OAuth authentication failed", err)
		return
	}

	h.responseHelper.Success(c, gin.H{
		"access_token":  authResult.AccessToken,
		"refresh_token": authResult.RefreshToken,
		"expires_in":    authResult.ExpiresIn,
		"token_type":    authResult.TokenType,
		"user": gin.H{
			"id":    authResult.User.ID,
			"name":  authResult.User.Name,
			"email": authResult.User.Email,
		},
	})
}

// RefreshToken handles token refresh requests
func (h *oauthHandler) RefreshToken(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		h.responseHelper.BadRequest(c, "invalid request", err.Error())
		return
	}

	authResult, err := h.authService.RefreshToken(req.RefreshToken)
	if err != nil {
		h.responseHelper.Unauthorized(c, "invalid refresh token")
		return
	}

	h.responseHelper.Success(c, gin.H{
		"access_token":  authResult.AccessToken,
		"refresh_token": authResult.RefreshToken,
		"expires_in":    authResult.ExpiresIn,
		"token_type":    authResult.TokenType,
	})
}

// GetProviders returns available OAuth providers
func (h *oauthHandler) GetProviders(c *gin.Context) {
	providers := make([]string, 0, len(h.enabledProviders))
	for _, p := range h.enabledProviders {
		providers = append(providers, string(p))
	}

	h.responseHelper.Success(c, gin.H{
		"providers": providers,
	})
}

func (h *oauthHandler) isProviderEnabled(provider domain.AuthProvider) bool {
	for _, p := range h.enabledProviders {
		if p == provider {
			return true
		}
	}
	return false
}

// generateState generates a cryptographically secure random state parameter for CSRF protection
func generateState() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		panic("failed to generate cryptographically secure random state: " + err.Error())
	}
	return base64.URLEncoding.EncodeToString(b)
}
