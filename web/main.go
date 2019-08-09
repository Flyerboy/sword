package main

import (
	"discuss/config"
	"discuss/model"
	r "discuss/web/router"
	"fmt"
	"github.com/gin-contrib/multitemplate"
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
)

func main() {
	config.Init("../config")
	model.Init()
	router := r.NewRouter()

	router.Static("/asset", "./asset")
	router.HTMLRender = loadTemplates("templates/")

	router.Run(fmt.Sprintf(":%d", viper.Get("web.port")))
}


func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "app.html")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "**/*.html")
	if err != nil {
		panic(err.Error())
	}

	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		includeNew := strings.Replace(include, "\\", "/", -1)
		name := strings.Replace(includeNew, templatesDir, "", -1)
		r.AddFromFiles(name, files...)
	}
	return r
}


