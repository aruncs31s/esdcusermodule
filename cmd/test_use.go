package main

import (
	model "github.com/aruncs31s/esdcmodels"
	"github.com/aruncs31s/esdcusermodule/utils"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {

	godotenv.Load(".env")
}
func main() {
	db := utils.GetDB()
	db.AutoMigrate(model.User{})
	// userRepo := repository.NewUserRepository(db)
	// userService := service.NewUserService(userRepo)
	// userHandler := handler.NewUserHandler(userService)
}
