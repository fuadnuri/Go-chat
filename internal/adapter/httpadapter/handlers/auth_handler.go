package handlers

import (
	"github.com/fuadnuri/Go-chat/internal/application/dto"
	"github.com/gin-gonic/gin"
	"log/slog"
)

func Login(ctx *gin.Context) {
	//add logging
	slog.Info("Login endpoint called")
	
	
	userDto := dto.AuthRequestDto{}
	err, ok := ctx.Bind(&userDto)
	if !ok {
		slog.Error("Failed to bind request", "error", err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	//check if user exists
	user, err := GetUserByEmail(userDto.Email)
	if err != nil {
		slog.Error("User not found", "email", userDto.Email)
		ctx.JSON(404, gin.H{"error": "User not found"})
		return
	}
	//if user exists, check if password is correct
	if user.Password != userDto.Password {
		slog.Error("Invalid password", "email", userDto.Email)
		ctx.JSON(401, gin.H{"error": "Invalid password"})
		return
	}
	//if password is correct, return token
	slog.Info("Login successful", "email", userDto.Email)
	ctx.JSON(200, gin.H{"token": "token"})
	return
	//if user does not exist, return error
	
}


