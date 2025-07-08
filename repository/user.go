package repository

import (
	"context"
	"fmt"

	"github.com/kaosdisablon/entity"
)

func (r *repository) InsertUser(ctx context.Context, user entity.User) (int64, error) {
	query := "INSERT INTO users(username,password,no_telp,alamat,role) VALUES ($1,$2,$3,$4,$5) RETURNING id"
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Password, user.NoTelephone, user.Alamat, user.Role).Scan(&user.Id)
	if err != nil {
		return 0, err
	}
	return user.Id, nil
}

func (r *repository) GetUsers(ctx context.Context) ([]entity.User, error) {
	var user []entity.User
	query := "SELECT id,username,password,no_telp,alamat,role FROM users"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return user, nil
	}
	defer rows.Close()

	for rows.Next() {
		var usr entity.User

		if err := rows.Scan(&usr.Id, &usr.Username, &usr.Password, &usr.NoTelephone, &usr.Alamat, &usr.Role); err != nil {
			return user, err
		}

		user = append(user, usr)
	}
	return user, nil
}

func (r *repository) UpdateUser(ctx context.Context, user entity.User) (int64, error) {
	query := "UPDATE users SET username=$2, password=$3, no_telp=$4, alamat=$5, role=$6 WHERE id=$1 RETURNING id"
	err := r.db.QueryRowContext(ctx, query, user.Id, user.Username, user.Password, user.NoTelephone, user.Alamat, user.Role).Scan(&user.Id)
	if err != nil {
		return 0, err
	}
	return user.Id, nil
}

func (r *repository) DeleteUser(ctx context.Context, id int64) error {
	query :="DELETE FROM users WHERE id =$1"
	_,err := r.db.ExecContext(ctx,query,id)

	if err != nil {
		fmt.Println("Error",err)
		return err
	}
	return nil
}
