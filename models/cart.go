package models

import "time"

type Cart struct {
	ID         int       `json:"id" gorm:"primary_key:auto_increment"`
	KodeProduk string    `json:"kode_produk" form:"kode_produk" gorm:"type: varchar(255)"`
	NamaProduk string    `json:"nama_produk" form:"nama_produk" gorm:"type: varchar(255)"`
	Kuantitas  int       `json:"kuantitas" form:"kuantitas" gorm:"type: int"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
}

type CartResponse struct {
	ID         int    `json:"id"`
	KodeProduk string `json:"kode_produk"`
	NamaProduk string `json:"nama_produk"`
	Kuantitas  int    `json:"kuantitas"`
}

func (CartResponse) TableName() string {
	return "carts"
}
