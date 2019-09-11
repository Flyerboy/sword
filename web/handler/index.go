package handler

import (
	"discuss/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var IndexHandler = &indexHandler{}

type indexHandler struct {}

func (this *indexHandler) Index(c *gin.Context)  {
	start, limit := getPage(c)
	posts, _ := service.PostService.Select(0, start, limit)

	data := getCommonData(c)
	data["posts"] = posts
	data["title"] = "首页"

	c.HTML(http.StatusOK, "index/index.html", data)
}
