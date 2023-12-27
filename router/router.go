package router

import (
	"net/http"

	"github.com/codepnw/gin-gorm/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *handler.TagsHandler) *gin.Engine {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})

	baseRouter := router.Group("/api")
	tagsRouter := baseRouter.Group("/tags")

	tagsRouter.GET("/", handler.FindAll)
	tagsRouter.GET("/:id", handler.FindById)
	tagsRouter.POST("/", handler.Create)
	tagsRouter.PATCH("/:id", handler.Update)
	tagsRouter.DELETE("/:id", handler.Delete)

	return router
}
