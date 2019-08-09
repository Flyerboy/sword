package router

import (
	"discuss/web/handler"
	"github.com/gin-gonic/gin"
	"html/template"
)

func NewRouter() *gin.Engine {

	router := gin.New()

	router.SetFuncMap(template.FuncMap{
		"html": html,
	})

	router.GET("/", handler.Index)

	// post
	router.GET("/posts", handler.GetPosts)
	router.GET("/posts/:id", handler.GetPostsDetail)

	return router
}

func html (x string) interface{} { return template.HTML(x) }