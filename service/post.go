package service

import "discuss/model"

func CreatePost(userId int, title string, content string) bool {

	res := model.Post{
		Title: title,
		Content: content,
	}.Create()

	return res
}

func GetPosts(start, limit int) (posts []*model.Post, err error){
	posts, err = model.Post{}.Select(start, limit)
	return
}

type PostService struct {

}

func (s *PostService) GetPostDetail(id int) (posts model.Post, err error) {
	posts, err = model.Post{}.Find(id)
	return
}