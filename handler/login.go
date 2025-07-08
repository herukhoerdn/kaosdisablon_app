package handler

import (
	"context"
	"encoding/json"
	"net/http"
)

func (h *handler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	user, err := h.uc.Login(context.Background(), loginRequest.Username, loginRequest.Password)
	if err != nil {
		http.Error(w, "Username atau password salah", http.StatusUnauthorized)
		return
	}

	// role buat frontend
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user_id":  user.Id,
		"username": user.Username,
		"role":     user.Role,
	})
}
