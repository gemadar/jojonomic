package models

type Rekening struct {
	Id    int     `json:"id"`
	Norek string  `json:"norek,ommitempty"`
	Saldo float64 `json:"saldo,ommitempty"`
}

func (Rekening) TableName() string {
	return "rekening"
}
