package models

type Topup struct {
	Gram  string `json:"gram,ommitempty"`
	Harga string `json:"harga,ommitempty"`
	Norek string `json:"norek"`
}

func (Topup) TableName() string {
	return "topup"
}
