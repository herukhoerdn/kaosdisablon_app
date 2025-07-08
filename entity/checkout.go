package entity

import "time"

type Checkout struct {
	Id               int64     `db:"id"`
	UserId           int64     `db:"user_id"`
	ProdukId         int64     `db:"produk_id"`
	Kuantiti         int64     `db:"kuantiti"`
	TotalHarga       float64   `db:"total_harga"`
	MetodePembayaran string    `db:"metode_pembayaran"`
	TanggalOrder     time.Time `db:"tanggal_order"`
	Status           string    `db:"status"`
	DesainId         int64     `db:"desain_id"`
}
