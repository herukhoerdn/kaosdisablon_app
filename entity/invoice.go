package entity

import "time"

type Invoice struct {
	Id          int64     `db:"id"`
	CheckoutId  int64     `db:"checkout_id"`
	UserId      int64     `db:"user_id"`
	TotalHarga  float64   `db:"total_harga"`
	TanggalBuat time.Time `db:"tanggal_buat"`
	FileInvoice string    `db:"file_invoice"` //fiel invoice nya itu akan di simpan dalam format pdf
}
