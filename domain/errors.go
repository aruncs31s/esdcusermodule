// Package domain contains domain-level errors following DDD principles.
package domain

import "errors"

// Domain-level errors for user operations
var (
	// ErrUserNotFound indicates that the requested user was not found
	ErrUserNotFound = errors.New("user not found")
	// ErrUserAlreadyExists indicates that a user with the same identifier already exists
	ErrUserAlreadyExists = errors.New("user already exists")
	// ErrInvalidEmail indicates that the email format is invalid
	ErrInvalidEmail = errors.New("invalid email format")
	// ErrInvalidCredentials indicates that the provided credentials are invalid
	ErrInvalidCredentials = errors.New("invalid credentials")
)

// Domain-level errors for authentication operations
var (
	// ErrTokenExpired indicates that the authentication token has expired
	ErrTokenExpired = errors.New("token expired")
	// ErrTokenInvalid indicates that the authentication token is invalid
	ErrTokenInvalid = errors.New("invalid token")
	// ErrProviderNotSupported indicates that the OAuth provider is not supported
	ErrProviderNotSupported = errors.New("OAuth provider not supported")
	// ErrProviderNotEnabled indicates that the OAuth provider is not enabled
	ErrProviderNotEnabled = errors.New("OAuth provider not enabled")
	// ErrOAuthStateMismatch indicates that the OAuth state parameter does not match
	ErrOAuthStateMismatch = errors.New("OAuth state mismatch")
	// ErrOAuthCodeExchangeFailed indicates that the OAuth code exchange failed
	ErrOAuthCodeExchangeFailed = errors.New("OAuth code exchange failed")
	// ErrRefreshTokenInvalid indicates that the refresh token is invalid
	ErrRefreshTokenInvalid = errors.New("invalid refresh token")
)
