package handler

import (
	"discuss/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var CommentHandler = &commentHandler{}

type commentHandler struct {

}

func (this *commentHandler) Index(c *gin.Context) {
	comments, err := service.CommentService.Select(1, 0, 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "get comments error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": comments,
		"total": 20,
	})
}

func (this *commentHandler) Create(c *gin.Context) {

}
