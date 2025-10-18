package user

import (
	"github.com/aruncs31s/esdcusermodule/handler"
	"github.com/aruncs31s/esdcusermodule/repository"
	"github.com/aruncs31s/esdcusermodule/routes"
	"github.com/aruncs31s/esdcusermodule/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserModule struct {
	handler handler.UserHandler
}

var userModule *UserModule

func GetUserRepo(db *gorm.DB) repository.UserRepository {
	return repository.NewUserRepository(db)
}

func InitUserModule(db *gorm.DB) error {
	userRepo := GetUserRepo(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	userModule = &UserModule{
		handler: userHandler,
	}
	return nil
}
func RegisterPublicUserRoutes(r *gin.Engine) {
	if userModule == nil {
		panic("User module not initialized")
	}
	routes.RegisterPublicUserRoutes(r, userModule.handler)
}
