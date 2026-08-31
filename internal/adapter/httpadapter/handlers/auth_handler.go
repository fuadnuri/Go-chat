package handlers

import (
	"github.com/fuadnuri/Go-chat/internal/application/dto"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	userDto := dto.AuthRequestDto{}
	ctx.Bind(userDto)

}
