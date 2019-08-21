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

	router.GET("/", handler.Index)

	// post
	postHandler := handler.PostHandler{}
	router.GET("/posts", postHandler.GetPosts)
	router.GET("/posts/:id", postHandler.GetPostDetail)
	router.GET("/create/posts", postHandler.CreatePost)
	router.POST("/create/posts", postHandler.CreatePost)
	router.GET("/update/posts/:id", postHandler.UpdatePost)
	router.POST("/update/posts/:id", postHandler.UpdatePost)

	// comment
	commentHandler := handler.CommentHandler{}
	router.POST("/create/comments", commentHandler.CreateComment)

	return router
}

func html (x string) interface{} { return template.HTML(x) }


