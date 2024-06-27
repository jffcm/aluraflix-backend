package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jffcm/aluraflix-backend/schemas"
	"github.com/jffcm/aluraflix-backend/utils"
)

func CreateUserHandler(ctx *gin.Context) {
	var request CreateUserRequest

	if err := ctx.BindJSON(&request); err != nil {
		log.Printf("Error decoding JSON request: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, "Failed to process request data")
		return
	}

	if err := request.Validate(); err != nil {
		log.Printf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user := schemas.User{
		Username: request.Username,
		Password: request.Password,
	}

	if err := db.Create(&user).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "Error creating user in database")
		return
	}

	sendSuccess(ctx, http.StatusCreated, "create-user", user)
}

func LoginUserHandler(ctx *gin.Context) {
	var request LoginUserRequest

	if err := ctx.BindJSON(&request); err != nil {
		log.Printf("Error decoding JSON request: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, "Failed to process request data")
		return
	}

	if err := request.Validate(); err != nil {
		log.Printf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var user schemas.User
	if err := db.Where("username = ?", request.Username).First(&user).Error; err != nil {
		log.Printf("User not found: %v", err)
		sendError(ctx, http.StatusUnauthorized, "Invalid username or password")
		return
	}

	if err := utils.ValidatePassword(user.Password, request.Password); err != nil {
		log.Printf("Password validation failed: %v", err)
		sendError(ctx, http.StatusUnauthorized, "Unauthorized")
		return
	}

	token, err := utils.CreateToken(user.Username)
	if err != nil {
		log.Printf("Error generating token: %v", err)
		sendError(ctx, http.StatusInternalServerError, "Could not generate token")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
