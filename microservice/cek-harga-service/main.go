package main

import (
	"cek-harga-service/handlers"
	db "cek-harga-service/storage"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// API Endpoint for cek harga
	router.HandleFunc("/api/check-harga", handlers.CekHarga).Methods("GET")

	//start the db
	db.InitDB()

	//get the port from the environment variable
	port := os.Getenv("PORT")

	log.Print("server has started")

	//pass the router and start listening with the server
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		log.Printf("error from router %v\n", err)
	}
}
