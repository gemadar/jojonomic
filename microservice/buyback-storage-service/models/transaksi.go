package models

import (
	"time"
)

type Transaksi struct {
	Date         time.Time `json:"-"`
	Type         string    `json:"type,ommitempty"`
	Gram         float64   `json:"gram,ommitempty"`
	HargaTopup   int       `json:"harga_topup,ommitempty"`
	HargaBuyback int       `json:"harga,ommitempty"`
	Norek        string    `json:"norek,ommitempty"`
}

func (Transaksi) TableName() string {
	return "transaksi"
}
