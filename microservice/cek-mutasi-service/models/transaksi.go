package models

import (
	"time"
)

type Transaksi struct {
	DateUnix     int64     `gorm:"-:all" json:"date"`
	Date         time.Time `json:"-"`
	Type         string    `json:"type,ommitempty"`
	Gram         float64   `json:"gram,ommitempty"`
	HargaTopup   float64   `json:"harga_topup,ommitempty"`
	HargaBuyback float64   `json:"harga_buyback,ommitempty"`
}

func (Transaksi) TableName() string {
	return "transaksi"
}

// Struct for request body
type Request struct {
	Norek     string `gorm:"-" json:"norek,ommitempty"`
	StartDate int64  `gorm:"->:false;<-:create" json:"start_date"`
	EndData   int64  `gorm:"->:false;<-:create" json:"end_date"`
}

type Result struct {
	DateUnix     int64     `gorm:"-:all" json:"date"`
	Date         time.Time `json:"-"`
	Type         string    `json:"type,ommitempty"`
	Gram         float64   `json:"gram,ommitempty"`
	HargaTopup   int       `json:"harga_topup,ommitempty"`
	HargaBuyback int       `json:"harga_buyback,ommitempty"`
	Saldo        float64   `json:"saldo,ommitempty"`
}
