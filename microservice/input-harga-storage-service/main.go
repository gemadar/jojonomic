package main

import (
	"context"
	"fmt"
	"input-harga-storage-service/handlers"
	"input-harga-storage-service/repositories"
	db "input-harga-storage-service/storage"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	store := &repositories.HargaStorage{
		Data: make(repositories.Hargas),
	}

	ctx, cancel := context.WithCancel(context.Background())
	go handlers.SetupConsumerGroup(ctx, store)
	defer cancel()

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
