package usecase

import (
	"context"
	"log"

	"github.com/kaosdisablon/entity"
)

func (u *usecase) GetPembayaran(ctx context.Context) ([]entity.Pembayaran, error) {
	Pembayaran, err := u.repo.GetPembayaran(ctx)
	if err != nil {
		log.Println("error get pembayaran")
		return Pembayaran, err
	}
	return Pembayaran, nil
}
