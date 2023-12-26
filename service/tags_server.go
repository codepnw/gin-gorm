package service

import (
	"github.com/codepnw/gin-gorm/data/request"
	"github.com/codepnw/gin-gorm/data/response"
)

type ITagsService interface {
	Create(tags request.CreateTagsRequest)
	Update(tags request.UpdateTagsRequest)
	Delete(id int)
	FindById(id int) response.TagsResponse
	FindAll() []response.TagsResponse
}
