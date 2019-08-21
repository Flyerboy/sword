package service

import "discuss/model"

type PostService struct {}

func (s *PostService) GetPosts(categoryId, start, limit int) (posts []*model.Post, err error){
	posts, err = model.Post{
		CategoryId: categoryId,
	}.Select(start, limit)
	return
}

func (s *PostService) GetPostDetail(id int) (posts model.Post, err error) {
	posts, err = model.Post{}.Find(id)
	return
}

func (s *PostService) CreatePost(categoryId int, title, content string) (bool) {
	res := model.Post{
		Title: title,
		CategoryId: categoryId,
		UserId: 1,
		Content: content,
	}.Create()

	return res
}

func (s *PostService) UpdatePost(id int, title, content string) (bool) {
	res := model.Post{
		Id: id,
		Title: title,
		UserId: 1,
		Content: content,
	}.Update()

	return res
}