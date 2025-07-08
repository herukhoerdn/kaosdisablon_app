package entity

type User struct {
	Id         int64  `db:"id"`
	Username   string `db:"username"`
	Password   string `db:"password"`
	NoTelephone string `db:"no_telp"`
	Alamat     string `db:"alamat"`
	Role       string `db:"role"`
}
