package handler

import (
	"github.com/gin-gonic/gin"
)

var PriseHandler = &priseHandler{}

type priseHandler struct{}

func (p *priseHandler) Create(c *gin.Context)  {
	//postId, _ := strconv.Atoi(c.PostForm("post_id"))
}




