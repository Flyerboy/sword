package router

import (
	"discuss/api/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {

	router := gin.New()

	// post
	router.GET("/posts", handler.PostHandler.Index)
	router.POST("/posts/create", handler.PostHandler.Create)

	// comment
	router.GET("/comments", handler.CommentHandler.Index)
	router.POST("/comments/create", handler.CommentHandler.Create)

	return router
}