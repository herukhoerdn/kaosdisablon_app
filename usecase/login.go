package usecase

import (
	"context"

	"github.com/kaosdisablon/entity"
)

func (u *usecase) Login(ctx context.Context, username, password string) (entity.User, error) {
	return u.repo.Login(ctx, username, password)
}
