package main

import (
	"net/http"

	"github.com/codepnw/gin-gorm/config"
	"github.com/codepnw/gin-gorm/handler"
	"github.com/codepnw/gin-gorm/helper"
	"github.com/codepnw/gin-gorm/model"
	"github.com/codepnw/gin-gorm/repository"
	"github.com/codepnw/gin-gorm/router"
	"github.com/codepnw/gin-gorm/service"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("server is running...")

	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.Tags{})

	repository := repository.NewTagsRepositoryImpl(db)
	service := service.NewTagsServiceImpl(repository, validate)
	handler := handler.NewTagsHandler(service)

	routes := router.NewRouter(handler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
