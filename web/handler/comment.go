package handler

import (
	"discuss/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var CommentHandler = &commentHandler{}

type commentHandler struct {}

func (this *commentHandler) Index(c *gin.Context) {

}

func (this *commentHandler) Create(c *gin.Context) {
	postId, _ := strconv.Atoi(c.PostForm("post_id"))
	content := c.PostForm("content")

	if postId > 0 && content != "" {
		if res, _ := service.CommentService.Create(postId, 1, content, 0); res {
			c.Redirect(http.StatusMovedPermanently, "/posts/" + c.PostForm("post_id"))
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "create error",
		})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"msg": "param error",
	})
}
