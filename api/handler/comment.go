package handler

import (
	"discuss/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetComments(c *gin.Context) {
	comments, err := service.GetComments(0, 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "get comments error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": comments,
		"total": 20,
	})
}

func CreateComment(c *gin.Context) {

}
