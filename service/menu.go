package service

import "discuss/model"

var MenuService = &menuService{}

type menuService struct {}

func (s *menuService) Select(start, limit int) (menus []*model.Menu, err error)  {
	menus, err = model.Menu{}.Select(start, limit)
	return
}