package repository

import (
	"context"
	"database/sql"

	"github.com/kaosdisablon/entity"
)

type repository struct {
	db *sql.DB
}

type Repository interface {
	//Checkout
	InsertCheckout(ctx context.Context, checkout entity.Checkout) (int64, error)
	GetCheckout(ctx context.Context) ([]entity.Checkout, error)
	UpdateCheckout(ctx context.Context, checkout entity.Checkout) (int64, error)
	DeleteCheckout(ctx context.Context, id int64) error
	//Print invoice
	FindCheckoutByDesainID(ctx context.Context, desainId int64) (*entity.Checkout, error)
	//cetak pdf invoice
	UpdateInvoiceFilePath(ctx context.Context, invoiceID int64, fileName string) error

	//Desain
	InsertDesain(ctx context.Context, desain entity.Desain) (int64, error)
	GetDesain(ctx context.Context) ([]entity.Desain, error)
	UpdateDesain(ctx context.Context, desain entity.Desain) (int64, error)
	DeleteDesain(ctx context.Context, id int64) error
	//desain update status
	UpdateStatusOnly(ctx context.Context, id int64, status string) error
	//detail pesanan
	GetDesainDetail(ctx context.Context) ([]entity.DesainDetail, error)

	//Kategori
	InsertKategori(ctx context.Context, kategori entity.Kategori) (int64, error)
	GetKategori(ctx context.Context) ([]entity.Kategori, error)
	Updatekategori(ctx context.Context, kategori entity.Kategori) (int64, error)
	DeleteKategori(ctx context.Context, id int64) error

	//Metode pembayaran
	GetPembayaran(ctx context.Context) ([]entity.Pembayaran, error)

	//Produk
	InsertProduk(ctx context.Context, produk entity.Produk) (int64, error)
	GetProduk(ctx context.Context) ([]entity.Produk, error)
	UpdateProduk(ctx context.Context, produk entity.Produk) (int64, error)
	DeleteProduk(ctx context.Context, id int64) error

	//User
	InsertUser(ctx context.Context, user entity.User) (int64, error)
	GetUsers(ctx context.Context) ([]entity.User, error)
	UpdateUser(ctx context.Context, user entity.User) (int64, error)
	DeleteUser(ctx context.Context, id int64) error

	//Login
	Login(ctx context.Context, username, password string) (entity.User, error)

	//Invoice
	CreateInvoice(ctx context.Context, invoice entity.Invoice) (int64, error)
	GetInvoicesByUser(ctx context.Context, userId int64) ([]entity.Invoice, error)
	//By id
	GetCheckoutByDesainID(ctx context.Context, desainId int64) (*entity.Checkout, error)
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}
