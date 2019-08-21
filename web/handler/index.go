package handler

import (
	"discuss/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context)  {
	start, limit := getPage(c)
	postService := service.PostService{}
	posts, _ := postService.GetPosts(0, start, limit)

	data := getCommonData(c)
	data["posts"] = posts
	data["title"] = "首页"

	c.HTML(http.StatusOK, "index/index.html", data)
}
