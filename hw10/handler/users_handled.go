package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/zhyzhkoslava/ITEA-GO/hw10/structure"
)

type UsersHandler struct {
	logger *log.Logger
	Users  []structure.User
}

func NewUsersHandler(logger *log.Logger, users []structure.User) *UsersHandler {
	return &UsersHandler{
		logger: logger,
		Users:  users,
	}
}

func (h *UsersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(h.Users); err != nil {
		h.logger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
