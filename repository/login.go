package repository

import (
	"context"

	"github.com/kaosdisablon/entity"
)

func (r *repository) Login(ctx context.Context, username, password string) (entity.User, error) {
	var user entity.User
	query := "SELECT id, username, password, no_telp, alamat, role FROM users WHERE username = $1 AND password = $2 LIMIT 1"
	err := r.db.QueryRowContext(ctx, query, username, password).Scan(
		&user.Id, &user.Username, &user.Password,
		&user.NoTelephone, &user.Alamat, &user.Role,
	)
	if err != nil {
		return user, err
	}
	return user, nil
}
