package model

import (
	"fmt"
	"time"
)

type Course struct {
	Id          int
	UserId      int
	User 		User
	Title       string `json:"title"`
	EnTitle		string 	`json:"en_title"`
	Image string `json:"image"`
	Times int	`json:"times"`
	Introduction     string `json:"introduction" gorm:"type:text"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	IsTop       uint8 `gorm:"default:0"`
	Status      uint8 `gorm:"default:1"`
}

func (c Course) Select(start, limit int) (courses []*Course, err error) {
	err = db.Where(&c).
		Order("is_top desc").
		Order("id desc").
		Offset(start).
		Limit(limit).
		Find(&courses).
		Error
	return
}

func (c Course) Find(id int) (course Course, err error) {
	err = db.First(&course, id).Error
	return
}

func (c Course) Create() (result bool, err error) {
	err = db.Create(&c).Error
	result = true
	if err != nil {
		fmt.Println(err)
		result = false
	}
	return
}


type CourseChapter struct {
	Id          int
	CourseId      int
	Course 		Course
	Title       string `json:"title"`
	EnTitle		string 	`json:"en_title"`
	Url string	`json:"url"`
	Times int	`json:"times"`
	IsCharge       uint8 `gorm:"default:0"`
	Status      uint8 `gorm:"default:1"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (c CourseChapter) Chapter(courseId int) (chapters []*CourseChapter, err error) {
	err = db.Where("course_id=? and status=1", courseId).
		Order("id asc").
		Find(&chapters).
		Error
	return
}

func (c CourseChapter) Find(id int) (chapter CourseChapter, err error) {
	err = db.First(&chapter, id).Error
	return
}

func (c CourseChapter) First(id int) (chapter CourseChapter, err error) {
	err = db.Where("course_id=?", id).
		Order("id asc").
		Limit(1).
		Find(&chapter).
		Error
	return
}