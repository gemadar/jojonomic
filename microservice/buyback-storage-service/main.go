package main

import (
	"buyback-storage-service/handlers"
	"buyback-storage-service/repositories"
	db "buyback-storage-service/storage"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	store := &repositories.TrxStorage{
		Data: make(repositories.Transactions),
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
