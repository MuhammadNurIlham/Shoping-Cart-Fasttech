package cartsdto

type CartResponse struct {
	ID         int    `json:"id"`
	KodeProduk string `json:"kode_produk" form:"kode_produk" validate:"required"`
	NamaProduk string `json:"nama_produk" form:"nama_produk" validate:"required"`
	Kuantitas  int    `json:"kuantitas" form:"kuantitas" validate:"required"`
}
