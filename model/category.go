package model

type Category struct {
	Id int
	Title string `gorm:"size:20" json:"title"`
	Status uint8 `gorm:"default:1" json:"status"`
	Sorts int `gorm:"default:0" json:"sorts"`
}

func (t Category) Select(start, limit int) (categories []*Category, err error){
	err = db.Offset(start).Limit(limit).Find(&categories).Error
	return
}