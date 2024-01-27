package models

type Harga struct {
	AdminId      string `json:"admin_id,ommitempty"`
	HargaTopup   int64  `json:"harga_topup,ommitempty"`
	HargaBuyback int64  `json:"harga_buyback,ommitempty"`
}

func (Harga) TableName() string {
	return "harga"
}
