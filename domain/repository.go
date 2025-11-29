// Package domain contains repository interfaces following DDD principles.
package domain

// UserRepository defines the interface for user persistence operations.
// This interface is part of the domain layer and defines the contract
// for data access without implementation details.
type UserRepository interface {
	// FindByID finds a user by their unique identifier
	FindByID(id UserID) (*User, error)
	// FindByEmail finds a user by their email address
	FindByEmail(email Email) (*User, error)
	// FindByExternalID finds a user by their external provider ID
	FindByExternalID(provider AuthProvider, externalID string) (*User, error)
	// Create creates a new user in the repository
	Create(user *User) error
	// Update updates an existing user
	Update(user *User) error
	// Delete removes a user from the repository
	Delete(id UserID) error
}

// CredentialsRepository defines the interface for credentials persistence operations.
type CredentialsRepository interface {
	// FindByUserID finds credentials by user ID
	FindByUserID(userID UserID) (*UserCredentials, error)
	// FindByProvider finds credentials by user ID and provider
	FindByProvider(userID UserID, provider AuthProvider) (*UserCredentials, error)
	// Create creates new credentials
	Create(credentials *UserCredentials) error
	// Update updates existing credentials
	Update(credentials *UserCredentials) error
	// Delete removes credentials
	Delete(userID UserID, provider AuthProvider) error
}
