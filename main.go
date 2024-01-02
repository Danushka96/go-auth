package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-auth/utils"
	"net/http"
)

func main() {
	binding.Validator = new(utils.DefaultValidator)
	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	r.POST("/auth/register", registerUser)
	r.POST("/auth/login", loginUser)
	r.Run()
}
