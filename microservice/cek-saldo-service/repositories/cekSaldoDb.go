package repositories

import (
	"cek-saldo-service/models"
	db "cek-saldo-service/storage"
)

var x = db.InitDB()

func CekSaldo(norek string) (map[string]any, error) {
	var data models.Rekening

	err := x.Select("norek", "saldo").Where("norek = ?", norek).Find(&data)

	if err.Error != nil {
		return nil, err.Error
	}

	if err.RowsAffected == 0 {
		return nil, nil
	}

	return map[string]any{"norek": data.Norek, "saldo": data.Saldo}, nil
}
