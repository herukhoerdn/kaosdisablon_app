package usecase

import (
	"context"
	"errors"
	"log"

	"github.com/kaosdisablon/entity"
)

func (u *usecase) InsertUser(ctx context.Context, user entity.User) (int64, error) {
	err := validateUser(user)
	if err != nil {
		return 0, err
	}

	id, err := u.repo.InsertUser(ctx, user)

	if err != nil {
		log.Println("error insert user")
		return 0, err
	}

	return id, nil

}

func validateUser(user entity.User) error {
	if user.Username == "" {
		return errors.New("nama tidak boleh kosong")
	}

	if user.Password == "" {
		return errors.New("password tidak boleh kosong")
	}
	if user.NoTelephone == "" {
		return errors.New("nomor tidak boleh kosong")
	}
	if user.Alamat == "" {
		return errors.New("alamat tidak boleh kosong")
	}
	if user.Role == "" {
		return errors.New("role harus diisi")
	}

	return nil
}
func (u *usecase) GetUsers(ctx context.Context) ([]entity.User, error) {
	users, err := u.repo.GetUsers(ctx)
	if err != nil {
		log.Println("error get user")
		return users, err
	}
	return users, nil
}
func (u *usecase) UpdateUser(ctx context.Context, user entity.User) (int64, error) {
	err := validateUser(user)
	if err != nil {
		log.Println("error update user")
		return 0, err
	}
	id, err := u.repo.UpdateUser(ctx, user)

	if err != nil {
		log.Println("error update user")
		return 0, err
	}
	return id, nil
}
func (u *usecase) DeleteUser(ctx context.Context, id int64) error {
	// Perform deletion logic

	// If deletion is successful
	err := u.repo.DeleteUser(ctx, id)
	if err != nil {
		log.Println("error update user")
		return err
	}

	return nil

	// If an error occurs
	// return 0, someError
}
