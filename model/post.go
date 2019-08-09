package model

import (
	"fmt"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Post struct {
	Id          int
	UserId      int
	CategoryId  int
	ReadNumber  int
	Category    Category
	Title       string `json:"title"`
	EnTitle		string 	`json:"en_title"`
	Content     string `json:"content" gorm:"type:text"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	IsTop       uint8 `gorm:"default:0"`
	IsRecommend uint8 `gorm:"default:0"`
	Status      uint8 `gorm:"default:1"`
}

func (t Post) Create() bool {
	err := db.Create(t).Error
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func (t Post) Select(start, limit int) (posts []*Post, err error) {
	err = db.Select("id,title,content").Offset(start).Limit(limit).Find(&posts).Error
	return
}

func (t Post) Find(id int) (post Post, err error) {
	err = db.Select("id,title,content").First(&post, id).Error
	return
}

