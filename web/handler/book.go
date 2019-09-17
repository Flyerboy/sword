package handler

import (
	"discuss/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var BookHandler = &bookHandler{}

type bookHandler struct {
}

func (this *bookHandler) Index(c *gin.Context) {
	data := getCommonData(c)
	data["Title"] = "图书"
	data["Books"], _ = service.BookService.Select(0, 10)
	c.HTML(http.StatusOK, "book/index.html", data)
}

func (this *bookHandler) Detail(c *gin.Context) {
	id := c.Param("id")
	bookId, _ := strconv.Atoi(id)

	book, err := service.BookService.Detail(bookId)
	if err != nil {

	} else {
		data := getCommonData(c)
		data["Title"] = "图书 - " + book.Title
		data["Book"] = book
		data["Chapters"], _ = service.BookService.Chapters(bookId)

		chapterId := c.DefaultQuery("chapter_id", "0")
		newChapterId, _ := strconv.Atoi(chapterId)

		data["Chapter"], _ = service.BookService.Chapter(bookId, newChapterId)
		c.HTML(http.StatusOK, "book/detail.html", data)
	}
}
