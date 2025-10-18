package user

import (
	"github.com/aruncs31s/esdcusermodule/repository"
	"gorm.io/gorm"
)

func GetUserRepo(db *gorm.DB) repository.UserRepository {
	return repository.NewUserRepository(db)
}
