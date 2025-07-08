package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/kaosdisablon/entity"
)

type InvoicePesanan struct {
	Id          int64     `json:"id"`
	CheckoutId  int64     `json:"checkout_id"`
	UserId      int64     `json:"user_id"`
	TotalHarga  float64   `json:"total_harga"`
	TanggalBuat time.Time `json:"tanggal_buat"`
	FileInvoice string    `json:"file_invoice"`
}

func (h *handler) CreateInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	var Invoice InvoicePesanan
	if err := json.NewDecoder(r.Body).Decode(&Invoice); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	id, err := h.uc.CreateInvoice(context.Background(), Invoice.CheckoutId, Invoice.UserId, Invoice.TotalHarga, "")
	if err != nil {
		log.Println("Error CreateInvoice:", err)
		http.Error(w, "Gagal membuat invoice", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int64{"invoice_id": id})

}

func (h *handler) GetInvoiceByUserHandler(w http.ResponseWriter, r *http.Request) {
	userIdStr := r.URL.Query().Get("user_id")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}

	invoices, err := h.uc.GetInvoicesByUser(context.Background(), userId)
	if err != nil {
		http.Error(w, "Gagal mengambil data invoice", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invoices)
}
func (h *handler) GetCheckoutByDesainID(w http.ResponseWriter, r *http.Request) {
	desainIDStr := mux.Vars(r)["desain_id"]
	desainID, err := strconv.ParseInt(desainIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid desain_id", http.StatusBadRequest)
		return
	}

	checkout, err := h.uc.GetCheckoutByDesainID(context.Background(), desainID)
	if err != nil {
		http.Error(w, "Checkout tidak ditemukan", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(checkout)
}
