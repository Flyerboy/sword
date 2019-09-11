package service

import "discuss/model"

var UserService = &userService{}

type userService struct {}

func (s *userService) GetUserDetail(id int) (user *model.User, err error) {
	user, err = (&model.User{}).DetailById(id)
	return
}
