package httpadapter

import (
	"net/http"

	"github.com/fuadnuri/Go-chat/internal/adapter/httpadapter/handlers"
)

func Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HandleRequest)
	return mux
}
