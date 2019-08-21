package handler

import (
	"discuss/model"
	"discuss/service"
	"discuss/web/util"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"strconv"
)

type PostHandler struct {}

func (h PostHandler) GetPosts(c *gin.Context) {
	start, limit := getPage(c)

	categoryId, _ := strconv.Atoi(c.DefaultQuery("category_id", "0"))
	sorts, _ := strconv.Atoi(c.DefaultQuery("sorts", "1"))

	postService := service.PostService{}
	posts, _ := postService.GetPosts(categoryId, start, limit)

	data := getCommonData(c)
	data["Posts"] = posts
	data["CategoryId"] = categoryId
	data["Sorts"] = sorts
	data["Title"] = "文章列表"
	data["Categories"], _ = service.GetCategories(0, 10)
	c.HTML(http.StatusOK, "post/index.html", data)
}

func (h PostHandler) GetPostDetail(c *gin.Context) {
	id := c.Param("id")
	postId, _ := strconv.Atoi(id)
	postService := service.PostService{}
	post, _ := postService.GetPostDetail(postId)

	content := util.MarkDown(post.Content)

	data := getCommonData(c)
	data["Post"] = post
	data["Title"] = "文章详情"
	data["Content"] = template.HTML(content)

	commentService := service.CommentService{}
	data["Comments"], _ = commentService.GetComments(postId, 0, 10)
	data["CommentTotal"] = commentService.GetTotal(postId)

	c.HTML(http.StatusOK, "post/detail.html", data)
}

func (h PostHandler) CreatePost(c *gin.Context) {
	var success, error string
	if c.Request.Method == "POST" {
		title := c.PostForm("title")
		categoryId, _ := strconv.Atoi(c.PostForm("category_id"))
		content := c.PostForm("content")

		postService := service.PostService{}
		res := postService.CreatePost(categoryId, title, content)
		if res {
			success = "发布成功"
		} else {
			error = "发布失败"
		}
	}

	data := getCommonData(c)
	data["Flush"] = FlushData{
		Success: success,
		Error: error,
	}
	data["Title"] = "创建文章"
	data["Post"] = model.Post{}
	data["Categories"], _ = service.GetCategories(0, 10)
	c.HTML(http.StatusOK, "post/create.html", data)
}

func (h PostHandler) UpdatePost(c *gin.Context) {
	var success, error string
	id, _ := strconv.Atoi(c.Param("id"))

	if c.Request.Method == "POST" {
		title := c.PostForm("title")
		content := c.PostForm("content")

		postService := service.PostService{}
		res := postService.UpdatePost(id, title, content)
		if res {
			success = "编辑成功"
		} else {
			error = "编辑失败"
		}
	}

	data := getCommonData(c)
	data["Flush"] = FlushData{
		Success: success,
		Error: error,
	}
	data["Title"] = "编辑文章"
	data["Post"], _ = model.Post{}.Find(id)
	data["Categories"], _ = service.GetCategories(0, 10)
	c.HTML(http.StatusOK, "post/update.html", data)
}
