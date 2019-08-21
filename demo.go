package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)


var tokenSecret = "AD#$@12232fd"

func main()  {
	router := gin.New()

	router.POST("/login", login)

	router.GET("/user", authToken(), user)

	router.Run(":9999")
}

func login(c *gin.Context)  {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "flyer" && password == "123456" {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, &JwtData{
			"flyer",
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 20).Unix(),
			},
		})

		tokenStr, err := token.SignedString([]byte(tokenSecret))

		fmt.Println(tokenStr)

		c.JSON(http.StatusOK, gin.H{
			"token": tokenStr,
			"err": err,
		})
	} else {
		c.JSON(http.StatusUnauthorized, nil)
	}
}

func user(c *gin.Context) {

	auth := c.GetHeader("authorization")
	fmt.Println(auth)

	c.JSON(http.StatusOK, gin.H{
		"user": "flyer",
	})
}

type JwtData struct {
	User string `json:"user"`
	jwt.StandardClaims
}

func authToken() gin.HandlerFunc {
	return func (c *gin.Context) {
		tokenString := c.Query("token")
		jwtObj := JwtData{}
		token, err := jwt.ParseWithClaims(tokenString, &jwtObj, func(token *jwt.Token) (interface{}, error) {
			return tokenSecret, nil
		})

		if err != nil && token != nil && token.Valid {
			fmt.Println(token.Claims)
			c.Next()
			return
		}

		c.Abort()
		c.JSON(401, gin.H{
			"code": -1,
			"msg": err,
		})
	}
}