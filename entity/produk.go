package entity

type Produk struct {
	Id         int64   `db:"id"`
	KategoriId int64   `db:"kategori_id"`
	Nama       string  `db:"nama"`
	Foto       string  `db:"foto"`
	Detail     string  `db:"detail"`
	Harga      float64 `db:"harga"`
	Stok       int64   `db:"stok"`
	IsCustom   bool    `db:"is_custom"`
	Bahan      string  `db:"bahan"`
	Ukuran     string  `db:"ukuran"`
}
