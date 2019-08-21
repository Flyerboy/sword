package model

import "time"

type Comment struct {
	Id          int `json:"id"`
	ParentId 	int
	UserId      int `json:"user_id"`
	PostId      int `json:"post_id"`
	Post        Post
	Content     string    `json:"content" gorm:"type:text"`
	CreatedAt   time.Time `json:"created_at"`
	PriseNumber int
	IsTop       uint8 `gorm:"default:0"`
	Status      uint8 `gorm:"default:1" json:"status"`
}

func (c Comment) Select(start, limit int) (comments []*Comment, err error) {
	err = db.Where(&c).Where("status=1").Order("id desc").Offset(start).Limit(limit).Find(&comments).Error
	return
}

func (c *Comment) Create() (res bool, err error) {
	err = db.Create(c).Error
	if err == nil {
		res = true
	}
	return
}

func (c Comment) Count() (total int) {
	err := db.Model(&Comment{}).Where(&c).Where("status=1").Count(&total).Error
	if err != nil {
		total = 0
	}
	return
}
