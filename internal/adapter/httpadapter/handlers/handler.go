package handlers

import (
	"log/slog"
	"net/http"
	"github.com/gin-gonic/gin"
)

func HandleRequest(ctx *gin.Context) {
	slog.Info("HandleRequest called")
	ctx.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}

