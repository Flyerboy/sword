package main

import (
	"fmt"

	r "discuss/api/router"
	"discuss/config"
	"discuss/model"
	"github.com/spf13/viper"
)

func main() {
	config.Init("../config")
	model.Init()
	router := r.NewRouter()
	router.Run(fmt.Sprintf(":%d", viper.Get("api.port")))
}
