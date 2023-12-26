package repository

import (
	"errors"

	"github.com/codepnw/gin-gorm/data/request"
	"github.com/codepnw/gin-gorm/helper"
	"github.com/codepnw/gin-gorm/model"
	"gorm.io/gorm"
)

type TagsRepositoryImpl struct {
	Db *gorm.DB
}

func NewTagsRepositoryImpl(Db *gorm.DB) ITagsRepository {
	return &TagsRepositoryImpl{Db: Db}
}

func (t *TagsRepositoryImpl) Save(tags model.Tags) {
	result := t.Db.Create(&tags)
	helper.ErrorPanic(result.Error)
}

func (t *TagsRepositoryImpl) Update(tags model.Tags) {
	var update = request.UpdateTagsRequest{
		Id:   tags.ID,
		Name: tags.Name,
	}
	result := t.Db.Model(&tags).Updates(update)
	helper.ErrorPanic(result.Error)
}

func (t *TagsRepositoryImpl) Delete(id int) {
	var tags model.Tags
	result := t.Db.Where("id = ?", id).Delete(&tags)
	helper.ErrorPanic(result.Error)
}

func (t *TagsRepositoryImpl) FindById(id int) (model.Tags, error) {
	var tag model.Tags
	result := t.Db.Find(&tag, id)
	if result != nil {
		return tag, nil
	} else {
		return tag, errors.New("tag is not found")
	}
}

func (t *TagsRepositoryImpl) FindAll() []model.Tags {
	var tags []model.Tags
	result := t.Db.Find(&tags)
	helper.ErrorPanic(result.Error)
	return tags
}
