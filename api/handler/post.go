package handler

import (
	"discuss/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var PostHandler = &postHandler{}

type postHandler struct {

}


func (p *postHandler) Index(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	size := c.DefaultQuery("size", "10")

	topics, err := service.PostService.Select(0, 0, 10)

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

func (p *postHandler) Create(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")

	res := service.PostService.Create(1, title, content)

	fmt.Println(res)

	c.JSON(http.StatusOK, gin.H{
		"msg": "OK",
		"data": gin.H{
			"title": title,
			"content": content,
		},
	})
}
