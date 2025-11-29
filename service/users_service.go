package service

import (
	"github.com/aruncs31s/esdcusermodule/dto"
	"github.com/aruncs31s/esdcusermodule/repository"
)

// UserService defines the interface for user-related business operations
type UserService interface {
	// SearchUsers searches for users by query string
	SearchUsers(query string) ([]dto.UserResponse, error)
	// GetUserByID retrieves a user by their ID
	GetUserByID(id uint) (*dto.UserResponse, error)
}

type userService struct {
	userRepo repository.UserRepository
}

// NewUserService creates a new UserService instance
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

// SearchUsers searches for users by query string
func (s *userService) SearchUsers(query string) ([]dto.UserResponse, error) {
	users, err := s.userRepo.SearchUsers(query)
	if err != nil {
		return nil, err
	}
	filteredUsers := make([]dto.UserResponse, 0)
	for _, user := range *users {
		filteredUsers = append(filteredUsers, dto.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return filteredUsers, nil
}

// GetUserByID retrieves a user by their ID
func (s *userService) GetUserByID(id uint) (*dto.UserResponse, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return &dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
