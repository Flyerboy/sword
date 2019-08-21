package service

import "discuss/model"

var User = &UserService{}

type UserService struct {}

func (s *UserService) GetUserDetail(id int) (user *model.User, err error) {
	user, err = (&model.User{}).DetailById(id)
	return
}
