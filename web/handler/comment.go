package handler

import (
	"discuss/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CommentHandler struct {}

func (h CommentHandler) CreateComment(c *gin.Context) {
	postId, _ := strconv.Atoi(c.PostForm("post_id"))
	content := c.PostForm("content")

	fmt.Println(postId)

	commentService := service.CommentService{}
	if res, _ := commentService.CreateComment(postId, 1, content, 0); res {
		c.Redirect(http.StatusMovedPermanently, "/posts/" + c.PostForm("post_id"))
		return
	}
	fmt.Println("error")
}
