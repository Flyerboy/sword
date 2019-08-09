package handler

import (
	"discuss/service"
	"discuss/web/util"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"strconv"
)

func GetPosts(c *gin.Context) {
	//page := c.DefaultQuery("page", "1")
	//size := c.DefaultQuery("size", "10")

	posts, _ := service.GetPosts(0, 10)
	categories, _ := service.GetCategories(0, 10)
	c.HTML(http.StatusOK, "post/index.html", gin.H{
		"posts": posts,
		"title": "文章列表",
		"categories": categories,
	})
}

func GetPostsDetail(c *gin.Context) {
	id := c.Param("id")
	postId, _ := strconv.Atoi(id)
	categories, _ := service.GetCategories(0, 10)
	postService := service.PostService{}
	post, _ := postService.GetPostDetail(postId)

	content := util.MarkDown(post.Content)
	c.HTML(http.StatusOK, "post/detail.html", gin.H{
		"post": post,
		"title": "文章详情",
		"categories": categories,
		"content": template.HTML(content),
	})
}

