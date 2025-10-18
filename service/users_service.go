package service

import (
	"github.com/aruncs31s/esdcusermodule/dto"
	"github.com/aruncs31s/esdcusermodule/repository"
)

type UserService interface {
	SearchUsers(query string) ([]dto.UserResponse, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) SearchUsers(query string) ([]dto.UserResponse, error) {
	users, err := s.userRepo.SearchUsers(query)
	if err != nil {
		return nil, err
	}
	filteredUsers := make([]dto.UserResponse, 0)
	for _, user := range users {
		filteredUsers = append(filteredUsers, dto.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return filteredUsers, nil
}
