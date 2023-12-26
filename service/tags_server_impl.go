package service

import (
	"github.com/codepnw/gin-gorm/data/request"
	"github.com/codepnw/gin-gorm/data/response"
	"github.com/codepnw/gin-gorm/helper"
	"github.com/codepnw/gin-gorm/model"
	"github.com/codepnw/gin-gorm/repository"
	"github.com/go-playground/validator/v10"
)

type TagsServiceImpl struct {
	TagsRepository repository.ITagsRepository
	validate       *validator.Validate
}

func NewTagsServiceImpl(repo repository.ITagsRepository, validate *validator.Validate) ITagsService {
	return &TagsServiceImpl{
		TagsRepository: repo,
		validate:       validate,
	}
}

func (t *TagsServiceImpl) Create(tags request.CreateTagsRequest) {
	err := t.validate.Struct(tags)
	helper.ErrorPanic(err)
	tagModel := model.Tags{
		Name: tags.Name,
	}
	t.TagsRepository.Save(tagModel)
}

func (t *TagsServiceImpl) Update(tags request.UpdateTagsRequest) {
	data, err := t.TagsRepository.FindById(tags.Id)
	helper.ErrorPanic(err)

	data.Name = tags.Name
	t.TagsRepository.Update(data)
}

func (t *TagsServiceImpl) Delete(id int) {
	t.TagsRepository.Delete(id)
}

func (t *TagsServiceImpl) FindById(id int) response.TagsResponse {
	data, err := t.TagsRepository.FindById(id)
	helper.ErrorPanic(err)

	tagResponse := response.TagsResponse{
		Id:   data.ID,
		Name: data.Name,
	}
	return tagResponse
}

func (t *TagsServiceImpl) FindAll() []response.TagsResponse {
	result := t.TagsRepository.FindAll()

	var tags []response.TagsResponse
	for _, v := range result {
		tag := response.TagsResponse{
			Id:   v.ID,
			Name: v.Name,
		}
		tags = append(tags, tag)
	}

	return tags
}
