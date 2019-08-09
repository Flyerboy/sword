package router

import (
	"discuss/api/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {

	router := gin.New()

	// post
	router.GET("/posts", handler.GetPosts)
	router.POST("/posts/create", handler.CreatePost)

	// comment
	router.GET("/comments", handler.GetComments)
	router.POST("/comments/create", handler.CreateComment)

	return router
}