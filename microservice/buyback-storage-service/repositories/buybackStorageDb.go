package repositories

import (
	"buyback-storage-service/models"
	db "buyback-storage-service/storage"
	"math"
	"sync"
	"time"
)

var x = db.InitDB()

type Transactions map[string][]models.Transaksi

type TrxStorage struct {
	Data Transactions
	mu   sync.RWMutex
}

func (ns *TrxStorage) Add(norek string,
	h models.Transaksi) {
	ns.mu.Lock()
	defer ns.mu.Unlock()
	ns.Data[norek] = append(ns.Data[norek], h)
}

func HandleInputTransaksi(trx models.Transaksi) string {
	trx.Type = "buyback"
	trx.Date = time.Now()

	result := x.Create(&trx)

	if result.Error != nil {
		return "failed"
	}
	return "success"
}

func HandleUpdateSaldo(gram float64, norek string) string {
	var currSaldo models.Rekening

	x.Select("saldo").Where("norek = ?", norek).Find(&currSaldo)

	err := x.Model(&models.Rekening{}).Where("norek = ?", norek).Update("saldo", math.Round((currSaldo.Saldo-gram)*1000)/1000)

	if err.Error != nil {
		return "failed"
	}

	return "success"
}
