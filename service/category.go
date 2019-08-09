package service

import "discuss/model"

func GetCategories(start, limit int) (categories []*model.Category, err error){
	categories, err = model.Category{}.Select(start, limit)
	return
}
