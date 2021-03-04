package app

import (
	"github.com/gin-gonic/gin"
	"github.com/JonathanUcles/HelloWorldGoApp/src/http"
)

var (
	router = gin.Default()
)

func StartApplication() {
	router.GET("/hello", http.GetHelloWorld)
	router.Run(":8080")
}
