package handlers

import (
	"cek-saldo-service/models"
	"cek-saldo-service/repositories"
	"encoding/json"
	"net/http"

	"github.com/teris-io/shortid"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func CekSaldo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var norek models.Rekening
	sid, err := shortid.Generate()
	json.NewDecoder(r.Body).Decode(&norek)

	if norek.Norek == "" {
		respondWithJSON(w, http.StatusOK, map[string]any{"error": true, "reff_id": sid, "message": "Kafka not ready"})
		return
	}

	data, err := repositories.CekSaldo(norek.Norek)

	if err != nil || data == nil {
		respondWithJSON(w, http.StatusOK, map[string]any{"error": true, "reff_id": sid, "message": "Kafka not ready"})
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]any{"error": false, "data": data})
}
