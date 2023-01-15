package cartsdto

type CreateCartRequest struct {
	KodeProduk string `json:"kode_produk" form:"kode_produk" validate:"required"`
	NamaProduk string `json:"nama_produk" form:"nama_produk" validate:"required"`
	Kuantitas  int    `json:"kuantitas" form:"kuantitas" validate:"required"`
}

type UpdateCartRequest struct {
	KodeProduk string `json:"kode_produk" form:"kode_produk"`
	NamaProduk string `json:"nama_produk" form:"nama_produk"`
	Kuantitas  int    `json:"kuantitas" form:"kuantitas"`
}
