package service

import "discuss/model"

var PostService = &postService{}

type postService struct {}

func (s *postService) Select(categoryId, start, limit int) (posts []*model.Post, err error){
	posts, err = model.Post{
		CategoryId: categoryId,
	}.Select(start, limit)
	return
}

func (s *postService) Detail(id int) (posts model.Post, err error) {
	posts, err = model.Post{}.Find(id)
	return
}

func (s *postService) Create(categoryId int, title, content string) (bool) {
	res := model.Post{
		Title: title,
		CategoryId: categoryId,
		UserId: 1,
		Content: content,
	}.Create()

	return res
}

func (s *postService) Update(id int, title, content string) (bool) {
	res := model.Post{
		Id: id,
		Title: title,
		UserId: 1,
		Content: content,
	}.Update()

	return res
}