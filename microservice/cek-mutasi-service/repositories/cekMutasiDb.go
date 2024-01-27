package repositories

import (
	"cek-mutasi-service/models"
	db "cek-mutasi-service/storage"
	"time"
)

var x = db.InitDB()

func CekMutasi(norek string, stDate int64, enDate int64) ([]models.Result, error) {
	var data []models.Result

	err := x.Model(&models.Transaksi{}).Select("*, rekening.saldo").Joins("left join rekening on rekening.norek = transaksi.norek").
		Where("transaksi.norek = ? AND date BETWEEN ? AND ? ", norek,
			time.Unix(stDate, 0).UTC().Format("2006-01-02 15:04:05"),
			time.Unix(enDate, 0).UTC().Format("2006-01-02 15:04:05")).Scan(&data)

	if err.Error != nil {
		return nil, err.Error
	}

	if err.RowsAffected == 0 {
		return nil, nil
	}

	if err.RowsAffected != 0 {
		for a := range data {
			data[a].DateUnix = data[a].Date.Unix()
		}
	}

	return data, nil
}
