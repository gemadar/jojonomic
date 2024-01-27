package handlers

import (
	"buyback-service/models"
	"buyback-service/repositories"
	db "buyback-service/storage"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

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
	var trx models.Transaksi
	var currHarga models.Harga
	var currSaldo models.Rekening

	sid, err := shortid.Generate()
	json.NewDecoder(r.Body).Decode(&trx)

	producer, err := setupProducer()
	if err != nil {
		log.Fatalf("failed to initialize producer: %v", err)
	}
	defer producer.Close()

	x.Select("harga_buyback").Find(&currHarga)
	x.Select("saldo").Find(&currSaldo)

	if trx.Gram > currSaldo.Saldo {
		respondWithJSON(w, http.StatusOK, map[string]any{"error": true, "reff_id": sid, "message": "Kafka not ready"})
		return
	}

	if trx.HargaBuyback != currHarga.HargaBuyback {
		respondWithJSON(w, http.StatusOK, map[string]any{"error": true, "reff_id": sid, "message": "Kafka not ready"})
		return
	}

	data := SendMessageHandler(producer, trx)

	if data != "success" {
		respondWithJSON(w, http.StatusOK, map[string]any{"error": true, "reff_id": sid, "message": "Kafka not ready"})
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]any{"error": false, "reff_id": sid})
}

func SendMessageHandler(producer sarama.SyncProducer, h models.Transaksi) string {

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
