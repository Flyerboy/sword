package model

import "time"

type Book struct {
	Id          uint
	UserId      uint
	User 		User
	Title       string `json:"title"`
	EnTitle		string 	`json:"en_title"`
	Image string `json:"image"`
	Introduction     string `json:"introduction" gorm:"type:text"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	IsTop       uint8 `gorm:"default:0"`
	Status      uint8 `gorm:"default:1"`
}


type BookChapter struct {
	Id          uint
	BookId      uint
	Book 		Book
	Title       string `json:"title"`
	EnTitle		string 	`json:"en_title"`
	Content string	`json:"content" gorm:"type:text"`
	IsCharge       uint8 `gorm:"default:0"`
	Status      uint8 `gorm:"default:1"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}


func (c Book) Select(start, limit int) (books []*Book, err error) {
	err = db.Where(&c).
		Order("is_top desc").
		Order("id desc").
		Offset(start).
		Limit(limit).
		Find(&books).
		Error
	return
}

func (c Book) Find(id int) (book Book, err error) {
	err = db.First(&book, id).Error
	return
}

func (c BookChapter) Chapters(bookId int) (chapters []*BookChapter, err error) {
	err = db.Where("book_id=? and status=1", bookId).
		Order("id asc").
		Find(&chapters).
		Error
	return
}

func (c BookChapter) Find(id int) (chapter BookChapter, err error) {
	err = db.First(&chapter, id).Error
	return
}

func (c BookChapter) First(id int) (chapter BookChapter, err error) {
	err = db.Where("book_id=?", id).
		Order("id asc").
		Limit(1).
		Find(&chapter).
		Error
	return
}

