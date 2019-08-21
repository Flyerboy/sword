package service

import "discuss/model"

type MenuService struct {}

func (s MenuService) GetMenus(start, limit int) (menus []*model.Menu, err error)  {
	menus, err = model.Menu{}.Select(start, limit)
	return
}