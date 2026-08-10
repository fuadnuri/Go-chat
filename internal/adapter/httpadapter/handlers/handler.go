package handlers

import (
	"net/http"

	"www.github.com/fuadnuri/Go-chat/internal/adapter/httpadapter/response"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	response.WriteJSON(w, http.StatusOK, "Hello, World!")
}
