package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"topup-storage-service/handlers"
	"topup-storage-service/repositories"
	db "topup-storage-service/storage"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	store := &repositories.TopStorage{
		Data: make(repositories.Topups),
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
