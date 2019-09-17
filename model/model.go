package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var db *gorm.DB

func Init()  {

	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true",
						viper.Get("database.user"),
						viper.Get("database.password"),
						viper.Get("database.host"),
						viper.GetInt("database.port"),
						viper.Get("database.db"))
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic("db link error")
	}

	db.AutoMigrate(&Post{}, &Comment{}, &Category{}, &Menu{}, &User{}, &Course{}, &CourseChapter{},
		&Book{}, &BookChapter{})

}
