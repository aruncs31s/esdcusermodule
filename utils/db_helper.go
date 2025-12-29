package utils

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDSN() string {

	host := os.Getenv("USER_DB_HOST")
	port := os.Getenv("USER_DB_PORT")
	user := os.Getenv("USER_DB_USER")
	password := os.Getenv("USER_DB_PASS")
	dbname := os.Getenv("USER_DB_NAME")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user,
		password,
		host,
		port,
		dbname,
	)
	fmt.Println(dsn)
	return dsn
}
func GetDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(GetDSN()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
