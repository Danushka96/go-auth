package auth_service

import (
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	auth := route.Group("/auth-service")
	{
		auth.POST("/login", LoginUser)
		auth.POST("/register", RegisterUser)
	}
}
