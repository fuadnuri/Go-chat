package handlers

import (
	"github.com/fuadnuri/Go-chat/internal/application/dto"
	"github.com/gin-gonic/gin"
	"github.com/fuadnuri/Go-chat/internal/utils"
	"github.com/fuadnuri/Go-chat/internal/application/usecases"
)

func AuthHandler(ctx *gin.Context) {
	userDto := dto.AuthRequestDto{}
	userDto.Username = ctx.Query("username")
	userDto.Password = ctx.Query("password")
	user_usecase := usecases.NewUserUsecase()
	
	if passwordHash,err:=utils.HashPassword(userDto.Password); err!=nil{
		fmt.Println("password hashing error")
	}
	userDto.Password=passwordHash
	user.Login(userDto)
	
}
