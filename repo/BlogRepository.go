package repo

import (
	"blogs/model"

	"gorm.io/gorm"
)

type BlogRepository struct {
	DatabaseConnection *gorm.DB
}

func (repository *BlogRepository) CreateBlog(blog *model.Blog) error {
	dbResult := repository.DatabaseConnection.Create(blog)
	if dbResult.Error != nil{
		return dbResult.Error
	}
	return nil
}