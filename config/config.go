package config

import (
	"github.com/spf13/viper"
)

func Init(path string) {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		//fmt.Println(err)
		panic(err)
	}

}
