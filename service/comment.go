package service

import "discuss/model"

var CommentService = &commentService{}

type commentService struct {}

func (s *commentService) Select(postId, start, limit int) (comments []*model.Comment, err error){
	comments, err = model.Comment{
		PostId: postId,
	}.Select(start, limit)
	return
}

func (s *commentService) Total(postId int) (total int)  {
	total = model.Comment{
		PostId: postId,
	}.Count()
	return
}

func (s *commentService) Create(postId, userId int, content string, parentId int) (res bool, err error) {
	commentModel := model.Comment{
		PostId: postId,
		UserId: userId,
		Content: content,
		ParentId: parentId,
	}
	res, err = commentModel.Create()
	return
}