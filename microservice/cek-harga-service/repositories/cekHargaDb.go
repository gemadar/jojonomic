package repositories

import (
	"cek-harga-service/models"
	db "cek-harga-service/storage"
)

var x = db.InitDB()

func CekHarga() (map[string]any, error) {
	var data models.Harga

	err := x.Select("harga_buyback", "harga_topup").Find(&data)

	if err.Error != nil {
		return nil, err.Error
	}

	if err.RowsAffected == 0 {
		return nil, nil
	}

	return map[string]any{"harga_buyback": data.HargaBuyback, "harga_topup": data.HargaTopup}, nil
}
