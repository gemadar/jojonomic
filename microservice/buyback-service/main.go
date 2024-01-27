package main

import (
	"buyback-service/handlers"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/buyback", handlers.HargaHandler).Methods("POST")

	log.Print("server has started")
	//start the db

	//get the port from the environment variable
	port := os.Getenv("PORT")

	//pass the router and start listening with the server
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		log.Printf("error from router %v\n", err)
	}
}
