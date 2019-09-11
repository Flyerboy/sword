package service

import "discuss/model"

var CategoryService = &categoryService{}

type categoryService struct {}

func GetCategories(start, limit int) (categories []*model.Category, err error){
	categories, err = model.Category{}.Select(start, limit)
	return
}
