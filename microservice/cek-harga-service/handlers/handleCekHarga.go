package handlers

import (
	"cek-harga-service/repositories"
	"encoding/json"
	"net/http"

	"github.com/teris-io/shortid"
)

// Generate JSON for API response
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func CekHarga(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := repositories.CekHarga()

	sid, err := shortid.Generate()

	if err != nil || data == nil {
		respondWithJSON(w, http.StatusOK, map[string]any{"error": true, "reff_id": sid, "message": "Kafka not ready"})
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]any{"error": false, "data": data})
}
