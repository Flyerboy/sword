package model

import (
	"fmt"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Post struct {
	Id          int
	UserId      int
	User 		User
	CategoryId  int
	ReadNumber  int	`gorm:"default:0"`
	PriseNumber int `gorm:"default:0"`
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
	err := db.Create(&t).Error
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func (p Post) Select(start, limit int) (posts []*Post, err error) {
	err = db.Select("id,user_id,title,content,is_top,is_recommend,read_number,prise_number").
			Where(&p).
			Order("is_top desc").
			Order("id desc").
			Offset(start).
			Limit(limit).
			Find(&posts).
			Error
	return
}

func (p Post) Find(id int) (post Post, err error) {
	err = db.First(&post, id).Error
	return
}

func (p Post) Update() bool {
	err := db.Model(&p).Updates(&Post{
		Title: p.Title,
		Content: p.Content,
	}).Error
	if err != nil {
		return false
	}
	return true
}

