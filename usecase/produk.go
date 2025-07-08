package usecase

import (
	"context"
	"errors"
	"log"

	"github.com/kaosdisablon/entity"
)

func (u *usecase) InsertProduk(ctx context.Context, produk entity.Produk) (int64, error) {
	err := validateProduk(produk)
	if err != nil {
		return 0, err
	}

	id, err := u.repo.InsertProduk(ctx, produk)

	if err != nil {
		log.Println("Error insert")
		return 0, err
	}

	return id, nil
}

func validateProduk(produk entity.Produk) error {
	if produk.KategoriId == 0 {
		return errors.New("masukan kategori id ")
	}
	if produk.Nama == "" {
		return errors.New("masukan Nama barang")
	}
	if len(produk.Foto) == 0 {
		return errors.New("pilih foto")
	}
	if produk.Detail == "" {
		return errors.New("silahkan lihat detail barang")
	}
	if produk.Harga == 0 {
		return errors.New("harga tidak boleh kosong silahkan cek dan isi harga")
	}
	if produk.Stok == 0 {
		return errors.New("silahkan lihat")
	}
	if !produk.IsCustom {
		return errors.New("produk harus bertipe custom, silakan unggah desain")
	}
	if produk.Bahan == "" {
		return errors.New("masukan nama bahan")
	}
	if produk.Ukuran == "" {
		return errors.New("masukan ukuran")
	}
	return nil
}

func (u *usecase) GetProduk(ctx context.Context) ([]entity.Produk, error) {
	produk, err := u.repo.GetProduk(ctx)
	if err != nil {
		log.Println("Error get produk")
		return produk, err
	}
	return produk, nil
}

func (u *usecase) UpdateProduk(ctx context.Context, produk entity.Produk) (int64, error) {
	err := validateProduk(produk)
	if err != nil {
		log.Println("Error update")
		return 0, err
	}
	id, err := u.repo.UpdateProduk(ctx, produk)

	if err != nil {
		log.Println("Error update produk")
		return 0, err
	}
	return id, nil
}

func (u *usecase) DeleteProduk(ctx context.Context, id int64) error {
	err := u.repo.DeleteProduk(ctx, id)
	if err != nil {
		log.Println("Error Delete")
		return err
	}
	return nil

}
