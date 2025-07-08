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

type KategoriHandler struct {
	Id   int64  `json:"id"`
	Nama string `json:"nama"`
}

func (h *handler) CreateKategoriHandler(w http.ResponseWriter, r *http.Request) {
	var kategori KategoriHandler
	if err := json.NewDecoder(r.Body).Decode(&kategori); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	kategoriID, err := h.uc.InsertKategori(context.Background(), entity.Kategori{
		Nama: kategori.Nama,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("Error cerating kaategori: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int64{"id": kategoriID})

}

func (h *handler) GetKategoriHandler(w http.ResponseWriter, r *http.Request) {
	kategori, err := h.uc.GetKategori(context.Background())
	if err != nil {
		http.Error(w, fmt.Sprintf("Error get kategori: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(kategori)
}
func (h *handler) UpdateKategoriHandler(w http.ResponseWriter, r *http.Request) {
	var kategori KategoriHandler
	if err := json.NewDecoder(r.Body).Decode(&kategori); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid id: %v", err), http.StatusBadRequest)
		return
	}

	kategoriID, err := h.uc.UpdateKategori(context.Background(), entity.Kategori{
		Id:   id,
		Nama: kategori.Nama,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating kategori: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(map[string]int64{"id": kategoriID})
}
func (h *handler) DeleteKategoriHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid id: %v", err), http.StatusInternalServerError)
		return
	}
	err = h.uc.DeleteKategori(context.Background(), id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error deleting kategori: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "succses delete"})
}
