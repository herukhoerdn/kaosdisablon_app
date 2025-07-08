package usecase

import (
	"context"
	"errors"
	"log"

	"github.com/kaosdisablon/entity"
)

func (u *usecase) InsertDesain(ctx context.Context, desain entity.Desain) (int64, error) {
	err := validateDesain(desain)
	if err != nil {
		return 0, err
	}

	id, err := u.repo.InsertDesain(ctx, desain)
	if err != nil {
		log.Println("Error insert")
		return 0, err
	}
	return id, nil
}

func validateDesain(desain entity.Desain) error {
	if desain.UserId == 0 {
		return errors.New("masukan user id")
	}
	if desain.FileDesain == "" {
		return errors.New("masukan file desain")
	}
	if desain.Catatan == "" {
		return errors.New("masukan catatan")
	}
	if desain.Status == "" {
		return errors.New("masukan status nya")
	}
	if desain.TanggalUpload.IsZero() {
		return errors.New("masukan tanggal upload")
	}
	return nil
}

func (u *usecase) GetDesain(ctx context.Context) ([]entity.Desain, error) {
	desain, err := u.repo.GetDesain(ctx)
	if err != nil {
		log.Println("Error get")
		return desain, err
	}
	return desain, nil
}

func (u *usecase) UpdateDesain(ctx context.Context, desain entity.Desain) (int64, error) {
	err := validateDesain(desain)
	if err != nil {
		log.Println("Error update")
		return 0, err
	}
	id, err := u.repo.UpdateDesain(ctx, desain)
	if err != nil {
		log.Println("Error delete")
		return 0, err
	}
	return id, nil
}
func (u *usecase) UpdateStatusOnly(ctx context.Context, id int64, status string) error {
	return u.repo.UpdateStatusOnly(ctx, id, status)
}

func (u *usecase) DeleteDesain(ctx context.Context, id int64) error {
	err := u.repo.DeleteDesain(ctx, id)
	if err != nil {
		log.Println("Error delete")
		return err
	}
	return nil
}
func (u *usecase) IsDesainUsed(ctx context.Context, desainId int64) (bool, error) {
	checkouts, err := u.repo.GetCheckout(ctx)
	if err != nil {
		return false, err
	}
	for _, c := range checkouts {
		if c.DesainId == desainId {
			return true, nil
		}
	}
	return false, nil
}
//Desail pesanan
func (u *usecase) GetDesainDetail(ctx context.Context) ([]entity.DesainDetail, error) {
	return u.repo.GetDesainDetail(ctx)
}
