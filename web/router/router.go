package router

import (
	"discuss/web/handler"
	"github.com/gin-gonic/gin"
	"html/template"
)

func NewRouter() *gin.Engine {

	router := gin.New()

	router.Use(handler.SetCommonData)

	router.SetFuncMap(template.FuncMap{
		"html": html,
	})

	router.GET("/", handler.IndexHandler.Index)

	// post
	router.GET("/posts", handler.PostHandler.Index)
	router.GET("/posts/:id", handler.PostHandler.Detail)
	router.GET("/create/posts", handler.PostHandler.Create)
	router.POST("/create/posts", handler.PostHandler.Create)
	router.GET("/update/posts/:id", handler.PostHandler.Update)
	router.POST("/update/posts/:id", handler.PostHandler.Update)

	// comment
	router.POST("/create/comments", handler.CommentHandler.Create)

	return router
}

func html (x string) interface{} { return template.HTML(x) }


