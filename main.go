package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-auth/auth-service"
	"go-auth/user-service"
	"go-auth/utils"
)

func main() {
	binding.Validator = new(utils.DefaultValidator)
	router := gin.Default()
	auth_service.Routes(router)
	user_service.Routes(router)
	router.Run()
}
