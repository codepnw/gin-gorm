package main

import (
	"net/http"

	"github.com/codepnw/gin-gorm/config"
	"github.com/codepnw/gin-gorm/handler"
	"github.com/codepnw/gin-gorm/helper"
	"github.com/codepnw/gin-gorm/model"
	"github.com/codepnw/gin-gorm/repository"
	"github.com/codepnw/gin-gorm/service"
	"github.com/gin-gonic/gin"
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
	_ = handler

	routes := gin.Default()

	routes.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
