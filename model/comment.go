package model

import "time"

type Comment struct {
	Id          int `json:"id"`
	UserId      int `json:"user_id"`
	PostId      int `json:"post_id"`
	Post        Post
	Content     string    `json:"content" gorm:"type:text"`
	CreatedAt   time.Time `json:"created_at"`
	PriseNumber int
	IsTop       uint8 `gorm:"default:0"`
	Status      uint8 `gorm:"default:1" json:"status"`
}

func (t Comment) Select(start, limit int) (comments []*Comment, err error) {
	err = db.Where("status=?", 1).Offset(start).Limit(limit).Find(&comments).Error
	return
}
