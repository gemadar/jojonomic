package repositories

import (
	"math"
	"strconv"
	"sync"
	"topup-storage-service/models"
	db "topup-storage-service/storage"
)

var x = db.InitDB()

type Topups map[string][]models.Topup

type TopStorage struct {
	Data Topups
	mu   sync.RWMutex
}

func (ns *TopStorage) Add(norek string,
	h models.Topup) {
	ns.mu.Lock()
	defer ns.mu.Unlock()
	ns.Data[norek] = append(ns.Data[norek], h)
}

func HandleInputTransaksi(top models.Topup) string {

	result := x.Select("gram", "harga").Create(&top)

	if result.Error != nil {
		return "failed"
	}
	return "success"
}

func HandleUpdateSaldo(gram string, norek string) string {
	var currSaldo models.Rekening
	gr, _ := strconv.ParseFloat(gram, 64)
	x.Select("saldo").Where("norek = ?", norek).Find(&currSaldo)

	err := x.Model(&models.Rekening{}).Where("norek = ?", norek).Update("saldo", math.Round((currSaldo.Saldo+gr)*1000)/1000)

	if err.Error != nil {
		return "failed"
	}

	return "success"
}
