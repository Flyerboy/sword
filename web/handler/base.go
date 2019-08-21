package handler

import (
	"discuss/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"strconv"
)

type CommonData map[string]interface{}

type FlushData struct {
	Success string
	Error string
}

func SetCommonData(c *gin.Context) {
	commonData := CommonData{}
	commonData["Menus"], _ = service.MenuService{}.GetMenus(0, 10)
	commonData["WebName"] = viper.Get("web.name")
	commonData["GitHub"] = "https://github.com/login/oauth/authorize?client_id=7e015d8ce32370079895&redirect_uri=http://localhost:8080/oauth/redirect"
	commonData["IsLogin"] = false
	commonData["Flush"] = FlushData{}
	c.Set("CommonData", &commonData)
	c.Next()
}

func getCommonData(c *gin.Context) CommonData {
	commonDataVal, _ := c.Get("CommonData")
	commonData := commonDataVal.(*CommonData)
	return *commonData
}

func getPage(c *gin.Context) (start, limit int)  {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ = strconv.Atoi(c.DefaultQuery("limit", "20"))

	start = (page - 1) * limit
	return
}
