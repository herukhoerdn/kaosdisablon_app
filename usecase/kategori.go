package usecase

import (
	"context"
	"errors"
	"log"

	"github.com/kaosdisablon/entity"
)

func (u *usecase) InsertKategori(ctx context.Context, kategori entity.Kategori) (int64, error) {
	err := validateKategori(kategori)
	if err != nil {
		return 0, err
	}

	id, err := u.repo.InsertKategori(ctx, kategori)

	if err != nil {
		log.Println("error insert ")
		return 0, err
	}

	return id, nil
}

func validateKategori(kategori entity.Kategori) error {
	if kategori.Nama == "" {
		return errors.New("silahkan inputkan")
	}
	return nil
}

func (u *usecase) GetKategori(ctx context.Context) ([]entity.Kategori, error) {
	kategori, err := u.repo.GetKategori(ctx)
	if err != nil {
		log.Println("Error Get kategori")
		return kategori, err
	}
	return kategori, nil
}

func (u *usecase) UpdateKategori(ctx context.Context, kategori entity.Kategori) (int64, error) {
	err := validateKategori(kategori)
	if err != nil {
		log.Println("Error update")
		return 0, err
	}
	id, err := u.repo.Updatekategori(ctx, kategori)

	if err != nil {
		log.Println("Error update")
		return 0, err
	}
	return id, nil
}
func (u *usecase) DeleteKategori(ctx context.Context, id int64) error {
	err := u.repo.DeleteKategori(ctx, id)
	if err != nil {
		log.Println("Error Delete")
		return err
	}
	return nil
}
