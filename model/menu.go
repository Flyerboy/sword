package model

type Menu struct {
	Id int
	Title string `gorm:"size:20" json:"title"`
	Router string `gorm:"size:50" json:"router"`
	Icon string `gorm:"size:20" json:"icon"`
	Status uint8 `gorm:"default:1" json:"status"`
	Sorts int `gorm:"default:0" json:"sorts"`
}

func (m Menu) Select(start, limit int) (menus []*Menu, err error){
	err = db.Where("status=1").Order("sorts desc").Offset(start).Limit(limit).Find(&menus).Error
	return
}
