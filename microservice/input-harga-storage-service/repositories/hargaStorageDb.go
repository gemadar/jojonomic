package repositories

import (
	"input-harga-storage-service/models"
	db "input-harga-storage-service/storage"
	"sync"
)

var x = db.InitDB()

type Hargas map[string][]models.Harga

type HargaStorage struct {
	Data Hargas
	mu   sync.RWMutex
}

func (ns *HargaStorage) Add(adminId string,
	h models.Harga) {
	ns.mu.Lock()
	defer ns.mu.Unlock()
	ns.Data[adminId] = append(ns.Data[adminId], h)
}

func HandleInputHarga(harga models.Harga) string {

	result := x.Model(&models.Harga{}).Where("id = ?", 1).
		Updates(models.Harga{AdminId: harga.AdminId, HargaTopup: harga.HargaTopup, HargaBuyback: harga.HargaBuyback})

	if result.Error != nil {
		return "failed"
	}
	return "success"
}
