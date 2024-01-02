package user_service

import (
	"github.com/gin-gonic/gin"
	"go-auth/constants"
	"go-auth/services"
	"go-auth/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

var userRepository *mongo.Collection

func init() {
	userRepository = services.GetDatabase().Collection(constants.UsersCollection)
}

func getUser(context *gin.Context) {
	userId, _ := services.ExtractTokenID(context)
	var currentUser types.User
	e := userRepository.FindOne(context, bson.D{{"_id", userId}}).Decode(&currentUser)
	if e != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "User not found"})
		return
	}
	context.JSON(http.StatusBadRequest, gin.H{"data": currentUser})
	return
}
