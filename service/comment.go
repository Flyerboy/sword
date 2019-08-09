package service

import "discuss/model"

func GetComments(start, limit int) (comments []*model.Comment, err error){
	comments, err = model.Comment{}.Select(start, limit)
	return
}