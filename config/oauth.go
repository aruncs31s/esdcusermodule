// Package config provides configuration management for the user module.
package config

import "github.com/aruncs31s/esdcusermodule/domain"

// OAuthConfig contains configuration for OAuth/SSO providers
type OAuthConfig struct {
	// Providers contains the configuration for each OAuth provider
	Providers map[domain.AuthProvider]ProviderConfig
	// JWTSecret is the secret key used to sign JWT tokens
	JWTSecret string
	// JWTExpiry is the token expiry duration in seconds
	JWTExpiry int64
	// JWTIssuer is the issuer identifier for JWT tokens
	JWTIssuer string
}

// ProviderConfig contains configuration for a single OAuth provider
type ProviderConfig struct {
	// ClientID is the OAuth client identifier
	ClientID string
	// ClientSecret is the OAuth client secret
	ClientSecret string
	// RedirectURL is the callback URL after OAuth authentication
	RedirectURL string
	// AuthURL is the authorization endpoint URL
	AuthURL string
	// TokenURL is the token endpoint URL
	TokenURL string
	// UserInfoURL is the user info endpoint URL
	UserInfoURL string
	// Scopes are the OAuth scopes to request
	Scopes []string
	// Enabled indicates if this provider is enabled
	Enabled bool
}

// DefaultOAuthConfig returns a default OAuth configuration
func DefaultOAuthConfig() *OAuthConfig {
	return &OAuthConfig{
		Providers: make(map[domain.AuthProvider]ProviderConfig),
		JWTExpiry: 3600,
		JWTIssuer: "esdcusermodule",
	}
}

// NewGoogleProvider creates a Google OAuth provider configuration
func NewGoogleProvider(clientID, clientSecret, redirectURL string) ProviderConfig {
	return ProviderConfig{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		AuthURL:      "https://accounts.google.com/o/oauth2/v2/auth",
		TokenURL:     "https://oauth2.googleapis.com/token",
		UserInfoURL:  "https://www.googleapis.com/oauth2/v2/userinfo",
		Scopes:       []string{"openid", "email", "profile"},
		Enabled:      true,
	}
}

// NewGitHubProvider creates a GitHub OAuth provider configuration
func NewGitHubProvider(clientID, clientSecret, redirectURL string) ProviderConfig {
	return ProviderConfig{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		AuthURL:      "https://github.com/login/oauth/authorize",
		TokenURL:     "https://github.com/login/oauth/access_token",
		UserInfoURL:  "https://api.github.com/user",
		Scopes:       []string{"user:email", "read:user"},
		Enabled:      true,
	}
}

// NewMicrosoftProvider creates a Microsoft SSO provider configuration
func NewMicrosoftProvider(clientID, clientSecret, redirectURL, tenantID string) ProviderConfig {
	baseURL := "https://login.microsoftonline.com/" + tenantID
	return ProviderConfig{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		AuthURL:      baseURL + "/oauth2/v2.0/authorize",
		TokenURL:     baseURL + "/oauth2/v2.0/token",
		UserInfoURL:  "https://graph.microsoft.com/v1.0/me",
		Scopes:       []string{"openid", "email", "profile", "User.Read"},
		Enabled:      true,
	}
}

// NewOIDCProvider creates a generic OpenID Connect provider configuration
func NewOIDCProvider(clientID, clientSecret, redirectURL, authURL, tokenURL, userInfoURL string, scopes []string) ProviderConfig {
	return ProviderConfig{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		AuthURL:      authURL,
		TokenURL:     tokenURL,
		UserInfoURL:  userInfoURL,
		Scopes:       scopes,
		Enabled:      true,
	}
}

// AddProvider adds a provider configuration to the OAuth config
func (c *OAuthConfig) AddProvider(provider domain.AuthProvider, config ProviderConfig) {
	c.Providers[provider] = config
}

// GetProvider retrieves a provider configuration by type
func (c *OAuthConfig) GetProvider(provider domain.AuthProvider) (ProviderConfig, bool) {
	config, ok := c.Providers[provider]
	return config, ok
}

// IsProviderEnabled checks if a provider is configured and enabled
func (c *OAuthConfig) IsProviderEnabled(provider domain.AuthProvider) bool {
	config, ok := c.Providers[provider]
	return ok && config.Enabled
}
