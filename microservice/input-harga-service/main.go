package main

import (
	"fmt"
	"input-harga-service/handlers"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(filepath.Join("./", ".env"))

	router := mux.NewRouter()
	router.HandleFunc("/api/input-harga", handlers.HargaHandler).Methods("POST")

	log.Print("server has started")

	//get the port from the environment variable
	port := os.Getenv("PORT")

	//pass the router and start listening with the server
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		log.Printf("error from router %v\n", err)
	}
}
