package handlers

import (
	"encoding/json"
	"net/http"

	"platzi.com/go/rest-ws/server"
)

type homeResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func HomeHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "aplicattion/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(homeResponse{
			Message: "Welcome to the home page",
			Status:  true,
		})
	}
}
