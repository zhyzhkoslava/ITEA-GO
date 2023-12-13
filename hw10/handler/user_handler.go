package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/zhyzhkoslava/ITEA-GO/hw10/structure"
)

type UserHandler struct {
	logger    *log.Logger
	Users     []structure.User
	UsersFile string
}

func NewUserHandler(logger *log.Logger, users []structure.User, usersFile string) *UserHandler {
	return &UserHandler{
		logger:    logger,
		Users:     users,
		UsersFile: usersFile,
	}
}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		h.handleGet(w, r)
	case http.MethodPost:
		h.handlePost(w, r)
	case http.MethodPatch:
		h.handlePatch(w, r)
	case http.MethodDelete:
		h.handleDelete(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *UserHandler) handleGet(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/user/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var foundUser *structure.User
	for _, user := range h.Users {
		if user.ID == id {
			foundUser = &user
			break
		}
	}

	if foundUser == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(foundUser); err != nil {
		h.logger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *UserHandler) handlePost(w http.ResponseWriter, r *http.Request) {
	var newUser structure.User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	if newUser.Name == "" || newUser.Age == 0 {
		http.Error(w, "Name and Age are required fields", http.StatusBadRequest)
		return
	}

	maxID := 0
	for _, user := range h.Users {
		if user.ID > maxID {
			maxID = user.ID
		}
	}
	newUser.ID = maxID + 1

	h.Users = append(h.Users, newUser)

	if err := writeUsersToFile(h.Users, h.UsersFile); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newUser); err != nil {
		h.logger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *UserHandler) handlePatch(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/user/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var targetUser *structure.User
	for i := range h.Users {
		if h.Users[i].ID == id {
			targetUser = &h.Users[i]
			break
		}
	}

	if targetUser == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	var updatedUser structure.User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	if updatedUser.Name != "" {
		targetUser.Name = updatedUser.Name
	}
	if updatedUser.Age != 0 {
		targetUser.Age = updatedUser.Age
	}

	if err := writeUsersToFile(h.Users, h.UsersFile); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(targetUser); err != nil {
		h.logger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (h *UserHandler) handleDelete(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/user/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	index := -1
	for i, user := range h.Users {
		if user.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	h.Users = append(h.Users[:index], h.Users[index+1:]...)

	if err := writeUsersToFile(h.Users, h.UsersFile); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func writeUsersToFile(users []structure.User, filePath string) error {
	usersData, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, usersData, 0644)
	if err != nil {
		return err
	}

	return nil
}
