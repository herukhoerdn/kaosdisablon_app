package entity

import (
	"database/sql"
	"time"
)

type Desain struct {
	Id            int64     `db:"id"`
	UserId        int64     `db:"user_id"`
	FileDesain    string    `db:"file_desain"`
	Catatan       string    `db:"catatan"`
	Status        string    `db:"status"`
	TanggalUpload time.Time `db:"tanggal_upload"`
}
type DesainDetail struct {
	Id            int64 `json:"Id"`
	UserId        int64 `json:"UserId"`
	Username      sql.NullString
	NamaProduk    sql.NullString
	FileDesain    string    `json:"FileDesain"`
	Catatan       string    `json:"Catatan"`
	Status        string    `json:"Status"`
	TanggalUpload time.Time `json:"TanggalUpload"`
}
