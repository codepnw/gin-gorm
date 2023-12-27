package handler

import (
	"net/http"
	"strconv"

	"github.com/codepnw/gin-gorm/data/request"
	"github.com/codepnw/gin-gorm/data/response"
	"github.com/codepnw/gin-gorm/helper"
	"github.com/codepnw/gin-gorm/service"
	"github.com/gin-gonic/gin"
)

type TagsHandler struct {
	tagsService service.ITagsService
}

func NewTagsHandler(service service.ITagsService) *TagsHandler {
	return &TagsHandler{
		tagsService: service,
	}
}

func (h *TagsHandler) Create(ctx *gin.Context) {
	createTagsReq := request.CreateTagsRequest{}
	err := ctx.ShouldBindJSON(&createTagsReq)
	helper.ErrorPanic(err)

	h.tagsService.Create(createTagsReq)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (h *TagsHandler) Update(ctx *gin.Context) {
	updateTagsReq := request.UpdateTagsRequest{}
	err := ctx.ShouldBindJSON(updateTagsReq)
	helper.ErrorPanic(err)

	tagId := ctx.Param("id")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	updateTagsReq.Id = id

	h.tagsService.Update(updateTagsReq)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (h *TagsHandler) Delete(ctx *gin.Context) {
	tagId := ctx.Param("id")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)

	h.tagsService.Delete(id)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (h *TagsHandler) FindById(ctx *gin.Context) {
	tagId := ctx.Param("id")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)

	tagResponse := h.tagsService.FindById(id)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   tagResponse,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (h *TagsHandler) FindAll(ctx *gin.Context) {
	tagResponse := h.tagsService.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   tagResponse,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
