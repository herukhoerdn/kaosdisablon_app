package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/kaosdisablon/entity"
)

type Checkout struct {
	Id               int64     `json:"id"`
	UserId           int64     `json:"user_id"`
	ProdukId         int64     `json:"produk_id"`
	Kuantiti         int64     `json:"kuantiti"`
	TotalHarga       float64   `json:"total_harga"`
	MetodePembayaran string    `json:"metode_pembayaran"`
	TanggalOrder     time.Time `json:"tanggal_order"`
	Status           string    `json:"status"`
	Desain           int64     `json:"desain"`
}

// POST /checkout
func (h *handler) CreateCheckoutHandler(w http.ResponseWriter, r *http.Request) {
	var checkout Checkout
	if err := json.NewDecoder(r.Body).Decode(&checkout); err != nil {
		log.Println("error:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	checkoutID, err := h.uc.InsertCheckout(context.Background(), entity.Checkout{
		UserId:           checkout.UserId,
		ProdukId:         checkout.ProdukId,
		Kuantiti:         checkout.Kuantiti,
		TotalHarga:       checkout.TotalHarga,
		MetodePembayaran: checkout.MetodePembayaran,
		TanggalOrder:     checkout.TanggalOrder,
		Status:           checkout.Status,
		DesainId:         checkout.Desain,
	})
	if err != nil {
		log.Println("error:", err)
		http.Error(w, fmt.Sprintf("Error creating : %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(map[string]int64{"id": checkoutID})
}

func (h *handler) GetCheckoutHandler(w http.ResponseWriter, r *http.Request) {
	checkout, err := h.uc.GetCheckout(context.Background())
	if err != nil {
		http.Error(w, fmt.Sprintf("Error Get : %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(checkout)
}

func (h *handler) UpdateCheckoutHandler(w http.ResponseWriter, r *http.Request) {
	var checkout Checkout
	if err := json.NewDecoder(r.Body).Decode(&checkout); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid id : %v", err), http.StatusBadRequest)
		return
	}

	checkoutID, err := h.uc.UpdateCheckout(context.Background(), entity.Checkout{
		Id:               id,
		UserId:           checkout.UserId,
		ProdukId:         checkout.ProdukId,
		Kuantiti:         checkout.Kuantiti,
		TotalHarga:       checkout.TotalHarga,
		MetodePembayaran: checkout.MetodePembayaran,
		TanggalOrder:     checkout.TanggalOrder,
		Status:           checkout.Status,
		DesainId:         checkout.Desain,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("Error updating: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(map[string]int64{"id": checkoutID})
}

func (h *handler) DeleteCheckoutHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid id: %v", err), http.StatusBadRequest)
		return
	}

	err = h.uc.DeleteCheckout(context.Background(), id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error deleting: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "success delete"})
}
func (h *handler) GetCheckoutFromDesain(w http.ResponseWriter, r *http.Request) {
	desainIdStr := r.URL.Query().Get("desain_id")
	desainId, err := strconv.ParseInt(desainIdStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid desain_id", http.StatusBadRequest)
		return
	}

	checkout, err := h.uc.FetchCheckoutByDesainID(context.Background(), desainId)
	if err != nil {
		http.Error(w, "Checkout tidak ditemukan", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(checkout)
}
