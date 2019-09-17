package service

import "discuss/model"

var BookService = &bookService{}

type bookService struct {}

func (this *bookService) Select(start, limit int) (books []*model.Book, err error) {
	books, err = model.Book{}.Select(start, limit)
	return
}

func (this *bookService) Detail(id int) (book model.Book, err error) {
	book, err = model.Book{}.Find(id)
	return
}

func (this *bookService) Chapters(bookId int) (chapters []*model.BookChapter, err error) {
	chapters, err = model.BookChapter{}.Chapters(bookId)
	return
}

func (this *bookService) Chapter(bookId, chapterId int) (chapter model.BookChapter, err error) {
	if chapterId > 0 {
		chapter, err = model.BookChapter{}.Find(chapterId)
	} else {
		chapter, err = model.BookChapter{}.First(bookId)
	}
	return
}
