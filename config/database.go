package config

import (
	"fmt"

	"github.com/codepnw/gin-gorm/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host = "localhost"
	port = 5433
	user = "root"
	password = "123456"
	dbName = "gin-gorm"
)

func DatabaseConnection() *gorm.DB {
	sqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName,
	)

	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	helper.ErrorPanic(err)

	return db
}