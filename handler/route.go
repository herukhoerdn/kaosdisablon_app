package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kaosdisablon/usecase"
)

type handler struct {
	uc usecase.Usecase
}

func InitRoute(r *mux.Router, uc usecase.Usecase) {
	h := handler{
		uc: uc,
	}
	//Checkout
	r.HandleFunc("/checkout/{id}", h.UpdateCheckoutHandler).Methods("PUT")
	r.HandleFunc("/checkout", h.CreateCheckoutHandler).Methods("POST")
	r.HandleFunc("/checkout", h.GetCheckoutHandler).Methods("GET")
	r.HandleFunc("/checkout/{id}", h.DeleteCheckoutHandler).Methods("DELETE")
	//Print invoice
	r.HandleFunc("/checkout/by-desain", h.GetCheckoutFromDesain).Methods("GET")

	//Desain
	r.HandleFunc("/desain/{id}", h.UpdateDesainHandler).Methods("PUT")
	r.HandleFunc("/desain", h.CreateDesainHandler).Methods("POST")
	r.HandleFunc("/desain", h.GetDesainHandler).Methods("GET")
	r.HandleFunc("/desain/{id}", h.DeleteDesainHandler).Methods("DELETE")

	//update status desain
	r.HandleFunc("/desain/{id}/status", h.UpdateStatusDesainHandler).Methods("PUT")
	//detail pesanan
	r.HandleFunc("/desain", h.GetDesainHandler).Methods("GET")
	//kategori
	r.HandleFunc("/kategori/{id}", h.UpdateKategoriHandler).Methods("PUT")
	r.HandleFunc("/kategori", h.CreateKategoriHandler).Methods("POST")
	r.HandleFunc("/kategori", h.GetKategoriHandler).Methods("GET")
	r.HandleFunc("/kategori/{id}", h.DeleteKategoriHandler).Methods("DELETE")

	//Metode pembayaran
	r.HandleFunc("/metode_pembayaran", h.GetPembayaranHandler).Methods("GET")

	//Produk
	r.HandleFunc("/produk/{id}", h.UpdateProdukHandler).Methods("PUT")
	r.HandleFunc("/produk", h.CreateProdukHandler).Methods("POST")
	r.HandleFunc("/produk", h.GetProdukHandler).Methods("GET")
	r.HandleFunc("/produk/{id}", h.DeleteProdukHandler).Methods("DELETE")

	//User
	r.HandleFunc("/users/{id}", h.UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/users", h.CreateUserHandler).Methods("POST")
	r.HandleFunc("/users", h.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", h.DeleteUserHandler).Methods("DELETE")

	//login
	r.HandleFunc("/login", h.LoginHandler).Methods("POST")

	//Invoie
	r.HandleFunc("/invoice", h.CreateInvoiceHandler).Methods("POST")
	r.HandleFunc("/invoice", h.GetInvoiceByUserHandler).Methods("GET")
	//By id
	r.HandleFunc("/checkout/by-desain/{desain_id}", h.GetCheckoutByDesainID).Methods("GET")

	assetsDir := http.Dir("./assets")
	assetsHandler := http.StripPrefix("/assets/", http.FileServer(assetsDir))
	r.PathPrefix("/assets/").Handler(assetsHandler)

}
