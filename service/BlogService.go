package service

import (
	"blogs/model"
	"blogs/repo"
)

type BlogService struct {
	BlogRepo *repo.BlogRepository
}

func (service *BlogService) CreateBlog(blog *model.Blog) error {
	result := service.BlogRepo.CreateBlog(blog)
	if result != nil {
		return result
	}
	return nil
}