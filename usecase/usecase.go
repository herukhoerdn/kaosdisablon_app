package usecase

import (
	"context"

	"github.com/kaosdisablon/entity"
	"github.com/kaosdisablon/repository"
)

type usecase struct {
	repo repository.Repository
}

type Usecase interface {
	//checkout
	InsertCheckout(ctx context.Context, checkout entity.Checkout) (int64, error)
	GetCheckout(ctx context.Context) ([]entity.Checkout, error)
	UpdateCheckout(ctx context.Context, checkout entity.Checkout) (int64, error)
	DeleteCheckout(ctx context.Context, id int64) error
	//Print Invoice
	FetchCheckoutByDesainID(ctx context.Context, desainId int64) (*entity.Checkout, error)
	//cetak pdf invoice
	GenerateInvoicePDF(filePath string, inv entity.Invoice) error

	//Desain
	InsertDesain(ctx context.Context, desain entity.Desain) (int64, error)
	GetDesain(ctx context.Context) ([]entity.Desain, error)
	UpdateDesain(ctx context.Context, desain entity.Desain) (int64, error)
	DeleteDesain(ctx context.Context, id int64) error
	//update desain status dan juga id
	UpdateStatusOnly(ctx context.Context, id int64, status string) error
	IsDesainUsed(ctx context.Context, desainId int64) (bool, error)
	//detail pesanan
	GetDesainDetail(ctx context.Context) ([]entity.DesainDetail, error)

	//kategori
	InsertKategori(ctx context.Context, kategori entity.Kategori) (int64, error)
	GetKategori(ctx context.Context) ([]entity.Kategori, error)
	UpdateKategori(ctx context.Context, kategori entity.Kategori) (int64, error)
	DeleteKategori(ctx context.Context, id int64) error

	//Metode pembayaran
	GetPembayaran(ctx context.Context) ([]entity.Pembayaran, error)

	//produk
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
	CreateInvoice(ctx context.Context, checkoutId, userId int64, totalHarga float64, fileInvoice string) (int64, error)
	GetInvoicesByUser(ctx context.Context, userId int64) ([]entity.Invoice, error)
	//By id
	GetCheckoutByDesainID(ctx context.Context, desainId int64) (*entity.Checkout, error)
}

func NewUsecase(r repository.Repository) Usecase {
	return &usecase{
		repo: r,
	}
}
