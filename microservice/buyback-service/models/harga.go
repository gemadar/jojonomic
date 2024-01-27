package models

type Harga struct {
	AdminId      string `json:"admin_id,ommitempty"`
	HargaTopup   int    `json:"harga_topup,ommitempty"`
	HargaBuyback int    `json:"harga_buyback,ommitempty"`
}

func (Harga) TableName() string {
	return "harga"
}
