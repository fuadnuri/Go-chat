package httpadapter

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"github.com/fuadnuri/Go-chat/internal/adapter/httpadapter/handlers"
)

func Routes() http.Handler {
	slog.Info("Initializing HTTP routes")
	mux := gin.New()

	mux.GET("/", handlers.HandleRequest)
	return mux
}
