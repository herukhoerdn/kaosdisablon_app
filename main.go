package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/kaosdisablon/config"
	"github.com/kaosdisablon/handler"
	"github.com/kaosdisablon/repository"
	"github.com/kaosdisablon/usecase"
)

func main() {
	db := config.InitDB()
	defer db.Close()

	repo := repository.NewRepository(db)
	uc := usecase.NewUsecase(repo)
	router := mux.NewRouter()
	handler.InitRoute(router, uc)
	//desain html
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("C:/Users/ihsan/go/src/kaosdisablon/assets"))))

	headersOk := handlers.AllowedHeaders([]string{"Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "FETCH"})

	// Start the HTTP server
	fmt.Println("server is running on port :8080")
	log.Fatal(http.ListenAndServe("localhost:8080", handlers.CORS(headersOk, originsOk, methodsOk)(router)))

}
