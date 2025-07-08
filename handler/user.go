package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kaosdisablon/entity"
)

type UserHandler struct {
	Id          int64  `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	NoTelephone string `json:"no_telephone"`
	Alamat      string `json:"alamat"`
	Role        string `json:"role"`
}

func (h *handler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user UserHandler
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	userID, err := h.uc.InsertUser(context.Background(), entity.User{
		Username:    user.Username,
		Password:    user.Password,
		NoTelephone: user.NoTelephone,
		Alamat:      user.Alamat,
		Role:        user.Role,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating user: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int64{"id": userID})
}
func (h *handler) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := h.uc.GetUsers(context.Background())
	if err != nil {
		http.Error(w, fmt.Sprintf("Error get user: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
func (h *handler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user UserHandler
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid id: %v", err), http.StatusBadRequest)
		return
	}

	userID, err := h.uc.UpdateUser(context.Background(), entity.User{
		Id:          id,
		Username:    user.Username,
		Password:    user.Password,
		NoTelephone: user.NoTelephone,
		Alamat:      user.Alamat,
		Role:        user.Role,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating user: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int64{"id": userID})
}
func (h *handler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid id: %v", err), http.StatusBadRequest)
		return
	}

	err = h.uc.DeleteUser(context.Background(), id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error deleting user: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "success delete"})
}
