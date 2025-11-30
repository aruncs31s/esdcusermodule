package repository

import (
	model "github.com/aruncs31s/esdcmodels"
)

func (r *userRepository) GetAllUsersEssentials() (*[]model.User, error) {
	return r.reader.GetAllUsersEssentials()
}
func (r *userRepositoryReader) GetAllUsersEssentials() (*[]model.User, error) {
	var users []model.User
	err := r.db.Select(model.User{}.GetEssentials()).Find(&users).Error
	return &users, err
}
