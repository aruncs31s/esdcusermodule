// Package domain contains the core business entities and value objects
// following Domain-Driven Design principles.
package domain

import "time"

// UserID represents a unique identifier for a user (value object)
type UserID uint

// Email represents a validated email address (value object)
type Email string

// User represents the core user entity in the domain layer.
// This entity encapsulates user data and business rules.
type User struct {
	ID        UserID
	Name      string
	Email     Email
	Username  string
	Provider  AuthProvider
	ExternalID string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// AuthProvider represents the authentication provider type (value object)
type AuthProvider string

const (
	// ProviderLocal represents local authentication
	ProviderLocal AuthProvider = "local"
	// ProviderGoogle represents Google OAuth authentication
	ProviderGoogle AuthProvider = "google"
	// ProviderGitHub represents GitHub OAuth authentication
	ProviderGitHub AuthProvider = "github"
	// ProviderMicrosoft represents Microsoft SSO authentication
	ProviderMicrosoft AuthProvider = "microsoft"
	// ProviderOIDC represents generic OpenID Connect authentication
	ProviderOIDC AuthProvider = "oidc"
)

// IsValid checks if the auth provider is a known valid provider
func (p AuthProvider) IsValid() bool {
	switch p {
	case ProviderLocal, ProviderGoogle, ProviderGitHub, ProviderMicrosoft, ProviderOIDC:
		return true
	default:
		return false
	}
}

// UserCredentials represents authentication credentials for a user
type UserCredentials struct {
	UserID       UserID
	Provider     AuthProvider
	ExternalID   string
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
}

// IsExpired checks if the credentials have expired
func (c *UserCredentials) IsExpired() bool {
	return time.Now().After(c.ExpiresAt)
}
