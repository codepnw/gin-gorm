package repository

import "github.com/codepnw/gin-gorm/model"

type TagsRepository interface {
	Save(tags model.Tags)
	Update(tags model.Tags)
	Delete(id int)
	FindById(id int) (model.Tags, error)
	FindAll() []model.Tags
}
