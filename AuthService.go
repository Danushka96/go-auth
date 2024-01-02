package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-auth/constants"
	"go-auth/types"
	"go-auth/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

var userRepository *mongo.Collection

func init() {
	userRepository = getDatabase().Collection(constants.UsersCollection)
}

func registerUser(context *gin.Context) {
	var req types.RegisterUserRequest
	if err := context.BindJSON(&req); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var currentUser types.User
	findError := userRepository.FindOne(context, bson.D{{"email", req.Email}}).Decode(&currentUser)
	if findError == nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "User already exists"})
		return
	}

	var id = uuid.Must(uuid.NewRandom()).String()
	password, _ := utils.HashPassword(req.Password)
	var user = types.User{Id: id, FirstName: req.FirstName, LastName: req.LastName, Email: req.Email, Password: password}

	_, err := userRepository.InsertOne(context, user)
	if err != nil {
		panic(err)
	}
	context.JSON(http.StatusOK, gin.H{"message": "Inserted with id"})
}

func loginUser(context *gin.Context) {
	var req types.LoginUserRequest
	if err := context.BindJSON(&req); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var currentUser types.User
	findError := userRepository.FindOne(context, bson.D{{"email", req.Email}}).Decode(&currentUser)
	if findError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "User not found"})
		return
	}

	isValid := utils.CheckPasswordHash(req.Password, currentUser.Password)

	if !isValid {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Password does not match"})
		return
	}

	context.JSON(http.StatusBadRequest, gin.H{"message": "User logged in"})
}
