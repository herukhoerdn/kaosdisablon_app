package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type PembayaranHandler struct {
	Id   int64  `json:"id"`
	Nama string `json:"nama"`
}

func (h *handler) GetPembayaranHandler(w http.ResponseWriter, r *http.Request) {
	pembayaran, err := h.uc.GetPembayaran(context.Background())
	if err != nil {
		http.Error(w, fmt.Sprintf("Error get pembayaran: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(pembayaran)
}
