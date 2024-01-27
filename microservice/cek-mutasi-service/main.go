package main

import (
	"cek-mutasi-service/handlers"
	db "cek-mutasi-service/storage"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	// API Endpoint for cek mutasi
	router.HandleFunc("/api/mutasi", handlers.CekMutasi).Methods("POST")

	log.Print("server has started")
	//start the db
	db.InitDB()

	//get the port from the environment variable
	port := os.Getenv("PORT")

	//pass the router and start listening with the server
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		log.Printf("error from router %v\n", err)
	}
}
