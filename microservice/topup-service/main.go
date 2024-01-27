package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"topup-service/handlers"
	db "topup-service/storage"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/topup", handlers.HargaHandler).Methods("POST")

	log.Print("server has started")
	db.InitDB()

	//get the port from the environment variable
	port := os.Getenv("PORT")

	//pass the router and start listening with the server
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		log.Printf("error from router %v\n", err)
	}
}
