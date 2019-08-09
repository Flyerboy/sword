package handler

import (
	"discuss/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPosts(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	size := c.DefaultQuery("size", "10")

	topics, err := service.GetPosts(0, 10)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "select error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": topics,
		"page": page,
		"size": size,
	})
}

func CreatePost(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")

	res := service.CreatePost(1, title, content)

	fmt.Println(res)

	c.JSON(http.StatusOK, gin.H{
		"msg": "OK",
		"data": gin.H{
			"title": title,
			"content": content,
		},
	})
}
