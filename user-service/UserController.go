package user_service

import (
	"github.com/gin-gonic/gin"
	"go-auth/middlewares"
)

func Routes(route *gin.Engine) {
	user := route.Group("/user-service")
	{
		user.Use(middlewares.JwtAuthMiddleware())
		user.GET("/", getUser)
	}
}
