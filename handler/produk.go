package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kaosdisablon/entity"
)

type Produk struct {
	Id         int64   `json:"id"`
	KategoriId int64   `json:"kategori_id"`
	Nama       string  `json:"nama"`
	Foto       string  `json:"foto"`
	Detail     string  `json:"detail"`
	Harga      float64 `json:"harga"`
	Stok       int64   `json:"stok"`
	IsCustom   bool    `json:"is_custom"`
	Bahan      string  `json:"bahan"`
	Ukuran     string  `json:"ukuran"`
}

// CREATE
func (h *handler) CreateProdukHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	kategoriId, _ := strconv.ParseInt(r.FormValue("kategori_id"), 10, 64)
	harga, _ := strconv.ParseFloat(r.FormValue("harga"), 64)
	stok, _ := strconv.ParseInt(r.FormValue("stok"), 10, 64)
	isCustom := r.FormValue("is_custom") == "true"

	nama := r.FormValue("nama")
	detail := r.FormValue("detail")
	bahan := r.FormValue("bahan")
	ukuran := r.FormValue("ukuran")

	// Handle file foto
	file, handlerFile, err := r.FormFile("foto")
	if err != nil {
		http.Error(w, fmt.Sprintf("Foto wajib diunggah: %v", err), http.StatusBadRequest)
		return
	}
	defer file.Close()

	filePath := fmt.Sprintf("assets/%s", handlerFile.Filename)
	dst, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Gagal menyimpan file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()
	io.Copy(dst, file)

	// Simpan foto ke Database
	produkID, err := h.uc.InsertProduk(context.Background(), entity.Produk{
		KategoriId: kategoriId,
		Nama:       nama,
		Foto:       handlerFile.Filename,
		Detail:     detail,
		Harga:      harga,
		Stok:       stok,
		IsCustom:   isCustom,
		Bahan:      bahan,
		Ukuran:     ukuran,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("Error insert: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int64{"id": produkID})
}

func (h *handler) GetProdukHandler(w http.ResponseWriter, r *http.Request) {
	produk, err := h.uc.GetProduk(context.Background())
	if err != nil {
		http.Error(w, fmt.Sprintf("Error get produk: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produk)
}

func (h *handler) UpdateProdukHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "ID tidak valid", http.StatusBadRequest)
		return
	}

	kategoriId, _ := strconv.ParseInt(r.FormValue("kategori_id"), 10, 64)
	harga, _ := strconv.ParseFloat(r.FormValue("harga"), 64)
	stok, _ := strconv.ParseInt(r.FormValue("stok"), 10, 64)
	isCustom := r.FormValue("is_custom") == "true"

	nama := r.FormValue("nama")
	detail := r.FormValue("detail")
	bahan := r.FormValue("bahan")
	ukuran := r.FormValue("ukuran")

	// Foto default tidak berubah
	foto := r.FormValue("foto_lama")

	// Ada file baru diunggah
	file, handlerFile, err := r.FormFile("foto")
	if err == nil {
		defer file.Close()
		filePath := fmt.Sprintf("assets/%s", handlerFile.Filename)
		dst, err := os.Create(filePath)
		if err == nil {
			defer dst.Close()
			io.Copy(dst, file)
			foto = handlerFile.Filename // update foto
		}
	}

	produkID, err := h.uc.UpdateProduk(context.Background(), entity.Produk{
		Id:         id,
		KategoriId: kategoriId,
		Nama:       nama,
		Foto:       foto,
		Detail:     detail,
		Harga:      harga,
		Stok:       stok,
		IsCustom:   isCustom,
		Bahan:      bahan,
		Ukuran:     ukuran,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("Error update produk: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]int64{"id": produkID})
}

func (h *handler) DeleteProdukHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "ID tidak valid", http.StatusBadRequest)
		return
	}

	err = h.uc.DeleteProduk(context.Background(), id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error delete: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Berhasil dihapus"})
}
