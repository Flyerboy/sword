package handler

import (
	"discuss/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context)  {
	posts, _ := service.GetPosts(0, 10)
	categories, _ := service.GetCategories(0, 10)
	c.HTML(http.StatusOK, "index/index.html", gin.H{
		"posts": posts,
		"title": "title",
		"categories": categories,
	})
}
