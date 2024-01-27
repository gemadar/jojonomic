package models

type Harga struct {
	Id           int     `json:"id"`
	AdminId      string  `json:"admin_id,ommitempty"`
	HargaTopup   float64 `json:"harga_topup,ommitempty"`
	HargaBuyback float64 `json:"harga_buyback,ommitempty"`
}

func (Harga) TableName() string {
	return "harga"
}
