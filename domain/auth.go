// Package domain contains the core business entities and value objects
// following Domain-Driven Design principles.
package domain

import "time"

// TokenClaims represents the claims contained in a JWT token
type TokenClaims struct {
	UserID    UserID
	Email     Email
	Provider  AuthProvider
	ExpiresAt time.Time
	IssuedAt  time.Time
	Issuer    string
	Subject   string
	Audience  []string
}

// IsValid checks if the token claims are still valid
func (c *TokenClaims) IsValid() bool {
	return time.Now().Before(c.ExpiresAt)
}

// OAuthUserInfo represents user information from OAuth providers
type OAuthUserInfo struct {
	ExternalID string
	Email      Email
	Name       string
	Provider   AuthProvider
	AvatarURL  string
	Verified   bool
}

// AuthResult represents the result of an authentication attempt
type AuthResult struct {
	User         *User
	AccessToken  string
	RefreshToken string
	ExpiresIn    int64
	TokenType    string
}
