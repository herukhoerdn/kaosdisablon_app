package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/kaosdisablon/entity"
)

type ProdukHandler struct {
	Id            int64     `json:"id"`
	UserId        int64     `json:"user_id"`
	FileDesain    string    `json:"file_desain"`
	Catatan       string    `json:"catatan"`
	Status        string    `json:"status"`
	TanggalUpload time.Time `json:"tanggal_upload"`
}
type ProdukHandlerDetail struct {
	Id            int64     `json:"Id"`
	UserId        int64     `json:"UserId"`
	Username      string    `json:"Username"`
	NamaProduk    string    `json:"NamaProduk"`
	FileDesain    string    `json:"FileDesain"`
	Catatan       string    `json:"Catatan"`
	Status        string    `json:"Status"`
	TanggalUpload time.Time `json:"TanggalUpload"`
}

func (h *handler) CreateDesainHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	userIdStr := r.FormValue("user_id")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid user_id: %v", err), http.StatusBadRequest)
		return
	}

	file, handlerFile, err := r.FormFile("file_desain")
	if err != nil {
		http.Error(w, "Masukan file desain", http.StatusBadRequest)
		return
	}
	defer file.Close()

	filePath := "C:/Users/ihsan/go/src/kaosdisablon/assets" + handlerFile.Filename

	dst, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Gagal menyimpan file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Gagal menyalin file", http.StatusInternalServerError)
		return
	}

	catatan := r.FormValue("catatan")
	status := r.FormValue("status")
	tanggalUploadStr := r.FormValue("tanggal_upload")
	layout := "2006-01-02"
	tanggalUpload, err := time.Parse(layout, tanggalUploadStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid tanggal_upload format: %v", err), http.StatusBadRequest)
		return
	}

	// Masukkan ke database (hanya nama filenya)
	desainID, err := h.uc.InsertDesain(context.Background(), entity.Desain{
		UserId:        userId,
		FileDesain:    handlerFile.Filename,
		Catatan:       catatan,
		Status:        status,
		TanggalUpload: tanggalUpload,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating desain: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int64{"id": desainID})
}

// func (h *handler) GetDesainHandler(w http.ResponseWriter, r *http.Request) {
// 	desain, err := h.uc.GetDesain(context.Background())
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Error Get desain : %v", err), http.StatusInternalServerError)
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	w.Header().Set("Content-type", "application/json")
// 	json.NewEncoder(w).Encode(desain)
// }

func (h *handler) UpdateDesainHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid ID: %v", err), http.StatusBadRequest)
		return
	}

	userIdStr := r.FormValue("user_id")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid user_id: %v", err), http.StatusBadRequest)
		return
	}

	catatan := strings.TrimSpace(r.FormValue("catatan"))
	status := strings.TrimSpace(r.FormValue("status"))
	tanggalUploadStr := r.FormValue("tanggal_upload")
	layout := "2006-01-02"
	tanggalUpload, err := time.Parse(layout, tanggalUploadStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Format tanggal tidak valid: %v", err), http.StatusBadRequest)
		return
	}

	file, handlerFile, err := r.FormFile("file_desain")
	if err != nil {
		http.Error(w, "File desain wajib diupload", http.StatusBadRequest)
		return
	}
	defer file.Close()

	filePath := "C:/Users/ihsan/go/src/kaosdisablon/assets" + handlerFile.Filename
	dst, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Gagal menyimpan file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Gagal menyimpan isi file", http.StatusInternalServerError)
		return
	}

	// Cek debug semua nilai
	fmt.Println("DEBUG UPDATE")
	fmt.Println("catatan:", catatan)
	fmt.Println("status:", status)

	desainID, err := h.uc.UpdateDesain(context.Background(), entity.Desain{
		Id:            id,
		UserId:        userId,
		FileDesain:    handlerFile.Filename,
		Catatan:       catatan,
		Status:        status,
		TanggalUpload: tanggalUpload,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("Error updating desain: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(map[string]int64{"id": desainID})
}
func (h *handler) UpdateStatusDesainHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "ID tidak valid", http.StatusBadRequest)
		return
	}
	var payload struct {
		Status string `json:"status"`
	}
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Body tidak valid", http.StatusBadRequest)
		return
	}
	if payload.Status != "approved" && payload.Status != "rejected" {
		http.Error(w, "Status hanya boleh approved atau rejected", http.StatusBadRequest)
		return
	}
	desains, err := h.uc.GetDesain(context.Background())
	if err != nil {
		http.Error(w, "Gagal ambil desain", http.StatusInternalServerError)
		return
	}

	var desain entity.Desain
	found := false
	for _, d := range desains {
		if d.Id == id {
			desain = d
			found = true
			break
		}
	}
	if !found {
		http.Error(w, "Desain tidak ditemukan", http.StatusNotFound)
		return
	}
	desain.Status = payload.Status

	_, err = h.uc.UpdateDesain(context.Background(), desain)
	if err != nil {
		http.Error(w, "Gagal update status", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Status berhasil diupdate"))
}

func (h *handler) DeleteDesainHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("ID tidak valid: %v", err), http.StatusBadRequest)
		return
	}
	used, err := h.uc.IsDesainUsed(context.Background(), id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Gagal cek penggunaan desain: %v", err), http.StatusInternalServerError)
		return
	}
	if used {
		http.Error(w, "Desain sedang digunakan dalam transaksi, tidak dapat dihapus", http.StatusBadRequest)
		return
	}
	err = h.uc.DeleteDesain(context.Background(), id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Gagal menghapus desain: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Desain berhasil dihapus"})
}

// detail pesanan
func (h *handler) GetDesainHandler(w http.ResponseWriter, r *http.Request) {
	desains, err := h.uc.GetDesainDetail(context.Background())
	if err != nil {
		http.Error(w, fmt.Sprintf("Error Get desain : %v", err), http.StatusInternalServerError)
		return
	}

	var responses []ProdukHandlerDetail
	for _, d := range desains {
		responses = append(responses, ProdukHandlerDetail{
			Id:            d.Id,
			UserId:        d.UserId,
			Username:      nullStringToString(d.Username),
			NamaProduk:    nullStringToString(d.NamaProduk),
			FileDesain:    d.FileDesain,
			Catatan:       d.Catatan,
			Status:        d.Status,
			TanggalUpload: d.TanggalUpload,
		})
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(responses)
}

func nullStringToString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return "-"
}
