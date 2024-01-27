package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"topup-service/models"
	"topup-service/repositories"
	db "topup-service/storage"

	"github.com/IBM/sarama"
	"github.com/teris-io/shortid"
)

var x = db.InitDB()

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func HargaHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var top models.Topup
	var currHarga models.Harga

	sid, err := shortid.Generate()

	json.NewDecoder(r.Body).Decode(&top)

	gr, _ := strconv.ParseFloat(top.Gram, 64)
	harga, _ := strconv.Atoi(top.Harga)

	if gr < 0.001 {
		respondWithJSON(w, http.StatusOK, map[string]any{"error": true, "reff_id": sid, "message": "Kafka not ready"})
		return
	}

	// Split the string at the decimal point
	parts := strings.Split(top.Gram, ".")

	// Check if there are exactly two digits after the decimal point
	if len(parts[1]) != 3 {
		respondWithJSON(w, http.StatusOK, map[string]any{"error": true, "reff_id": sid, "message": "Kafka not ready"})
		return
	}

	producer, err := setupProducer()
	if err != nil {
		log.Fatalf("failed to initialize producer: %v", err)
	}
	defer producer.Close()

	x.Select("harga_topup").Find(&currHarga)

	if harga != currHarga.HargaTopup {
		respondWithJSON(w, http.StatusOK, map[string]any{"error": true, "reff_id": sid, "message": "Kafka not ready"})
		return
	}

	data := SendMessageHandler(producer, top)

	if data != "success" {
		respondWithJSON(w, http.StatusOK, map[string]any{"error": true, "reff_id": sid, "message": "Kafka not ready"})
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]any{"error": false, "reff_id": sid})
}

func SendMessageHandler(producer sarama.SyncProducer, h models.Topup) string {

	err := repositories.SendKafkaMessage(producer, h)
	if err != nil {
		return "error"
	}

	return "success"
}

func setupProducer() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer([]string{os.Getenv("KAFKA_BROKERS")},
		config)
	if err != nil {
		return nil, fmt.Errorf("failed to setup producer: %w", err)
	}
	return producer, nil
}
